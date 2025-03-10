package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type CaptchaRepo interface {
	// DeleteCodeFromCache delete code cache from Redis
	DeleteCodeFromCache(ctx context.Context, email string) error
	SetCodeToCache(ctx context.Context, email, code string, expiration time.Duration) error
	GetCodeFromCache(ctx context.Context, email string) (string, error)
}

type CaptchaUsecase struct {
	logger *log.Helper
	repo   CaptchaRepo
}

func NewCaptchaUsecase(logger log.Logger, repo CaptchaRepo) *CaptchaUsecase {
	return &CaptchaUsecase{
		logger: log.NewHelper(logger),
		repo:   repo,
	}
}

func (cu *CaptchaUsecase) DeleteCodeFromCache(ctx context.Context, email string) error {
	return cu.repo.DeleteCodeFromCache(ctx, email)
}

func (cu *CaptchaUsecase) SetCodeToCache(ctx context.Context, email, code string, expiration time.Duration) error {
	return cu.repo.SetCodeToCache(ctx, email, code, expiration)
}
