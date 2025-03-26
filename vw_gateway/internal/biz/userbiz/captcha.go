package userbiz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vw_gateway/internal/pkg/captcha"
)

type CaptchaRepo interface {
	// DeleteCodeFromCache delete code from Redis cache
	DeleteCodeFromCache(ctx context.Context, email string) error

	// SetCodeToCache set code to Redis cache
	SetCodeToCache(ctx context.Context, email, code string) error
}

type CaptchaUsecase struct {
	logger      *log.Helper
	captchaRepo CaptchaRepo
	email       *captcha.Email
}

func NewCaptchaUsecase(
	logger log.Logger,
	repo CaptchaRepo,
	email *captcha.Email,
) *CaptchaUsecase {
	return &CaptchaUsecase{
		logger:      log.NewHelper(logger),
		captchaRepo: repo,
		email:       email,
	}
}

func (cu *CaptchaUsecase) SendCodeCaptcha(ctx context.Context, email string) (string, error) {
	err := cu.captchaRepo.DeleteCodeFromCache(ctx, email)
	if err != nil {
		return "", err
	}
	code := cu.email.CreateVerificationCode()
	err = cu.email.SendCode(ctx, email, code)
	if err != nil {
		return "", err
	}
	err = cu.captchaRepo.SetCodeToCache(ctx, email, code)
	if err != nil {
		return "", err
	}
	return code, nil
}
