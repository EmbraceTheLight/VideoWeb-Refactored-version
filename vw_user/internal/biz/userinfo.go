package biz

import (
	"context"
	stderr "errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"unicode/utf8"
	"util/helper"
	"vw_user/internal/data/dal/model"
	"vw_user/internal/domain"
	"vw_user/internal/pkg/ecode/errdef"
)

type UserInfoRepo interface {
	// GetUserInfoByUserId get user summary info by user id
	GetUserInfoByUserId(ctx context.Context, userId int64) (*domain.UserInfo, error)

	// GetUserInfoByUsername get user summary info by username
	GetUserInfoByUsername(ctx context.Context, username string) (*domain.UserInfo, error)

	// GetUserLevelById get user level by user id
	GetUserLevelById(ctx context.Context, userId int64) (*model.UserLevel, error)

	// GetUserFansByUserID get user fans by user id
	GetUserFansByUserID(ctx context.Context, userId int64) ([]*domain.UserInfo, error)

	// GetUserFollowersByUserIDFavoriteID GetUserFollowersByUserID get user followers by user id and followList id
	GetUserFollowersByUserIDFollowListID(ctx context.Context, userId, followListID int64) ([]*domain.UserSummary, error)

	// UpdateEmail update user email
	UpdateEmail(ctx context.Context, userId int64, newEmail string) error

	// UpdatePassword update user password. Front-end should check if the password is valid
	UpdatePassword(ctx context.Context, userId int64, newPassword string) error

	// UpdateUserSignature update user signature
	UpdateUserSignature(ctx context.Context, userId int64, signature string) error

	// CheckUsernameConflict check if the new username is conflict with existing username
	CheckUsernameConflict(ctx context.Context, newUsername string) (ok bool, err error)

	// UpdateUsername update user username
	UpdateUsername(ctx context.Context, userId int64, newUsername string) error

	// UpdateCntFollows update user fans count
	UpdateCntFollows(ctx context.Context, userId int64, change int64) error

	// UpdateCntFans update user fans count
	UpdateCntFans(ctx context.Context, userId int64, change int64) error
}

type UserInfoUsecase struct {
	infoRepo UserInfoRepo
	captcha  CaptchaRepo
	logger   *log.Helper
}

func NewUserInfoUsecase(repo UserInfoRepo, captcha CaptchaRepo, logger log.Logger) *UserInfoUsecase {
	return &UserInfoUsecase{
		infoRepo: repo,
		captcha:  captcha,
		logger:   log.NewHelper(logger),
	}
}

func (uic *UserInfoUsecase) GetUserinfo(ctx context.Context, userId int64) (*domain.UserInfo, error) {
	user, err := uic.infoRepo.GetUserInfoByUserId(ctx, userId)
	if err != nil {
		return nil, helper.HandleError(errdef.ErrGetUserInfoFailed, err)
	}

	// TODO：向 vw_video 系统请求视频数量，并填充到 UserInfo 结构体中
	return user, nil

}

func (uic *UserInfoUsecase) ModifyEmail(ctx context.Context, userID int64, email string, inputCode string) (newEmail string, err error) {
	// * DELETE the verify code from Redis cache after updating the email,
	// * NO MATTER this update is successful or not.
	defer uic.captcha.DeleteCodeFromCache(ctx, email)

	// Get the verify code from Redis cache and compare it with the input code
	verifyCode, err := uic.captcha.GetCodeFromCache(ctx, email)
	if err != nil {
		return "", helper.HandleError(errdef.ErrModifyEmailFailed, err)
	}
	if verifyCode != inputCode {
		return "", helper.HandleError(errdef.ErrModifyEmailFailed, stderr.New("用户输入的验证码错误"))
	}

	// Update the email in the database
	err = uic.infoRepo.UpdateEmail(ctx, userID, email)
	if err != nil {
		return "", helper.HandleError(errdef.ErrModifyEmailFailed, err)
	}

	return email, nil
}

func (uic *UserInfoUsecase) ModifyPassword(ctx context.Context, userId int64, oldPassword, newPassword string) error {
	info, err := uic.infoRepo.GetUserInfoByUserId(ctx, userId)
	if err != nil {
		return helper.HandleError(errdef.ErrModifyPasswordFailed, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(oldPassword))
	if err != nil {
		return helper.HandleError(errdef.ErrUserPasswordError, nil)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	err = uic.infoRepo.UpdatePassword(ctx, userId, string(hashedPassword))
	if err != nil {
		return helper.HandleError(errdef.ErrModifyPasswordFailed, err)
	}
	return nil
}

func (uic *UserInfoUsecase) ModifyUserSignature(ctx context.Context, userId int64, signature string) (newSignature string, err error) {
	// Check if the signature length is valid
	if utf8.RuneCountInString(signature) > 25 {
		return "", helper.HandleError(errdef.ErrModifySignatureFailed, stderr.New("签名不能超过25个字符"))
	}

	err = uic.infoRepo.UpdateUserSignature(ctx, userId, signature)
	if err != nil {
		return "", helper.HandleError(errdef.ErrModifySignatureFailed, err)
	}
	return signature, nil
}

func (uic *UserInfoUsecase) ForgetPassword(ctx context.Context, userId int64, email string, inputCode string, newPassword string) error {
	// * DELETE the verify code from Redis cache after this function call.
	defer uic.captcha.DeleteCodeFromCache(ctx, email)

	// Get the verify code from Redis cache and compare it with the input code
	verify, err := uic.captcha.GetCodeFromCache(ctx, email)
	if err != nil {
		return helper.HandleError(errdef.ErrForgetPasswordFailed, err)
	}
	if verify != inputCode {
		return helper.HandleError(errdef.ErrForgetPasswordFailed, errdef.ErrVerifyCodeNotMatch)
	}

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	// Update the password in the database
	err = uic.infoRepo.UpdatePassword(ctx, userId, string(encryptedPassword))
	if err != nil {
		return helper.HandleError(errdef.ErrForgetPasswordFailed, err)
	}

	return nil
}

func (uic *UserInfoUsecase) ModifyUsername(ctx context.Context, userId int64, username string) (newUsername string, err error) {
	// Check if the new username is conflict with existing username
	ok, err := uic.infoRepo.CheckUsernameConflict(ctx, username)
	if !ok {
		return "", helper.HandleError(errdef.ErrModifyUsernameFailed, stderr.New("该用户名已存在"))
	}
	if err != nil {
		return "", helper.HandleError(errdef.ErrModifyUsernameFailed, err)
	}

	// The new username is valid, update the username in the database
	err = uic.infoRepo.UpdateUsername(ctx, userId, username)
	if err != nil {
		return "", helper.HandleError(errdef.ErrModifyUsernameFailed, err)
	}

	return username, nil
}
