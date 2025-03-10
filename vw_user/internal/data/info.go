package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"vw_user/internal/biz"
	"vw_user/internal/data/dal/model"
	"vw_user/internal/data/dal/query"
	"vw_user/internal/domain"
	"vw_user/internal/pkg/ecode/errdef"
)

type userInfoRepo struct {
	data   *Data
	logger *log.Helper
}

func NewUserInfoRepo(data *Data, logger log.Logger) biz.UserInfoRepo {
	return &userInfoRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

func (u *userInfoRepo) GetUserInfoByUsername(ctx context.Context, username string) (*domain.UserInfo, error) {
	user := query.User
	userdo := user.WithContext(ctx)
	userInfo, err := userdo.Where(user.Username.Eq(username)).
		Select(user.UserID, user.Username, user.IsAdmin, user.Email, user.Password,
			user.Gender, user.AvatarPath, user.Birthday, user.Signature, user.Shells,
			user.CntFollows, user.CntFans).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdef.ErrUserNotFound
		}
		return nil, err
	}

	return domain.NewUserInfo(userInfo), nil
}

func (u *userInfoRepo) GetUserInfoByUserId(ctx context.Context, userid int64) (*domain.UserInfo, error) {
	user := query.User
	userdo := user.WithContext(ctx)
	userInfo, err := userdo.Where(user.UserID.Eq(userid)).
		Select(user.UserID, user.Username, user.IsAdmin, user.Email, user.Password,
			user.Gender, user.AvatarPath, user.Birthday, user.Signature).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdef.ErrUserNotFound
		}
		return nil, err
	}
	return domain.NewUserInfo(userInfo), nil
}

func (u *userInfoRepo) GetUserLevelById(ctx context.Context, userId int64) (*model.UserLevel, error) {
	userLevel := query.UserLevel
	userLevelDo := userLevel.WithContext(ctx)
	userLevelInfo, err := userLevelDo.Where(userLevel.UserID.Eq(userId)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdef.ErrUserNotFound
		}
		return nil, err
	}
	return userLevelInfo, nil
}

func (u *userInfoRepo) GetUserFansByUserID(ctx context.Context, userId int64) ([]*domain.UserInfo, error) {
	// get user's fans' Ids
	var fansIds = make([]int64, 0)
	result := u.data.mysql.Select("followed_id").Where("user_id = ?", userId).Find(&fansIds)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errdef.ErrUserNotFound
		}
		return nil, result.Error
	}

	// get user's fans' info by fansIds
	users, err := getUsersSummaryInfoByIDs(ctx, fansIds)
	if err != nil {
		return nil, err
	}
	return domain.NewUserInfos(users...), nil
}

func (u *userInfoRepo) GetUserFollowersByUserIDFavoriteID(ctx context.Context, userId, followListId int64) ([]*domain.UserSummary, error) {
	//find followers' ids by user_id and followlist_id
	followersIds := make([]int64, 0)
	result := u.data.mysql.Select("user_id").
		Where("user_id = ?,followlist_id = ?", userId, followListId).Find(&followersIds)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errdef.ErrUserNotFound
		}
		return nil, result.Error
	}

	//find followers' info by followers' ids
	users, err := getUsersSummaryInfoByIDs(ctx, followersIds)
	if err != nil {
		return nil, err
	}
	return domain.NewUserSummaries(users...), nil
}

func (u *userInfoRepo) UpdateEmail(ctx context.Context, userId int64, newEmail string) error {
	user := query.User
	userDo := user.WithContext(ctx)
	_, err := userDo.Where(user.UserID.Eq(userId)).Update(user.Email, newEmail)
	if err != nil {
		return err
	}
	return nil
}

func (u *userInfoRepo) UpdatePassword(ctx context.Context, userId int64, encryptedPassword string) error {
	user := query.User
	userDo := user.WithContext(ctx)
	_, err := userDo.Where(user.UserID.Eq(userId)).Update(user.Password, encryptedPassword)
	if err != nil {
		return err
	}
	return nil
}

func (u *userInfoRepo) UpdateUserSignature(ctx context.Context, userId int64, signature string) error {
	user := query.User
	userDo := user.WithContext(ctx)
	_, err := userDo.Where(user.UserID.Eq(userId)).Update(user.Signature, signature)
	if err != nil {
		return err
	}
	return nil
}

// CheckUsernameConflict check if the new username is conflict with existing username
func (u *userInfoRepo) CheckUsernameConflict(ctx context.Context, newUsername string) (ok bool, err error) {
	user := query.User
	userDo := user.WithContext(ctx)
	count, err := userDo.Where(user.Username.Eq(newUsername)).Count()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
	}
	if count > 0 {
		return false, nil
	}
	return true, nil
}

// UpdateUsername update user username
func (u *userInfoRepo) UpdateUsername(ctx context.Context, userId int64, newUsername string) error {
	user := query.User
	userDo := user.WithContext(ctx)
	_, err := userDo.Where(user.UserID.Eq(userId)).Update(user.Username, newUsername)
	if err != nil {
		return err
	}
	return nil
}

func getUsersSummaryInfoByIDs(ctx context.Context, userIds []int64) ([]*model.User, error) {
	user := query.User
	userdo := user.WithContext(ctx)
	userInfo, err := userdo.Where(user.UserID.In(userIds...)).Find()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdef.ErrUserNotFound
		}
		return nil, err
	}
	return userInfo, nil
}
