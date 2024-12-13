package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
	"vw_user/internal/biz"
	"vw_user/internal/data/dal/query"
	"vw_user/internal/pkg/captcha"
)

const (
	login      = "login"
	expOfLogin = 5
)

type userIdentityRepo struct {
	data   *Data
	logger *log.Helper
}

func (u *userIdentityRepo) DeleteCodeFromCache(ctx context.Context, email string) {
	u.data.redis.Del(ctx, email)
}

func (u *userIdentityRepo) GetCodeFromCache(ctx context.Context, email string) (code string, err error) {
	return u.data.redis.Get(ctx, email).Result()
}

func (u *userIdentityRepo) SetCodeCache(ctx context.Context, email, code string) error {
	_, err := u.data.redis.Set(ctx, email, code, captcha.ExpirationTime).Result()
	if err != nil {
		return err
	}
	return nil
}

func (u *userIdentityRepo) AddExpForLogin(ctx context.Context, userId int64) error {
	userLevel := query.UserLevel
	userLevelDo := userLevel.WithContext(ctx)
	level, err := userLevelDo.Where(userLevel.UserID.Eq(userId)).First()
	if err != nil {
		return err
	}
	loginKey := strconv.FormatInt(userId, 10) + login
	_, err = u.data.redis.Get(ctx, loginKey).Result()
	// find user info of logging in redis user has logged in today,don't add the exp
	if err == nil {
		return nil
	} else {
		// no record in redis,add the exp and set the logging record in redis
		if errors.Is(err, redis.Nil) {
			//add the exp
			level.AddExp(expOfLogin)
			_, err = userLevelDo.Where(userLevel.UserID.Eq(userId)).Updates(level)
			if err != nil {
				return err
			}
			//get tomorrow time at 00:00:00 for setting the expiration of login record
			now := time.Now()
			t := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
			u.data.redis.Set(ctx, loginKey, "1", t.Sub(now))
		} else { // other redis error
			return err
		}
	}
	return nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserIdentityRepo {
	return &userIdentityRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}
}
