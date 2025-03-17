package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vw_user/internal/biz"
	"vw_user/internal/data/dal/model"
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

// CreateFollowRecord inserts a new follow record into the db.
func (repo *followRepo) CreateFollowRecord(ctx context.Context, follow *model.UserFollow) error {
	userFollow := getQuery(ctx).UserFollow
	userFollowDo := userFollow.WithContext(ctx)
	return userFollowDo.Create(follow)
}

// DeleteFollowRecord deletes a follow record from the db.
func (repo *followRepo) DeleteFollowRecord(ctx context.Context, follow *model.UserFollow) error {
	userFollow := getQuery(ctx).UserFollow
	userFollowDo := userFollow.WithContext(ctx)
	_, err := userFollowDo.Where(
		userFollow.FollowUserID.Eq(follow.FollowUserID),
		userFollow.UserID.Eq(follow.UserID),
	).Delete(follow)
	return err
}
