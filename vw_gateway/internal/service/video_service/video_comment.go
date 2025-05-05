package video

import (
	"context"
	stderr "errors"
	"github.com/go-kratos/kratos/v2/log"
	"util/helper"
	vcommentv1 "vw_gateway/api/v1/video/video_comment"
	"vw_gateway/internal/biz/videobiz"
	"vw_gateway/internal/pkg/ecode/errdef"
)

type CommentService struct {
	vcommentv1.UnimplementedVideoCommentServer
	vcUsecase *videobiz.VideoCommentUsecase
	log       *log.Helper
}

func NewVideoCommentService(vcUsecase *videobiz.VideoCommentUsecase, logger log.Logger) *CommentService {
	return &CommentService{
		vcUsecase: vcUsecase,
		log:       log.NewHelper(logger),
	}
}

func (s *CommentService) CommentVideo(ctx context.Context, req *vcommentv1.CommentVideoReq) (*vcommentv1.CommentVideoResp, error) {
	publisherId, ok := ctx.Value(userIdKey).(int64)
	if !ok {
		return nil, helper.HandleError(errdef.INTERNAL_ERROR, stderr.New("用户id获取失败"))
	}

	commentId, err := s.vcUsecase.CommentVideo(ctx, publisherId, req.VideoId, req.ParentId, req.Content)
	if err != nil {
		return nil, err
	}
	return &vcommentv1.CommentVideoResp{CommentId: commentId}, nil
}

func (s *CommentService) GetCommentList(ctx context.Context, req *vcommentv1.GetCommentListReq) (*vcommentv1.GetCommentListResp, error) {
	resp, err := s.vcUsecase.GetCommentList(ctx, req.VideoId, req.PageNum, req.PageSize, req.SortBy, req.Order)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *CommentService) GetCommentReplies(ctx context.Context, req *vcommentv1.GetCommentRepliesReq) (*vcommentv1.GetCommentRepliesResp, error) {
	resp, err := s.vcUsecase.GetCommentReplies(ctx, req.CommentId, req.PageNum, req.PageSize, req.SortBy, req.Order)
	if err != nil {
		return nil, err
	}
	return &vcommentv1.GetCommentRepliesResp{Replies: resp.Replies}, nil
}

func (s *CommentService) UpvoteComment(ctx context.Context, req *vcommentv1.UpvoteCommentReq) (*vcommentv1.UpvoteCommentResp, error) {
	userId, ok := ctx.Value(userIdKey).(int64)
	if !ok {
		return nil, helper.HandleError(errdef.INTERNAL_ERROR, stderr.New("用户id获取失败"))
	}
	isUpvoted, err := s.vcUsecase.UpvoteComment(ctx, req.CommentId, userId)
	if err != nil {
		return nil, err
	}

	var msg string
	if isUpvoted {
		msg = "评论取消点赞成功"
	} else {
		msg = "评论点赞成功"
	}
	return &vcommentv1.UpvoteCommentResp{
		StatusCode: 200,
		Message:    msg,
	}, nil
}
