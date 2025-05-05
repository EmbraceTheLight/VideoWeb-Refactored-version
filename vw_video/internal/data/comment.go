package data

import (
	"context"
	stderr "errors"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gen/field"
	"vw_video/internal/biz"
	"vw_video/internal/data/dal/model"
	"vw_video/internal/data/dal/query"
)

type videoCommentRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoCommentRepo(data *Data, logger log.Logger) biz.VideoCommentRepo {
	return &videoCommentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

const separator = "::"

var orderMap map[string]field.Expr

func initOrderMap() {
	orderMap = map[string]field.Expr{
		"created_at" + separator + "asc":  query.Comment.CreatedAt.Asc(),
		"created_at" + separator + "desc": query.Comment.CreatedAt.Desc(),
		"cnt_upvote" + separator + "desc": query.Comment.CntUpvote.Desc(),
	}
}

func (v *videoCommentRepo) GetRootComment(ctx context.Context, videoId int64, sortBy, order string, pageNum, pageSize int32) ([]*model.Comment, error) {
	comment := getQuery(ctx).Comment
	commentDo := comment.WithContext(ctx)
	ret, err := commentDo.
		Where(comment.VideoID.Eq(videoId), comment.ParentID.Eq(-1)).
		Order(orderMap[sortBy+separator+order]).
		Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize)).
		Find()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (v *videoCommentRepo) GetReplyComments(ctx context.Context, parentId int64, sortBy, order string, pageNum, pageSize int32) ([]*model.Comment, error) {
	comment := getQuery(ctx).Comment
	commentDo := comment.WithContext(ctx)
	ret, err := commentDo.
		Where(comment.ParentID.Eq(parentId)).
		Order(orderMap[sortBy+separator+order]).
		Offset(int((pageNum - 1) * pageSize)).
		Limit(int(pageSize)).
		Find()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (v *videoCommentRepo) GetRepliesByParentID(ctx context.Context, parentId int64, sortBy, order string, pageNum, pageSize int32) ([]*model.Comment, error) {
	comment := getQuery(ctx).Comment
	commentDo := comment.WithContext(ctx)

	ret, err := commentDo.
		Debug().
		Where(comment.ParentID.Eq(parentId)).
		Order(orderMap[sortBy+separator+order]).
		Offset(int((pageNum - 1) * pageSize)).
		Limit(int(pageSize)).
		Find()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (v *videoCommentRepo) InsertComment(ctx context.Context, newComment *model.Comment) (int64, error) {
	comment := getQuery(ctx).Comment
	commentDo := comment.WithContext(ctx)
	err := commentDo.Create(newComment)
	if err != nil {
		return 0, err
	}

	// If the comment is a root comment, its root_id is -1 now.
	// We need to update root_id as the same as the comment_id.
	if newComment.ParentID == -1 {
		_, err = commentDo.Where(comment.CommentID.Eq(newComment.CommentID)).Updates(map[string]any{
			"root_id": newComment.CommentID,
		})
		if err != nil {
			return 0, err
		}
	}
	return newComment.CommentID, nil
}

func (v *videoCommentRepo) GetRootIDByParentID(ctx context.Context, videoId, parentId int64) (int64, error) {
	// parent_id == -1 indicates that the comment is a root comment., so its root_id is -1.
	if parentId == -1 {
		return -1, nil
	}

	comment := getQuery(ctx).Comment
	commentDo := comment.WithContext(ctx)
	ret, err := commentDo.Select(comment.RootID).Where(comment.VideoID.Eq(videoId), comment.CommentID.Eq(parentId)).First()
	if err != nil {
		return 0, err
	}

	return ret.RootID, nil
}

func (v *videoCommentRepo) IncrementCommentCntReplies(ctx context.Context, commentId int64) error {
	// This case indicated that the child comment is a root comment, so we don't need to increment the cnt_replies field.
	if commentId < 0 {
		return nil
	}
	comment := query.Comment
	commentDo, commentModel, err := addCommentModel(ctx, commentId)
	if err != nil {
		return err
	}
	_, err = commentDo.Where(comment.CommentID.Eq(commentId)).Updates(
		&model.Comment{
			CntReplies: commentModel.CntReplies + 1,
		})
	return nil
}

func (v *videoCommentRepo) IncrementCommentCntUpvote(ctx context.Context, commentId int64) error {
	comment := getQuery(ctx).Comment
	commentDo, commentModel, err := addCommentModel(ctx, commentId)
	if err != nil {
		return err
	}
	_, err = commentDo.
		Where(comment.CommentID.Eq(commentId)).
		Updates(&model.Comment{
			CntUpvote: commentModel.CntUpvote + 1,
		})
	return err
}

func (v *videoCommentRepo) DecrementCommentCntUpvote(ctx context.Context, commentId int64) error {
	comment := getQuery(ctx).Comment
	commentDo, commentModel, err := addCommentModel(ctx, commentId)
	if err != nil {
		return err
	}
	_, err = commentDo.
		Where(comment.CommentID.Eq(commentId)).
		Updates(map[string]any{
			"cnt_upvote": commentModel.CntUpvote - 1,
		})
	return err
}

func (v *videoCommentRepo) AddVideoHotAndCntCommentsByComment(ctx context.Context, videoId int64) error {
	video := query.Video
	videoDo, videoModel, err := addVideoModel(ctx, videoId)
	if err != nil {
		return err
	}
	_, err = videoDo.Where(video.VideoID.Eq(videoId)).Updates(
		map[string]any{
			"cnt_comments": videoModel.CntComments + 1,
			"hot":          videoModel.Hot + HotEachReply,
		})
	return nil
}

func (v *videoCommentRepo) AddVideoHotAndCntCommentsByReply(ctx context.Context, videoId int64) error {
	video := query.Video
	videoDo, videoModel, err := addVideoModel(ctx, videoId)
	if err != nil {
		return err
	}
	_, err = videoDo.Where(video.VideoID.Eq(videoId)).Updates(
		map[string]any{
			"cnt_comments": videoModel.CntComments + 1,
			"hot":          videoModel.Hot + HotEachReply,
		})
	return nil
}

func (v *videoCommentRepo) GetCommentInfo(ctx context.Context, commentId int64) (*model.Comment, error) {
	comment := getQuery(ctx).Comment
	commentDo := comment.WithContext(ctx)
	c, err := commentDo.Where(comment.CommentID.Eq(commentId)).First()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (v *videoCommentRepo) CommentSetUpvoted(ctx context.Context, commentId int64, userId int64) error {
	return v.data.mongo.InsertOne(ctx, uc_status, &model.UserCommentUpvoted{
		UserID:    userId,
		CommentID: commentId,
	})
}

func (v *videoCommentRepo) CommentCancelUpvoted(ctx context.Context, commentId int64, userId int64) error {
	return v.data.mongo.DeleteOne(ctx, &model.UserCommentUpvoted{
		UserID:    userId,
		CommentID: commentId,
	})
}

func (v *videoCommentRepo) CheckIfUserUpvotedComment(ctx context.Context, commentId int64, userId int64) (bool, error) {
	filter := &model.UserCommentUpvoted{
		CommentID: commentId,
		UserID:    userId,
	}
	_, err := v.data.mongo.FindOne(ctx, uc_status, filter)
	if err != nil {
		if stderr.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
