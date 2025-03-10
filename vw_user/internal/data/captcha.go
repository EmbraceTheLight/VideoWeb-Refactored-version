package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"vw_user/internal/biz"
)

type captchaRepo struct {
	logger *log.Helper
	data   *Data
}

func NewCaptRepo(data *Data, logger log.Logger) biz.CaptchaRepo {
	return &captchaRepo{
		logger: log.NewHelper(logger),
		data:   data,
	}
}

func (r *captchaRepo) DeleteCodeFromCache(ctx context.Context, email string) error {
	_, err := r.data.redis.Del(ctx, email).Result()
	return err
}

func (r *captchaRepo) SetCodeToCache(ctx context.Context, email, code string, expiration time.Duration) error {
	_, err := r.data.redis.SetEx(ctx, email, code, expiration).Result()
	return err

}

func (r *captchaRepo) GetCodeFromCache(ctx context.Context, email string) (string, error) {
	code, err := r.data.redis.Get(ctx, email).Result()
	return code, err
}
