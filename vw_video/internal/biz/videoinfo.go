package biz

import (
	"context"
	stderr "errors"
	"github.com/go-kratos/kratos/v2/log"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"util"
	utilCtx "util/context"
	"util/ffmpeg"
	"util/getid"
	"util/helper"
	"util/helper/file"
	videoinfov1 "vw_video/api/v1/videoinfo"
	"vw_video/cmd/gorm_gen/methods"
	"vw_video/internal/data/dal/model"
	"vw_video/internal/pkg/ecode/errdef"
)

type VideoInfoRepo interface {
	GetVideoInfoById(ctx context.Context, videoId int64) (*model.Video, error)
	GetVideoListByClass(ctx context.Context, class []string, pageNum int32, pageSize int32) ([]*model.Video, error)
	GetVideoFilePathById(ctx context.Context, videoId int64) (string, error)
	AddVideoHot(ctx context.Context, videoId, hot int64) error
	SetCoverPath(ctx context.Context, videoId int64, coverPath string) error
	CreateBasicVideoInfo(ctx context.Context, video *model.Video) error

	UpdateVideoDurationSizeInfo(ctx context.Context, videoId int64, duration string, size int64) error
	GetPublisherIdByVideoId(CTX context.Context, videoId int64) (publisherId int64, err error)
	UpdateVideoFilePath(ctx context.Context, videoId int64, videoFilePath string) error
}

type VideoInfoUsecase struct {
	repo   VideoInfoRepo
	tx     util.Transaction
	logger *log.Helper
}

func NewVideoInfoUsecase(repo VideoInfoRepo, tx util.Transaction, logger log.Logger) *VideoInfoUsecase {
	return &VideoInfoUsecase{
		repo:   repo,
		tx:     tx,
		logger: log.NewHelper(logger),
	}
}

func (uc *VideoInfoUsecase) GetVideoInfoById(ctx context.Context, videoId int64) (*model.Video, error) {
	videoDetail, err := uc.repo.GetVideoInfoById(ctx, videoId)
	if err != nil {
		if stderr.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.HandleError(errdef.ErrGetVideoInfoFailed, errdef.ErrVideoNotFound)
		}
		return nil, helper.HandleError(errdef.ErrGetVideoInfoFailed, err)
	}
	return videoDetail, nil
}

func (uc *VideoInfoUsecase) GetVideoListByClass(ctx context.Context, class []string, pageNum int32, pageSize int32) ([]*model.Video, error) {
	videoList, err := uc.repo.GetVideoListByClass(ctx, class, pageNum, pageSize)
	if err != nil {
		return nil, helper.HandleError(errdef.ErrGetVideoListFailed, err)
	}
	return videoList, nil
}

func (uc *VideoInfoUsecase) UploadVideoInfo(ctx context.Context, info *videoinfov1.VideoMetaInfo) error {
	videoInfo := &model.Video{
		VideoID:       getid.GetID(),
		Title:         info.Title,
		Description:   info.Description,
		Class:         strings.Join(info.Classes, methods.Separator),
		Tags:          strings.Join(info.Tags, methods.Separator),
		PublisherID:   info.PublisherId,
		PublisherName: info.PublisherName,
		Duration:      info.Duration,
	}
	err := uc.repo.CreateBasicVideoInfo(ctx, videoInfo)
	if err != nil {
		return helper.HandleError(errdef.ErrUploadVideoInfoFailed, err)
	}
	return nil
}

