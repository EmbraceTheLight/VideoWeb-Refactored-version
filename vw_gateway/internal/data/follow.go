package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vw_gateway/internal/biz"
	followv1 "vw_user/api/v1/follow"
)

type followRepo struct {
	data *Data
	log  *log.Helper
}

func NewFollowRepo(data *Data, logger log.Logger) biz.FollowRepo {
	return &followRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *followRepo) FollowOtherUser(ctx context.Context, followerId, followedId, followListId int64) error {
	_, err := r.data.followClient.FollowUser(ctx, &followv1.FollowUserReq{
		UserId:         followerId,
		FollowerUserId: followedId,
		FollowListId:   followListId,
	})
	return err
}

func (r *followRepo) UnfollowOtherUser(ctx context.Context, followerId, followedId int64) error {
	_, err := r.data.followClient.UnfollowUser(ctx, &followv1.UnfollowUserReq{
		UserId:         followerId,
		FolloweeUserId: followedId,
	})
	return err
}
