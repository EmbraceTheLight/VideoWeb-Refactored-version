package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type FollowRepo interface {
	FollowOtherUser(ctx context.Context, followerId, followedId, followListId int64) error
	UnfollowOtherUser(ctx context.Context, followerId, followedId int64) error
}

type FollowUsecase struct {
	repo FollowRepo
	log  *log.Helper
}

func NewFollowUsecase(repo FollowRepo, logger log.Logger) *FollowUsecase {
	return &FollowUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (u *FollowUsecase) FollowOtherUser(ctx context.Context, followerId, followedId, followListId int64) error {
	return u.repo.FollowOtherUser(ctx, followerId, followedId, followListId)
}

func (u *FollowUsecase) UnfollowOtherUser(ctx context.Context, followerId, followedId int64) error {
	return u.repo.UnfollowOtherUser(ctx, followerId, followedId)
}
