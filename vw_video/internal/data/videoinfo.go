package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vw_video/internal/biz"
	"vw_video/internal/data/dal/model"
	"vw_video/internal/data/dal/query"
)

type videoInfoRepo struct {
	data   *Data
	logger *log.Helper
}

func NewVideoInfoRepo(data *Data, logger log.Logger) biz.VideoInfoRepo {
	return &videoInfoRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

func (v *videoInfoRepo) AddVideoHot(ctx context.Context, videoId, hot int64) error {

	//TODO implement me
	panic("implement me")
}

func (v *videoInfoRepo) SetCoverPath(ctx context.Context, videoId int64, coverPath string) error {
	video := getQuery(ctx).Video
	videoDo, _, err := addVideoModel(ctx, videoId)
	if err != nil {
		return err
	}

	_, err = videoDo.Where(video.VideoID.Eq(videoId)).Updates(&model.Video{
		CoverPath: coverPath,
	})

	return err
}

func (v *videoInfoRepo) UpdateVideoDurationSizeInfo(ctx context.Context, videoId int64, duration string, size int64) error {
	video := getQuery(ctx).Video
	videoDo, _, err := addVideoModel(ctx, videoId)
	if err != nil {
		return err
	}

	_, err = videoDo.Where(video.VideoID.Eq(videoId)).Updates(&model.Video{
		Duration: duration,
		Size:     size,
	})
	return err
}

func (v *videoInfoRepo) UpdateVideoFilePath(ctx context.Context, videoId int64, videoFilePath string) error {
	video := getQuery(ctx).Video
	videoDo, _, err := addVideoModel(ctx, videoId)
	if err != nil {
		return err
	}

	_, err = videoDo.Where(video.VideoID.Eq(videoId)).Updates(&model.Video{
		VideoPath: videoFilePath,
	})
	return err
}

func (v *videoInfoRepo) CreateBasicVideoInfo(ctx context.Context, videoInfo *model.Video) error {
	video := getQuery(ctx).Video
	videoDo := video.WithContext(ctx)
	return videoDo.Create(videoInfo)

}

func (v *videoInfoRepo) GetVideoInfoById(ctx context.Context, videoId int64) (*model.Video, error) {
	video := query.Video
	videoDo, _, err := addVideoModel(ctx, videoId)
	if err != nil {
		return nil, err
	}
	return videoDo.Where(video.VideoID.Eq(videoId)).First()
}

func (v *videoInfoRepo) GetVideoListByClass(ctx context.Context, class []string, pageNum int32, pageSize int32) ([]*model.Video, error) {
	video := getQuery(ctx).Video
	videoDo := video.WithContext(ctx)
	res, _, err := videoDo.
		Where(video.Class.In(class...)).
		Order(video.Hot.Desc()).
		FindByPage(int((pageNum-1)*pageSize), int(pageSize))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (v *videoInfoRepo) GetVideoFilePathById(ctx context.Context, videoId int64) (string, error) {
	video := getQuery(ctx).Video
	videoDo := video.WithContext(ctx)
	res, err := videoDo.Where(video.VideoID.Eq(videoId)).Select(video.VideoPath).First()
	if err != nil {
		return "", err
	}
	return res.VideoPath, nil
}

func (v *videoInfoRepo) GetPublisherIdByVideoId(ctx context.Context, videoId int64) (publisherId int64, err error) {
	video := getQuery(ctx).Video
	videoDo := video.WithContext(ctx)
	res, err := videoDo.Where(video.VideoID.Eq(videoId)).Select(video.PublisherID).First()
	if err != nil {
		return 0, err
	}
	return res.PublisherID, nil
}
