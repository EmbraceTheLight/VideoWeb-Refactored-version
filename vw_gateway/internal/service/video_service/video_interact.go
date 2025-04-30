package video

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
	interactv1 "vw_gateway/api/v1/video/video_interact"
	"vw_gateway/internal/biz/videobiz"
)

type InteractService struct {
	interactv1.UnimplementedVideoInteractServer
	interactUsecase *videobiz.InteractUsecase
	log             *log.Helper
}

func NewInteractService(interactUsecase *videobiz.InteractUsecase, logger log.Logger) *InteractService {
	return &InteractService{
		interactUsecase: interactUsecase,
		log:             log.NewHelper(logger),
	}
}

func (s *InteractService) VideoFavorite(ctx context.Context, req *interactv1.VideoFavoriteReq) (*interactv1.VideoFavoriteResp, error) {
	curUserId := ctx.Value(userIdKey).(int64)
	isFavorited, err := s.interactUsecase.Favorite(ctx, curUserId, req.VideoId, req.FavoriteId)
	if err != nil {
		return nil, err
	}

	var msg string
	// isFavorited is true means user has favorited this video before, so this is a cancel favorite request
	if isFavorited {
		msg = "取消收藏成功"
	} else {
		msg = "收藏成功"
	}
	return &interactv1.VideoFavoriteResp{
		StatusCode: http.StatusOK,
		Message:    msg,
	}, nil
}

func (s *InteractService) VideoUpvoteBarrage(ctx context.Context, req *interactv1.VideoUpvoteBarrageReq) (*interactv1.VideoUpvoteBarrageResp, error) {
	curUserId := ctx.Value(userIdKey).(int64)
	err := s.interactUsecase.UpvoteBarrage(ctx, curUserId, req.BarrageId)
	if err != nil {
		return nil, err
	}
	return &interactv1.VideoUpvoteBarrageResp{
		StatusCode: http.StatusOK,
		Message:    "点赞弹幕成功",
	}, nil

}

func (s *InteractService) VideoSendBarrage(ctx context.Context, req *interactv1.VideoSendBarrageReq) (*interactv1.VideoSendBarrageResp, error) {
	curUserId := ctx.Value(userIdKey).(int64)
	err := s.interactUsecase.SendBarrage(ctx, curUserId, req)
	if err != nil {
		return nil, err
	}
	return &interactv1.VideoSendBarrageResp{
		StatusCode: http.StatusOK,
		Message:    "发送弹幕成功",
	}, nil
}

func (s *InteractService) VideoShare(ctx context.Context, req *interactv1.VideoShareReq) (*interactv1.VideoShareResp, error) {
	curUserId := ctx.Value(userIdKey).(int64)
	url, err := s.interactUsecase.Share(ctx, curUserId, req.VideoId)
	if err != nil {
		return nil, err
	}
	return &interactv1.VideoShareResp{
		Url:        url,
		StatusCode: http.StatusOK,
		Message:    "分享链接已生成",
	}, nil

}

func (s *InteractService) VideoThrowShells(ctx context.Context, req *interactv1.VideoThrowShellsReq) (*interactv1.VideoThrowShellsResp, error) {
	curUserId := ctx.Value(userIdKey).(int64)
	err := s.interactUsecase.ThrowShells(ctx, curUserId, req.VideoId, req.Shells)
	if err != nil {
		return nil, err
	}
	return &interactv1.VideoThrowShellsResp{
		StatusCode: http.StatusOK,
		Message:    "投币成功",
	}, nil

}

func (s *InteractService) VideoUpvote(ctx context.Context, req *interactv1.VideoUpvoteReq) (*interactv1.VideoUpvoteResp, error) {
	curUserId := ctx.Value(userIdKey).(int64)
	isUpvoted, err := s.interactUsecase.Upvote(ctx, curUserId, req.VideoId)
	if err != nil {
		return nil, err
	}
	var msg string
	// isUpvoted is true means user has upvoted this video before, so this is a cancel upvote request
	if isUpvoted {
		msg = "取消点赞成功"
	} else {
		msg = "点赞成功"
	}
	return &interactv1.VideoUpvoteResp{
		StatusCode: http.StatusOK,
		Message:    msg,
	}, nil
}
