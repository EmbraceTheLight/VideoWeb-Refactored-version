package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	commentv1 "vw_video/api/v1/comment"
	"vw_video/internal/biz"
)

type VideoCommentService struct {
	commentv1.UnimplementedVideoCommentServer
	vcUsecase *biz.VideoCommentUsecase
	log       *log.Helper
}

func NewVideoCommentService(vcUsecase *biz.VideoCommentUsecase, logger log.Logger) *VideoCommentService {
	return &VideoCommentService{
		vcUsecase: vcUsecase,
		log:       log.NewHelper(logger),
	}
}

func (s *VideoCommentService) PublishComment(ctx context.Context, req *commentv1.PublishCommentReq) (*commentv1.PublishCommentResp, error) {
	commentId, err := s.vcUsecase.CreateComment(ctx, req.VideoId, req.ParentId, req.PublisherId, req.Content, req.IpAddr)
	if err != nil {
		return nil, err
	}
	return &commentv1.PublishCommentResp{CommentId: commentId}, nil
}

func (s *VideoCommentService) GetCommentList(ctx context.Context, req *commentv1.GetCommentListReq) (*commentv1.GetCommentListResp, error) {
	commentList, err := s.vcUsecase.GetCommentList(ctx, req.VideoId, req.SortBy, req.Order, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &commentv1.GetCommentListResp{Comments: commentList, TotalCount: int32(len(commentList))}, nil
}

func (s *VideoCommentService) GetCommentReplies(ctx context.Context, req *commentv1.GetCommentRepliesReq) (*commentv1.GetCommentRepliesResp, error) {
	replies, err := s.vcUsecase.GetCommentReplies(ctx, req.CommentId, req.SortBy, req.Order, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &commentv1.GetCommentRepliesResp{Replies: replies}, nil
}

func (s *VideoCommentService) UpvoteComment(ctx context.Context, req *commentv1.UpvoteCommentReq) (*emptypb.Empty, error) {
	err := s.vcUsecase.UpvoteComment(ctx, req.CommentId, req.UserId, req.IsUpvoted)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *VideoCommentService) UpvoteCommentRevert(ctx context.Context, req *commentv1.UpvoteCommentReq) (*emptypb.Empty, error) {
	req.IsUpvoted = !req.IsUpvoted
	return s.UpvoteComment(ctx, req)
}

func (s *VideoCommentService) CheckIfUserUpvotedComment(ctx context.Context, req *commentv1.CheckIfUserUpvotedCommentReq) (*commentv1.CheckIfUserUpvotedCommentResp, error) {
	isUpvoted, err := s.vcUsecase.CheckIfUserUpvotedComment(ctx, req.CommentId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &commentv1.CheckIfUserUpvotedCommentResp{IsUpvoted: isUpvoted}, nil
}

func (s *VideoCommentService) GetPublisherId(ctx context.Context, req *commentv1.GetPublisherIdReq) (*commentv1.GetPublisherIdResp, error) {
	id, err := s.vcUsecase.GetPublisherId(ctx, req.CommentId)
	if err != nil {
		return nil, err
	}
	return &commentv1.GetPublisherIdResp{PublisherId: id}, nil
}
