package userdata

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vw_gateway/internal/biz/userbiz"
	"vw_gateway/internal/domain"
	followv1 "vw_user/api/v1/follow"
)

type followRepo struct {
	data *Data
	log  *log.Helper
}

func NewFollowRepo(data *Data, logger log.Logger) userbiz.FollowRepo {
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

func (r *followRepo) GetFolloweeInfo(ctx context.Context, userId int64, followListId int64, pageNum int32, pageSize int32) ([]*domain.UserSummary, error) {
	followeesList, err := r.data.followClient.GetFolloweesInfo(ctx, &followv1.GetFolloweesInfoReq{
		UserId:       userId,
		FollowListId: followListId,
		PageNum:      pageNum,
		PageSize:     pageSize,
	})
	if err != nil {
		return nil, err
	}

	infoList := make([]*domain.UserSummary, len(followeesList.FolloweesInfo))
	for i, followee := range followeesList.FolloweesInfo {
		infoList[i] = &domain.UserSummary{
			Username:   followee.Username,
			Email:      followee.Email,
			Signature:  followee.Signature,
			AvatarPath: followee.AvatarPath,
			Gender:     followee.Gender,
			Birthday:   followee.Birthday,
		}
	}
	return infoList, err
}