func (uc *VideoInfoUsecase) UploadVideoFile(stream grpc.ClientStreamingServer[videoinfov1.UploadVideoFileReq, emptypb.Empty]) error {
	uc.logger.Info("receiving the video file")
	var (
		retErr        error
		videoId       int64
		videoFilePath string
		videoFile     *os.File
	)

	defer func() {
		if videoFile != nil {
			_ = videoFile.Close()
			if retErr != nil {
				err := os.RemoveAll(filepath.Dir(videoFilePath))
				// TODO: use message queue to retry later
				if err != nil {

				}
			}
		}
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF { // 传输完成
			uc.logger.Infof("receive the video file %s complete.", videoFilePath)
			break
		}
		if err != nil {
			retErr = helper.HandleError(errdef.ErrUploadVideoFileFailed, err)
			return retErr
		}

		switch reqData := req.VideoFile.(type) {
		// 1. Handle the metadata of the video file
		// 1.1 Receive the metadata of the video file
		case *videoinfov1.UploadVideoFileReq_VideoMetadata:
			videoId = reqData.VideoMetadata.VideoId
			publisherId := reqData.VideoMetadata.PublisherId

			// 1.2 Create video file, which path is computed by the video id and publisher(user) id.
			videoFilePath = filepath.Join(
				resourcePath,
				strconv.FormatInt(publisherId, 10),
				videoPath,
				strconv.FormatInt(videoId, 10),
				reqData.VideoMetadata.VideoFileName)
			videoFile, err = file.CreateFile(videoFilePath, os.ModePerm)
			if err != nil {
				retErr = helper.HandleError(errdef.ErrUploadVideoFileFailed, err)
				return retErr
			}
			uc.logger.Info("create video file success")
		case *videoinfov1.UploadVideoFileReq_VideoFileData:
			// 2. Handle the data of the video cover file
			fileContent := reqData.VideoFileData
			_, err = videoFile.Write(fileContent)
			if err != nil {
				retErr = helper.HandleError(errdef.ErrUploadVideoFileFailed, err)
				return retErr
			}
		}
	}
	err := stream.SendAndClose(&emptypb.Empty{})
	if err != nil {
		retErr = helper.HandleError(errdef.ErrUploadVideoFileFailed, err)
		return retErr
	}

	// 3. Handle the video file: make DASH segments files.
	// 3.1 Make DASH directory, which path is computed by the video file path.
	_ = os.MkdirAll(filepath.Dir(videoFilePath)+"/"+dashPath, os.ModePerm)

	// 3.2 Make DASH segments files.
	err = ffmpeg.MakeDASH(videoFilePath, filepath.Dir(videoFilePath)+"/"+dashPath, dashName)
	if err != nil {
		retErr = helper.HandleError(errdef.ErrUploadVideoFileFailed, err)
		return retErr
	}
	uc.logger.Info("make DASH segments files success.")

	// 4. Update the video info record
	// 4.1 Get the duration of the video file.
	ctx, cancel := utilCtx.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = uc.tx.WithTx(ctx, func(ctx context.Context) error {
		duration, err := ffmpeg.GetVideoDuration(videoFilePath)
		if err != nil {
			return err
		}
		durationStr, _ := helper.SecondToTime(duration)

		// 4.2 Get the size of the video file.
		size, err := file.GetFileSize(videoFilePath)
		if err != nil {
			return err
		}

		// 4.3 Set the video size and duration information of the video info record.
		ctx, cancel := utilCtx.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		err = uc.repo.UpdateVideoDurationSizeInfo(ctx, videoId, durationStr, size)
		if err != nil {
			return err
		}

		// 4.4 Set video_path of the video info record.
		err = uc.repo.UpdateVideoFilePath(ctx, videoId, videoFilePath)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		retErr = helper.HandleError(errdef.ErrUploadVideoFileFailed, err)
		return retErr
	}

	uc.logger.Info("update video size and duration information success.")
	uc.logger.Info("upload video file success.")
	return nil
}

func (uc *VideoInfoUsecase) UploadVideoCover(stream grpc.ClientStreamingServer[videoinfov1.UploadVideoCoverReq, emptypb.Empty]) error {
	var (
		retErr        error
		coverFilePath string
		videoId       int64
		coverFile     *os.File
	)

	defer func() {
		if coverFile != nil {
			_ = coverFile.Close()
			if retErr != nil {
				_ = os.Remove(coverFilePath)
			}
		}
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF { // 传输完成
			break
		}
		if err != nil {
			retErr = helper.HandleError(errdef.ErrUploadVideoCoverFailed, err)
			return retErr
		}

		switch reqData := req.CoverFile.(type) {
		case *videoinfov1.UploadVideoCoverReq_CoverMetadata:
			// 1. Handle the metadata of the video cover file
			// 1.1 Receive the metadata of the video cover file
			videoId = reqData.CoverMetadata.VideoId
			publisherId := reqData.CoverMetadata.PublisherId

			// 1.2 Create cover file, which path is computed by the video id and publisher(user) id.
			coverFilePath = filepath.Join(
				resourcePath,
				strconv.FormatInt(publisherId, 10),
				videoPath,
				strconv.FormatInt(videoId, 10),
				coverName+filepath.Ext(reqData.CoverMetadata.CoverFileName))
			coverFile, err = file.CreateFile(coverFilePath, os.ModePerm)
			if err != nil {
				retErr = helper.HandleError(errdef.ErrUploadVideoCoverFailed, err)
				return retErr
			}
		case *videoinfov1.UploadVideoCoverReq_CoverFileData:
			// 2. Handle the data of the video cover file
			fileContent := reqData.CoverFileData
			_, err = coverFile.Write(fileContent)
			if err != nil {
				retErr = helper.HandleError(errdef.ErrUploadVideoCoverFailed, err)
				return retErr
			}
		}
	}
	err := stream.SendAndClose(&emptypb.Empty{})
	if err != nil {
		retErr = helper.HandleError(errdef.ErrUploadVideoCoverFailed, err)
		return retErr
	}

	// 3. Set the cover path of the video info record.
	absCoverPath, err := filepath.Abs(coverFilePath)
	if err != nil {
		retErr = helper.HandleError(errdef.ErrUploadVideoCoverFailed, err)
		return retErr
	}
	ctx, cancel := utilCtx.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = uc.repo.SetCoverPath(ctx, videoId, absCoverPath)
	if err != nil {
		retErr = helper.HandleError(errdef.ErrUploadVideoCoverFailed, err)
		return retErr
	}
	return nil
}

func (uc *VideoInfoUsecase) GetVideoFile(videoId int64, stream grpc.ServerStreamingServer[videoinfov1.GetVideoFileResp]) error {
	videoFilePath, err := uc.repo.GetVideoFilePathById(context.Background(), videoId)
	if err != nil {
		return helper.HandleError(errdef.ErrGetVideoFileFailed, err)
	}
	// 1. Send the video file name to the client.
	//err = stream.Send(&videoinfov1.GetVideoFileResp{
	//	VideoFile: &videoinfov1.FileResp{
	//		File: &videoinfov1.FileResp_Filename{
	//			Filename: url.PathEscape(filepath.Base(videoFilePath)),
	//		},
	//	},
	//})
	//if err != nil {
	//	return helper.HandleError(errdef.ErrGetVideoFileFailed, err)
	//}
	//
	//// 2. Send the video file data to the client.
	//videoFile, err := os.Open(videoFilePath)
	//if err != nil {
	//	return helper.HandleError(errdef.ErrGetVideoFileFailed, err)
	//}
	//defer videoFile.Close()
	//buf := make([]byte, 1*mb)
	//for {
	//	v, err := videoFile.Read(buf)
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		return helper.HandleError(errdef.ErrGetVideoFileFailed, err)
	//	}
	//	err = stream.Send(&videoinfov1.GetVideoFileResp{
	//		VideoFile: &videoinfov1.FileResp{
	//			File: &videoinfov1.FileResp_FileData{
	//				FileData: buf[:v],
	//			},
	//		},
	//	})
	//	if err != nil {
	//		return helper.HandleError(errdef.ErrGetVideoFileFailed, err)
	//	}
	//}

	filenameMsgFn := func(filename string) *videoinfov1.GetVideoFileResp {
		return &videoinfov1.GetVideoFileResp{
			VideoFile: &videoinfov1.FileResp{
				File: &videoinfov1.FileResp_Filename{
					Filename: filename,
				},
			},
		}
	}
	fileDataMsgFn := func(data []byte) *videoinfov1.GetVideoFileResp {
		return &videoinfov1.GetVideoFileResp{
			VideoFile: &videoinfov1.FileResp{
				File: &videoinfov1.FileResp_FileData{
					FileData: data,
				},
			},
		}
	}
	err = sendFileStream(stream, videoFilePath, filenameMsgFn, fileDataMsgFn)
	return nil
}

func (uc *VideoInfoUsecase) GetVideoMpd(videoId int64, stream grpc.ServerStreamingServer[videoinfov1.GetVideoMpdResp]) error {
	// TODO: 设置 context 限制超时时间，目前先不设置
	publisherId, err := uc.repo.GetPublisherIdByVideoId(utilCtx.NewBaseContext(), videoId)
	if err != nil {
		return helper.HandleError(errdef.ErrGetVideoCoverFailed, err)
	}
	mpdPath := filepath.Join(
		resourcePath,
		strconv.FormatInt(publisherId, 10),
		videoPath,
		strconv.FormatInt(videoId, 10),
		dashPath,
		"dash.mpd")

	filenameMsgFn := func(filename string) *videoinfov1.GetVideoMpdResp {
		return &videoinfov1.GetVideoMpdResp{
			MpdFile: &videoinfov1.FileResp{
				File: &videoinfov1.FileResp_Filename{
					Filename: filename,
				},
			},
		}
	}

	fileDataMsgFn := func(data []byte) *videoinfov1.GetVideoMpdResp {
		return &videoinfov1.GetVideoMpdResp{
			MpdFile: &videoinfov1.FileResp{
				File: &videoinfov1.FileResp_FileData{
					FileData: data,
				},
			},
		}
	}
	err = sendFileStream(stream, mpdPath, filenameMsgFn, fileDataMsgFn)
	if err != nil {
		return helper.HandleError(errdef.ErrGetVideoMpdFailed, err)
	}
	return nil
}

func (uc *VideoInfoUsecase) GetVideoSegment(videoSegmentPath string, stream grpc.ServerStreamingServer[videoinfov1.GetVideoSegmentResp]) error {
	filenameMsgFn := func(filename string) *videoinfov1.GetVideoSegmentResp {
		return &videoinfov1.GetVideoSegmentResp{
			SegmentFile: &videoinfov1.FileResp{
				File: &videoinfov1.FileResp_Filename{
					Filename: filepath.Base(videoSegmentPath),
				},
			},
		}
	}

	fileDataMsgFn := func(data []byte) *videoinfov1.GetVideoSegmentResp {
		return &videoinfov1.GetVideoSegmentResp{
			SegmentFile: &videoinfov1.FileResp{
				File: &videoinfov1.FileResp_FileData{
					FileData: data,
				},
			},
		}
	}
	err := sendFileStream(stream, videoSegmentPath, filenameMsgFn, fileDataMsgFn)
	if err != nil {
		return helper.HandleError(errdef.ErrGetVideoSegmentFailed, err)
	}
	return nil
}

func (uc *VideoInfoUsecase) GetVideoCover(videoId int64, stream grpc.ServerStreamingServer[videoinfov1.GetVideoCoverResp]) error {
	// TODO: 设置 context 限制超时时间，目前先不设置
	publisherId, err := uc.repo.GetPublisherIdByVideoId(utilCtx.NewBaseContext(), videoId)
	if err != nil {
		return helper.HandleError(errdef.ErrGetVideoCoverFailed, err)
	}
	coverDirPath := filepath.Join(
		resourcePath,
		strconv.FormatInt(publisherId, 10),
		videoPath,
		strconv.FormatInt(videoId, 10))
	coverFilePath, err := file.NewFileSearcher().Find(coverDirPath, coverName)

	filenameMsgFn := func(filename string) *videoinfov1.GetVideoCoverResp {
		return &videoinfov1.GetVideoCoverResp{
			CoverFile: &videoinfov1.FileResp{
				File: &videoinfov1.FileResp_Filename{
					Filename: filepath.Base(coverFilePath),
				},
			},
		}
	}

	fileDataMsgFn := func(data []byte) *videoinfov1.GetVideoCoverResp {
		return &videoinfov1.GetVideoCoverResp{
			CoverFile: &videoinfov1.FileResp{
				File: &videoinfov1.FileResp_FileData{
					FileData: data,
				},
			},
		}
	}
	err = sendFileStream(stream, coverFilePath, filenameMsgFn, fileDataMsgFn)
	if err != nil {
		return helper.HandleError(errdef.ErrGetVideoCoverFailed, err)
	}
	return nil
}

func sendFileStream[resT any, streamT grpc.ServerStreamingServer[resT]](
	stream streamT,
	filePath string,
	fileMetaMsgFn func(filename string) *resT,
	fileDataMsgFn func(data []byte) *resT,
) error {
	filename := filepath.Base(filePath)
	// 1. Send the file name to the client.
	if err := stream.Send(fileMetaMsgFn(filename)); err != nil {
		return nil
	}

	// 2. Send the file data to the client.
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 1*mb)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		err = stream.Send(fileDataMsgFn(buf[:n]))
		if err != nil {
			return err
		}
	}
	return nil
}

func protoVideoInfoToModel(info *videoinfov1.VideoMetaInfo) *model.Video {
	return &model.Video{
		Title:         info.Title,
		Description:   info.Description,
		Class:         strings.Join(info.Classes, methods.Separator),
		Hot:           info.Hot,
		Tags:          strings.Join(info.Tags, methods.Separator),
		PublisherID:   info.PublisherId,
		PublisherName: info.PublisherName,
		Duration:      info.Duration,
	}
}
