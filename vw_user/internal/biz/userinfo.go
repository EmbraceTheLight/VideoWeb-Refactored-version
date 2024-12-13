package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"vw_user/internal/data/dal/model"
	"vw_user/internal/pkg/captcha"
	"vw_user/internal/pkg/middlewares/auth"
)

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

// Padding fills the UserSummaryInfo with user data
func (u *UserSummaryInfo) Padding(user *model.User) {
	u.UserID = user.UserID
	u.Username = user.Username
	u.Password = user.Password
	u.Email = user.Email
	u.Gender = user.Gender
	u.Signature = user.Signature
	u.AvatarPath = user.AvatarPath
	u.IsAdmin = user.IsAdmin
	u.Birthday = user.Birthday
}

type UserInfoRepo interface {
	// GetUserSummaryInfoByUserId get user summary info by user id
	GetUserSummaryInfoByUserId(ctx context.Context, userId int64) (*UserSummaryInfo, error)

	// GetUserSummaryInfoByUsername get user summary info by username
	GetUserSummaryInfoByUsername(ctx context.Context, username string) (*UserSummaryInfo, error)

	// GetUserInfo get user detail by user id
	GetUserInfo(ctx context.Context, userId int64) (*model.User, error)

	// GetUserLevelById get user level by user id
	GetUserLevelById(ctx context.Context, userId int64) (*model.UserLevel, error)

	// GetUserFansByUserID get user fans by user id
	GetUserFansByUserID(ctx context.Context, userId int64) ([]*UserSummaryInfo, error)

	// GetUserFollowersByUserIDFavoriteID GetUserFollowersByUserID get user followers by user id and followList id
	GetUserFollowersByUserIDFavoriteID(ctx context.Context, userId, followListID int64) ([]*UserSummaryInfo, error)
}

type UserInfoUsecase struct {
	infoRepo      UserInfoRepo
	email         *captcha.Email
	jwtAuthorizer *auth.JWTAuth
	logger        *log.Helper
}

func NewUserInfoUsecase(repo UserInfoRepo, jwtAuthorizer *auth.JWTAuth, email *captcha.Email, logger log.Logger) *UserInfoUsecase {
	return &UserInfoUsecase{
		infoRepo:      repo,
		jwtAuthorizer: jwtAuthorizer,
		email:         email,
		logger:        log.NewHelper(logger),
	}
}

func (u *UserInfoUsecase) GetUserDetail(ctx context.Context, userId int64) (*model.User, *model.UserLevel, error) {
	userinfo, err := u.infoRepo.GetUserInfo(ctx, userId)
	if err != nil {
		return nil, nil, err
	}
	userlevel, err := u.infoRepo.GetUserLevelById(ctx, userId)
	if err != nil {
		return nil, nil, err
	}
	return userinfo, userlevel, nil
}

//func (u *UserInfoUsecase)

// GetCaptchaCode get captcha code and send to `sendto`
func (u *UserInfoUsecase) GetCaptchaCode(ctx context.Context, sendto string) (string, error) {
	code := u.email.CreateVerificationCode()
	err := u.email.SendCode(ctx, sendto, code)
	if err != nil {
		return "", err
	}
	return code, nil
}
