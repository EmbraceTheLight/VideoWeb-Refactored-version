package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"vw_gateway/internal/pkg/middlewares/auth"
)

type RegisterInfo struct {
	Username       string
	Password       string
	RepeatPassword string
	Gender         int32
	InputCode      string
	Email          string
	Signature      string
	Birthday       time.Time
}
type UserSummaryInfo struct {
	UserID     int64
	Username   string
	Password   string
	Email      string
	Gender     int
	Signature  string
	AvatarPath string
	Birthday   time.Time
	IsAdmin    bool
}

type LoginResponse struct {
	UserId   int64
	Username string
	IsAdmin  bool
}

type UserIdentityRepo interface {
	CacheAccessToken(ctx context.Context, accessToken string, expireTime time.Duration) error
	Register(ctx context.Context, registerInfo *RegisterInfo) (userID int64, isAdmin bool, err error)
	Login(ctx context.Context, username, password string) (*LoginResponse, error)
	Logout(ctx context.Context, accessToken string) error
}

type UserIdentityUsecase struct {
	identityRepo UserIdentityRepo
	jwt          *auth.JWTAuth
	logger       *log.Helper
}

func NewUserIdentityUsecase(repo UserIdentityRepo, jwt *auth.JWTAuth, logger log.Logger) *UserIdentityUsecase {
	return &UserIdentityUsecase{
		identityRepo: repo,
		jwt:          jwt,
		logger:       log.NewHelper(logger),
	}
}

func (uc *UserIdentityUsecase) Register(ctx context.Context, registerInfo *RegisterInfo) (accessToken, RefreshToken string, err error) {
	userID, isAdmin, err := uc.identityRepo.Register(ctx, registerInfo)
	if err != nil {
		return "", "", err
	}

	accessToken, RefreshToken, err = uc.jwt.CreateToken(userID, registerInfo.Username, isAdmin)
	return accessToken, RefreshToken, err
}

func (uc *UserIdentityUsecase) Login(ctx context.Context, username, password string) (accessToken, refreshToken string, err error) {
	resp, err := uc.identityRepo.Login(ctx, username, password)
	if err != nil {
		return "", "", err
	}

	atoken, rtoken, err := uc.jwt.CreateToken(resp.UserId, resp.Username, resp.IsAdmin)
	if err != nil {
		return "", "", err
	}

	err = uc.identityRepo.CacheAccessToken(ctx, atoken, uc.jwt.AccessExpireTime)
	if err != nil {
		// TODO: 做异步处理
		return "", "", err
	}

	return atoken, rtoken, nil
}

func (uc *UserIdentityUsecase) Logout(ctx context.Context, token string) error {
	return uc.identityRepo.Logout(ctx, token)
}
