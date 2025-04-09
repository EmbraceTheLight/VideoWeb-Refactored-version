package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
	videoinfov1 "vw_video/api/v1/videoinfo"
	"vw_video/internal/biz"
	"vw_video/internal/data/dal/model"
)

type VideoInfoService struct {
	videoinfov1.UnimplementedVideoInfoServer
	info   *biz.VideoInfoUsecase
	logger *log.Helper
}

func NewVideoInfo(info *biz.VideoInfoUsecase, logger log.Logger) *VideoInfoService {
	return &VideoInfoService{
		info:   info,
		logger: log.NewHelper(logger),
	}
}

func (s *VideoInfoService) GetVideoInfo(ctx context.Context, req *videoinfov1.GetVideoInfoReq) (*videoinfov1.GetVideoInfoResp, error) {
	videoDetail, err := s.info.GetVideoInfoById(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	return &videoinfov1.GetVideoInfoResp{
		VideoDetail: modelToResp(videoDetail),
	}, nil
}

func (s *VideoInfoService) GetVideoList(ctx context.Context, req *videoinfov1.GetVideoListReq) (*videoinfov1.GetVideoListResp, error) {
	spanCtx := trace.SpanContextFromContext(ctx)
	fmt.Println("[video grpc service GetVideoList] Span ID:", spanCtx.SpanID().String())
	fmt.Println("[video grpc service GetVideoList] Trace ID:", spanCtx.TraceID().String())
	videoDetails, err := s.info.GetVideoListByClass(ctx, req.Class, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	resp := &videoinfov1.GetVideoListResp{
		VideoDetails: make([]*videoinfov1.VideoMetaInfo, len(videoDetails)),
	}
	for i, videoDetail := range videoDetails {
		resp.VideoDetails[i] = modelToResp(videoDetail)
	}
	return resp, nil
}

func (s *VideoInfoService) UploadVideoInfo(ctx context.Context, req *videoinfov1.UploadVideoInfoReq) (*emptypb.Empty, error) {
	err := s.info.UploadVideoInfo(ctx, req.VideoInfo)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *VideoInfoService) UploadVideoFile(stream grpc.ClientStreamingServer[videoinfov1.UploadVideoFileReq, emptypb.Empty]) error {
	return s.info.UploadVideoFile(stream)
}

func (s *VideoInfoService) UploadVideoCover(stream grpc.ClientStreamingServer[videoinfov1.UploadVideoCoverReq, emptypb.Empty]) error {
	return s.info.UploadVideoCover(stream)
}

func (s *VideoInfoService) GetVideoFile(req *videoinfov1.GetVideoFileReq, stream grpc.ServerStreamingServer[videoinfov1.GetVideoFileResp]) error {
	return s.info.GetVideoFile(req.VideoId, stream)
}

func (s *VideoInfoService) GetVideoMpd(req *videoinfov1.GetVideoMpdReq, stream grpc.ServerStreamingServer[videoinfov1.GetVideoMpdResp]) error {
	return s.info.GetVideoMpd(req.VideoId, stream)
}

func (s *VideoInfoService) GetVideoSegments(req *videoinfov1.GetVideoSegmentReq, stream grpc.ServerStreamingServer[videoinfov1.GetVideoSegmentResp]) error {
	return s.info.GetVideoSegment(req.VideoSegmentPath, stream)
}

func (s *VideoInfoService) GetVideoCover(req *videoinfov1.GetVideoCoverReq, stream grpc.ServerStreamingServer[videoinfov1.GetVideoCoverResp]) error {
	return s.info.GetVideoCover(req.VideoId, stream)
}

func modelToResp(videoInfo *model.Video) *videoinfov1.VideoMetaInfo {
	return &videoinfov1.VideoMetaInfo{
		VideoId:       videoInfo.VideoID,
		PublisherId:   videoInfo.PublisherID,
		PublisherName: videoInfo.PublisherName,
		Title:         videoInfo.Title,
		Description:   videoInfo.Description,
		VideoPath:     videoInfo.VideoPath,
		Classes:       strings.Split(videoInfo.Class, ","),
		Tags:          strings.Split(videoInfo.Tags, ","),
		Hot:           videoInfo.Hot,
		Records: &videoinfov1.VideoMetaInfo_Records{
			CntBarrages:  uint32(videoInfo.CntBarrages),
			CntShares:    uint32(videoInfo.CntShares),
			CntViewed:    uint32(videoInfo.CntViewed),
			CntFavorites: uint32(videoInfo.CntFavorited),
		},
		Duration:  videoInfo.Duration,
		CoverPath: videoInfo.CoverPath,
	}
}
