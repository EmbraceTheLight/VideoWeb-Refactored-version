package videobiz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"util/iputil"
	vcommentv1 "vw_gateway/api/v1/video/video_comment"
	"vw_gateway/internal/pkg/ecode/errdef"
)

type VideoCommentRepo interface {
	PublishVideoComment(ctx context.Context, publisherId int64, videoId int64, parentId int64, content string, ipAddr string) (int64, error)
	GetCommentList(ctx context.Context, videoId int64, pageNum int32, pageSize int32, sortBy string, order string) (*vcommentv1.GetCommentListResp, error)
	GetCommentReplies(ctx context.Context, commentId int64, pageNum int32, pageSize int32, sortBy string, order string) (*vcommentv1.GetCommentRepliesResp, error)
	UpvoteComment(ctx context.Context, commentId int64, userId, publisherId int64, isUpvoted bool) error
	CheckIfUserUpvotedComment(ctx context.Context, commentId int64, userId int64) (bool, error)
	GetPublisherId(ctx context.Context, commentId int64) (int64, error)
}

type VideoCommentUsecase struct {
	repo VideoCommentRepo
	log  *log.Helper
}

func NewVideoCommentUsecase(repo VideoCommentRepo, logger log.Logger) *VideoCommentUsecase {
	return &VideoCommentUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (u *VideoCommentUsecase) CommentVideo(ctx context.Context, publisherId int64, videoId int64, parentId int64, content string) (int64, error) {
	httpRequest, ok := khttp.RequestFromServerContext(ctx)
	if !ok {
		return 0, errdef.INTERNAL_ERROR
	}

	ipAddr, err := iputil.GetClientIp(httpRequest)
	if err != nil {
		return 0, errdef.INTERNAL_ERROR
	}

	return u.repo.PublishVideoComment(ctx, publisherId, videoId, parentId, content, ipAddr)
}

func (u *VideoCommentUsecase) GetCommentList(ctx context.Context, videoId int64, pageNum int32, pageSize int32, sortBy string, order string) (*vcommentv1.GetCommentListResp, error) {
	resp, err := u.repo.GetCommentList(ctx, videoId, pageNum, pageSize, sortBy, order)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *VideoCommentUsecase) GetCommentReplies(ctx context.Context, commentId int64, pageNum int32, pageSize int32, sortBy string, order string) (*vcommentv1.GetCommentRepliesResp, error) {
	resp, err := u.repo.GetCommentReplies(ctx, commentId, pageNum, pageSize, sortBy, order)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *VideoCommentUsecase) UpvoteComment(ctx context.Context, commentId int64, userId int64) (bool, error) {
	isUpvoted, err := u.repo.CheckIfUserUpvotedComment(ctx, commentId, userId)
	if err != nil {
		return false, err
	}

	publisherId, err := u.repo.GetPublisherId(ctx, commentId)
	if err != nil {
		return false, err
	}

	return isUpvoted, u.repo.UpvoteComment(ctx, commentId, userId, publisherId, isUpvoted)
}
