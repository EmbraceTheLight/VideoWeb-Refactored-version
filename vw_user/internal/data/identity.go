package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
	"util/helper"
	"util/snowflake"
	"vw_user/internal/biz"
	"vw_user/internal/data/dal/model"
	"vw_user/internal/data/dal/query"
	"vw_user/internal/pkg/ecode/errdef"
)

const (
	login      = "login"
	expOfLogin = 5
)

type userIdentityRepo struct {
	data   *Data
	logger *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserIdentityRepo {
	return &userIdentityRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

func (u *userIdentityRepo) AddExpForLogin(ctx context.Context, userId int64) error {
	userLevel := query.UserLevel
	userLevelDo := userLevel.WithContext(ctx)
	level, err := userLevelDo.Where(userLevel.UserID.Eq(userId)).First()
	if err != nil {
		return err
	}
	loginKey := strconv.FormatInt(userId, 10) + login

	// check if user has logged in today
	_, err = u.data.redis.Get(ctx, loginKey).Result()

	if err != nil {
		// no record in redis,add the exp and set the logging record in redis
		if errors.Is(err, redis.Nil) {
			//add the exp and update it in db
			level.AddExp(expOfLogin)
			_, err = userLevelDo.Where(userLevel.UserID.Eq(userId)).Updates(level)
			if err != nil {
				return err
			}

			//get tomorrow time at 00:00:00 for setting the expiration of login record
			now := time.Now()
			t := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
			u.data.redis.SetEx(ctx, loginKey, "1", t.Sub(now))
		} else { // other redis error
			return err
		}
	}
	return nil
}

func (u *userIdentityRepo) CacheAccessToken(ctx context.Context, accessToken string, expiration time.Duration) error {
	_, err := u.data.redis.SetEx(ctx, accessToken, "1", expiration).Result()
	if err != nil {
		return helper.HandleError(errdef.ErrCacheAccessToken, err)
	}
	return nil
}

func (u *userIdentityRepo) CreatRecordsForRegister(ctx context.Context, newUser *model.User) error {
	var err error
	defaultFavorites := &model.UserFavorite{
		FavoriteID:   snowflake.GetID(),
		UserID:       newUser.UserID,
		FavoriteName: "默认收藏夹",
		Description:  "",
		IsPrivate:    1,
	}
	privateFavorites := &model.UserFavorite{
		FavoriteID:   snowflake.GetID(),
		UserID:       newUser.UserID,
		FavoriteName: "私密收藏夹",
		Description:  "",
		IsPrivate:    -1,
	}
	userLevel := &model.UserLevel{
		UserID: newUser.UserID,
	}
	defaultFollowList := &model.FollowList{
		ListID:   snowflake.GetID(),
		UserID:   newUser.UserID,
		ListName: "默认关注列表",
	}

	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.User.Create(newUser)
	if err != nil {
		return err
	}

	err = tx.FollowList.Create(defaultFollowList)
	if err != nil {
		return err
	}

	err = tx.UserFavorite.Create(defaultFavorites, privateFavorites)
	if err != nil {
		return err
	}

	err = tx.UserLevel.Create(userLevel)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (u *userIdentityRepo) DeleteCachedAccessToken(ctx context.Context, accessToken string) error {
	_, err := u.data.redis.Del(ctx, accessToken).Result()
	return err
}
