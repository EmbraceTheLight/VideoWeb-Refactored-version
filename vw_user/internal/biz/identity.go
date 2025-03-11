package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"unicode/utf8"
	"util/helper"
	"util/snowflake"
	idv1 "vw_user/api/v1/identity"
	"vw_user/internal/data/dal/model"
	"vw_user/internal/data/dal/query"
	"vw_user/internal/pkg/ecode/errdef"
)

type UserIdentityRepo interface {
	CacheAccessToken(ctx context.Context, userId, accessToken string, expiration time.Duration) error
	AddExpForLogin(ctx context.Context, userID int64) error
	CheckPassword(ctx context.Context, password, hashedPassword string) error
	CreatRecordsForRegister(ctx context.Context, newUser *model.User) error
	DeleteCachedAccessToken(ctx context.Context, userId string) error
}

type RegisterInfo struct {
	UserID         *int64
	Username       string
	Password       string
	RepeatPassword string
	Gender         int32
	InputCode      string
	Email          string
	Signature      string
	Birthday       time.Time
	AvatarFilePath *string
}

type UserIdentityUsecase struct {
	infoRepo     UserInfoRepo
	identityRepo UserIdentityRepo
	captcha      CaptchaRepo
	logger       *log.Helper
}

func NewUserIdentityUsecase(idRepo UserIdentityRepo, infoRepo UserInfoRepo, captcha CaptchaRepo, logger log.Logger) *UserIdentityUsecase {
	return &UserIdentityUsecase{
		identityRepo: idRepo,
		infoRepo:     infoRepo,
		captcha:      captcha,
		logger:       log.NewHelper(logger),
	}
}

func (uc *UserIdentityUsecase) Register(ctx context.Context, registerInfo *RegisterInfo) (userID int64, isAdmin bool, err error) {
	newUser, err := uc.handleRegisterInfo(ctx, registerInfo)
	if err != nil {
		return 0, false, err
	}

	// * set default avatar for user, if he/she don't upload avatar.
	var avatarFilePath string
	if registerInfo.AvatarFilePath == nil {
		userDir := filepath.Join(resourcePath, strconv.FormatInt(*registerInfo.UserID, 10))
		err = helper.CreateDir(userDir, os.ModePerm)
		if err != nil {
			return 0, false, helper.HandleError(errdef.ErrCreateUserDirFailed, err)
		}

		defer func() {
			/* this defer is used to remove the user dir
			if there is any error Behind this logic*/
			if err != nil {
				_ = os.RemoveAll(userDir)
			}
		}()

		avatarFilePath = filepath.Join(userDir, defaultAvatarName)
		err = helper.WriteToNewFile(avatarFilePath, defaultAvatarPath)
		if err != nil {
			return 0, false, helper.HandleError(errdef.ErrCreateUserDirFailed, err)
		}
		newUser.AvatarPath = avatarFilePath
	}

	// create new user's records
	err = uc.identityRepo.CreatRecordsForRegister(ctx, newUser)
	if err != nil {
		return 0, false, helper.HandleError(errdef.ErrCreateUserRecordsFailed, err)
	}
	return newUser.UserID, newUser.IsAdmin, nil
}

func (uc *UserIdentityUsecase) CacheAccessToken(ctx context.Context, userId, accessToken string, expiration time.Duration) error {
	return uc.identityRepo.CacheAccessToken(ctx, userId, accessToken, expiration)
}

func (uc *UserIdentityUsecase) AddExpForLogin(ctx context.Context, userID int64) error {
	err := uc.identityRepo.AddExpForLogin(ctx, userID)
	if err != nil {
		// TODO: 若用户经验更新失败，使用异步方法将消息传送到消息队列中，后台异步更新用户经验，现阶段不做处理
	}

	return nil
}

func (uc *UserIdentityUsecase) Logout(ctx context.Context, userId string) error {
	err := uc.identityRepo.DeleteCachedAccessToken(ctx, userId)
	if err != nil {
		return errdef.ErrLogoutFailed
	}
	return nil
}

// handleRegisterInfo check the register information
// create a new user model and fill its dynamic fields (such as encrypted password, e.g.)
// and return the user model, if there is no error.
func (uc *UserIdentityUsecase) handleRegisterInfo(ctx context.Context, registerInfo *RegisterInfo) (*model.User, error) {
	// check register info
	userdo := query.User
	count, err := userdo.Where(userdo.Username.Eq(registerInfo.Username)).Count()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.HandleError(errdef.INTERNAL_ERROR, err)
		}
	}

	verifyCode, err := uc.captcha.GetCodeFromCache(ctx, registerInfo.Email)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, helper.HandleError(errdef.ErrVerifyCodeExpired, nil)
		} else {
			return nil, helper.HandleError(errdef.INTERNAL_ERROR, err)
		}
	}

	switch {
	case count > 0: //已有同名用户
		return nil, errdef.ErrUserAlreadyExist
	case len(registerInfo.Password) < 6: //密码长度小于6位
		return nil, errdef.ErrPasswordTooShort
	case registerInfo.Password != registerInfo.RepeatPassword: //第一次输入的密码与第二次输入的密码不一致.
		return nil, errdef.ErrPasswordNotMatch
	case utf8.RuneCountInString(registerInfo.Signature) > 25:
		return nil, errdef.ErrSignatureTooLong
	case registerInfo.InputCode != verifyCode:
		return nil, errdef.ErrVerifyCodeNotMatch
	}

	// encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errdef.ErrEncryptPasswordFailed
	}

	// create new user model
	// -> this `if` block checks if the user upload avatar file while registering.
	// -> if the user upload avatar file, the UserID will be set by the client.
	if *registerInfo.UserID <= 0 {
		*registerInfo.UserID = snowflake.GetID()
	}
	newUser := &model.User{
		UserID:    *registerInfo.UserID,
		Username:  registerInfo.Username,
		Password:  string(hashedPassword),
		Gender:    int(registerInfo.Gender),
		Email:     registerInfo.Email,
		Signature: registerInfo.Signature,
		Birthday:  registerInfo.Birthday,
	}

	return newUser, nil
}

func (uc *UserIdentityUsecase) Login(ctx context.Context, username, password string) (*idv1.LoginResp, error) {
	userinfo, err := uc.infoRepo.GetUserInfoByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	err = uc.identityRepo.CheckPassword(ctx, password, userinfo.Password)
	if err != nil {
		return nil, err
	}

	err = uc.identityRepo.AddExpForLogin(ctx, userinfo.UserId)
	if err != nil {
		return nil, err
	}
	return &idv1.LoginResp{
		Code:     200,
		Msg:      "用户登录成功",
		UserId:   userinfo.UserId,
		Username: userinfo.Username,
		IsAdmin:  userinfo.IsAdmin,
	}, nil
}
