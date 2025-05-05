package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	vcommentv1 "vw_comment/api/v1/video_comment"
	"vw_comment/internal/biz"
	"vw_comment/internal/pkg/ecode/errdef"
)

type VideoCommentService struct {
	vcommentv1.UnimplementedVideoCommentServer
	vcUsecase *biz.VideoCommentUsecase
	log       *log.Helper
}

func NewVideoCommentService(vcUsecase *biz.VideoCommentUsecase, logger log.Logger) *VideoCommentService {
	return &VideoCommentService{
		vcUsecase: vcUsecase,
		log:       log.NewHelper(logger),
	}
}

func (s *VideoCommentService) PublishComment(ctx context.Context, req *vcommentv1.PublishCommentReq) (*vcommentv1.PublishCommentResp, error) {
	ipAddr, ok := ctx.Value("client_ip").(string)
	if !ok {
		return nil, errdef.ErrInternal
	}
	commentId, err := s.vcUsecase.CreateComment(ctx, req.VideoId, req.ParentId, req.PublisherId, req.Content, ipAddr)
	if err != nil {
		return nil, err
	}
	return &vcommentv1.PublishCommentResp{CommentId: commentId}, nil
}

func (s *VideoCommentService) GetCommentList(ctx context.Context, req *vcommentv1.GetCommentListReq) (*vcommentv1.GetCommentListResp, error) {
	commentList, err := s.vcUsecase.GetCommentList(ctx, req.VideoId, req.SortBy, req.Order, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &vcommentv1.GetCommentListResp{Comments: commentList, TotalCount: int32(len(commentList))}, nil
}

func (s *VideoCommentService) GetCommentReplies(ctx context.Context, req *vcommentv1.GetCommentRepliesReq) (*vcommentv1.GetCommentRepliesResp, error) {
	replies, err := s.vcUsecase.GetCommentReplies(ctx, req.CommentId, req.SortBy, req.Order, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &vcommentv1.GetCommentRepliesResp{Replies: replies}, nil
}

func (s *VideoCommentService) UpvoteComment(ctx context.Context, req *vcommentv1.UpvoteCommentReq) (*emptypb.Empty, error) {
	err := s.vcUsecase.UpvoteComment(ctx, req.CommentId, req.UserId, req.IsUpvoted)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *VideoCommentService) CheckIfUserUpvotedComment(ctx context.Context, req *vcommentv1.CheckIfUserUpvotedCommentReq) (*vcommentv1.CheckIfUserUpvotedCommentResp, error) {
	isUpvoted, err := s.vcUsecase.CheckIfUserUpvotedComment(ctx, req.CommentId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &vcommentv1.CheckIfUserUpvotedCommentResp{IsUpvoted: isUpvoted}, nil
}
