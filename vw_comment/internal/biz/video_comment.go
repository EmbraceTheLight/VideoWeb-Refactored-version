package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"util/dbutil/mysqlutil"
	"util/iputil"
	vcommentv1 "vw_comment/api/v1/video_comment"
	"vw_comment/internal/data/dal/model"
)

type VideoCommentRepo interface {
	GetRootComment(ctx context.Context, videoId int64, sortBy, order string, pageNum, pageSize int32) ([]*model.Comment, error)
	GetReplyComments(ctx context.Context, parentId int64, sortBy, order string, pageNum, pageSize int32) ([]*model.Comment, error)
	GetRepliesByParentID(ctx context.Context, parentId int64, sortBy, order string, pageNum, pageSize int32) ([]*model.Comment, error)
	GetRootIDByParentID(ctx context.Context, videoId, parentId int64) (int64, error)
	InsertComment(ctx context.Context, comment *model.Comment) (int64, error)
	IncrementCommentCntReplies(ctx context.Context, commentId int64) error
	IncrementCommentCntUpvote(ctx context.Context, commentId int64) error
	DecrementCommentCntUpvote(ctx context.Context, commentId int64) error
	AddVideoHotByComment(ctx context.Context, videoId int64) error

	/*MongoDB related methods*/
	// CommentSetUpvoted Insert a document into comment_upvoted table.
	CommentSetUpvoted(ctx context.Context, commentId int64, userId int64) error

	// CommentCancelUpvoted Delete a document from comment_upvoted table.
	CommentCancelUpvoted(ctx context.Context, commentId int64, userId int64) error

	// CheckIfUserUpvotedComment Check if a user has upvoted a comment.
	CheckIfUserUpvotedComment(ctx context.Context, commentId int64, userId int64) (bool, error)
}

type VideoCommentUsecase struct {
	repo VideoCommentRepo
	tx   mysqlutil.Transaction
	log  *log.Helper
}

func NewVideoCommentUsecase(repo VideoCommentRepo, tx mysqlutil.Transaction, logger log.Logger) *VideoCommentUsecase {
	return &VideoCommentUsecase{
		repo: repo,
		tx:   tx,
		log:  log.NewHelper(logger),
	}
}

func (u *VideoCommentUsecase) CreateComment(ctx context.Context, videoId, parentId, publisherId int64, content string, ipAddr string) (int64, error) {
	country, city, err := iputil.GetCountryAndCity(ipAddr)
	if err != nil {
		return 0, err
	}
	if country == "" {
		country = "未知地区"
	}

	var newCommentId int64
	err = u.tx.WithTx(ctx, func(ctx context.Context) error {
		// 1. Get root id by parent id.
		rootId, err := u.repo.GetRootIDByParentID(ctx, videoId, parentId)
		if err != nil {
			return err
		}

		// 2. Insert comment.
		comment := &model.Comment{
			VideoID:     videoId,
			ParentID:    parentId,
			RootID:      rootId,
			PublisherID: publisherId,
			Content:     content,
			IPLocation:  country + " " + city,
		}
		newCommentId, err = u.repo.InsertComment(ctx, comment)
		if err != nil {
			return err
		}

		// 3. Increment cnt_replies of parent comment.
		err = u.repo.IncrementCommentCntReplies(ctx, parentId)
		if err != nil {
			return err
		}

		// 4. Increment hot of video.

		return nil
	})
	if err != nil {
		return -1, err
	}

	return newCommentId, nil
}

func (u *VideoCommentUsecase) GetCommentList(ctx context.Context, videoId int64, sortBy, order string, page, pageSize int32) ([]*vcommentv1.CommentInfo, error) {
	// 1. Get root comments. Which is the comments to the video.
	comments, err := u.repo.GetRootComment(ctx, videoId, sortBy, order, page, pageSize)
	if err != nil {
		return nil, err
	}

	// 2. Padding resp with root comments.
	result := paddingResp(comments...)
	return result, nil
}

func (u *VideoCommentUsecase) GetCommentReplies(ctx context.Context, commentId int64, sortBy, order string, pageNum, pageSize int32) ([]*vcommentv1.CommentInfo, error) {
	// 1. Get all replies of this comment.
	replies, err := u.repo.GetRepliesByParentID(ctx, commentId, sortBy, order, pageNum, pageSize)
	if err != nil {
		return nil, err
	}

	repliesLen := int32(len(replies))
	if repliesLen < (pageNum-1)*pageSize {
		return nil, nil
	}

	// 2. Padding resp with replies.
	result := paddingResp(replies[(pageNum-1)*pageSize : min(repliesLen, pageNum*pageSize)]...)
	return result, nil
}

func (u *VideoCommentUsecase) UpvoteComment(ctx context.Context, commentId int64, userId int64, isUpvoted bool) error {
	err := u.tx.WithTx(ctx, func(ctx context.Context) error {
		var err error
		// 1. Update comment's cnt_upvote field.
		// 2. Insert or delete a document into user_comment_upvoted table.
		if isUpvoted {
			err = u.repo.IncrementCommentCntUpvote(ctx, commentId)
			if err != nil {
				return err
			}

			err = u.repo.CommentSetUpvoted(ctx, commentId, userId)
			if err != nil {
				return err
			}
		} else {
			err = u.repo.DecrementCommentCntUpvote(ctx, commentId)
			if err != nil {
				return err
			}

			err = u.repo.CommentCancelUpvoted(ctx, commentId, userId)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *VideoCommentUsecase) CheckIfUserUpvotedComment(ctx context.Context, commentId int64, userId int64) (bool, error) {
	return u.repo.CheckIfUserUpvotedComment(ctx, commentId, userId)
}

func paddingResp(comments ...*model.Comment) []*vcommentv1.CommentInfo {
	result := make([]*vcommentv1.CommentInfo, len(comments))
	for i, comment := range comments {
		result[i] = &vcommentv1.CommentInfo{
			CommentId:   comment.CommentID,
			PublisherId: comment.PublisherID,
			RootId:      comment.RootID,
			Content:     comment.Content,
			CreatedAt:   timestamppb.New(comment.CreatedAt),
			UpvoteCount: comment.CntUpvote,
			IpAddress:   comment.IPLocation,
			CntReplies:  uint32(comment.CntReplies),
		}
	}
	return result
}
