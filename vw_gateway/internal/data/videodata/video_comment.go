package videodata

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"util/resolver"
	vcommentv1 "vw_gateway/api/v1/video/video_comment"
	"vw_gateway/internal/biz/videobiz"
	userinfoGRPC "vw_user/api/v1/userinfo"
	vcommentGRPC "vw_video/api/v1/comment"
)

type videoCommentRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoCommentRepo(data *Data, logger log.Logger) videobiz.VideoCommentRepo {
	return &videoCommentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (v *videoCommentRepo) PublishVideoComment(ctx context.Context, publisherId int64, videoId int64, parentId int64, content string, ipAddr string) (int64, error) {
	resp, err := v.data.videoCommentClient.PublishComment(ctx, &vcommentGRPC.PublishCommentReq{
		VideoId:     videoId,
		ParentId:    parentId,
		PublisherId: publisherId,
		Content:     content,
		IpAddr:      ipAddr,
	})
	if err != nil {
		return 0, err
	}
	return resp.CommentId, nil
}

func (v *videoCommentRepo) GetCommentList(ctx context.Context, videoId int64, pageNum int32, pageSize int32, sortBy string, order string) (*vcommentv1.GetCommentListResp, error) {
	resp, err := v.data.videoCommentClient.GetCommentList(ctx, &vcommentGRPC.GetCommentListReq{
		VideoId:  videoId,
		PageNum:  pageNum,
		PageSize: pageSize,
		SortBy:   sortBy,
		Order:    order,
	})
	if err != nil {
		return nil, err
	}

	return &vcommentv1.GetCommentListResp{
		Comments:   convert(resp.Comments),
		TotalCount: resp.TotalCount,
	}, nil
}

func (v *videoCommentRepo) GetCommentReplies(ctx context.Context, commentId int64, pageNum int32, pageSize int32, sortBy string, order string) (*vcommentv1.GetCommentRepliesResp, error) {
	resp, err := v.data.videoCommentClient.GetCommentReplies(ctx, &vcommentGRPC.GetCommentRepliesReq{
		CommentId: commentId,
		PageNum:   pageNum,
		PageSize:  pageSize,
		SortBy:    sortBy,
		Order:     order,
	})
	if err != nil {
		return nil, err
	}

	return &vcommentv1.GetCommentRepliesResp{
		Replies: convert(resp.Replies),
	}, nil
}

func (v *videoCommentRepo) UpvoteComment(ctx context.Context, commentId int64, userId, publisherId int64, isUpvoted bool) error {
	videoAddr, err := resolver.GetRandomAddr(ctx, videoService)
	if err != nil {
		return err
	}

	userAddr, err := resolver.GetRandomAddr(ctx, userService)
	if err != nil {
		return err
	}

	// iu is used for update publisher's upvote count
	var iu int32
	if isUpvoted {
		iu = 1
	} else {
		iu = -1
	}

	gid, saga := wrapSaga(ctx, v.data.dtmServerAddr,
		// 1. upvote/cancel upvote comment
		&sagaMember{
			method:           videoAddr + vcommentGRPC.VideoComment_UpvoteComment_FullMethodName,
			compensateMethod: videoAddr + vcommentGRPC.VideoComment_UpvoteCommentRevert_FullMethodName,
			payload: &vcommentGRPC.UpvoteCommentReq{
				CommentId: commentId,
				UserId:    userId,
				IsUpvoted: isUpvoted,
			},
		},

		// 2. update the publisher's upvote count
		&sagaMember{
			method:           userAddr + userinfoGRPC.Userinfo_UpdateUserCntLikes_FullMethodName,
			compensateMethod: userAddr + userinfoGRPC.Userinfo_UpdateUserCntLikesRevert_FullMethodName,
			payload: &userinfoGRPC.UpdateUserCntLikesReq{
				UserId:     publisherId,
				UpvoteFlag: -iu,
			},
		},
	)
	log.Infof("[UpvoteComment] gid: %s", gid)

	err = saga.Submit()
	if err != nil {
		return err
	}

	return nil
}

func (v *videoCommentRepo) CheckIfUserUpvotedComment(ctx context.Context, commentId int64, userId int64) (bool, error) {
	resp, err := v.data.videoCommentClient.CheckIfUserUpvotedComment(ctx, &vcommentGRPC.CheckIfUserUpvotedCommentReq{
		CommentId: commentId,
		UserId:    userId,
	})
	if err != nil {
		return false, err
	}
	return resp.IsUpvoted, nil
}

func (v *videoCommentRepo) GetPublisherId(ctx context.Context, commentId int64) (int64, error) {
	resp, err := v.data.videoCommentClient.GetPublisherId(ctx, &vcommentGRPC.GetPublisherIdReq{
		CommentId: commentId,
	})
	if err != nil {
		return 0, err
	}
	return resp.PublisherId, nil
}

func convert(comments []*vcommentGRPC.CommentInfo) []*vcommentv1.CommentInfo {
	resp := make([]*vcommentv1.CommentInfo, len(comments))
	for i, comment := range comments {
		resp[i] = &vcommentv1.CommentInfo{
			CommentId:   comment.CommentId,
			PublisherId: comment.PublisherId,
			RootId:      comment.RootId,
			Content:     comment.Content,
			CreatedAt:   comment.CreatedAt,
			UpvoteCount: comment.UpvoteCount,
			IpAddress:   comment.IpAddress,
			CntReplies:  comment.CntReplies,
		}
	}
	return resp
}
