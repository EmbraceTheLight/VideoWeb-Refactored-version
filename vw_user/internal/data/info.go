package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"vw_user/internal/biz"
	"vw_user/internal/data/dal/model"
	"vw_user/internal/data/dal/query"
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

func (u *userInfoRepo) GetUserSummaryInfoByUsername(ctx context.Context, username string) (*biz.UserSummaryInfo, error) {
	user := query.User
	userdo := user.WithContext(ctx)
	tmp, err := userdo.Where(user.Username.Eq(username)).
		Select(user.UserID, user.Username, user.IsAdmin, user.Email, user.Password,
			user.Gender, user.AvatarPath, user.Birthday, user.Signature).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdef.ErrUserNotFound
		}
		return nil, err
	}
	return modelUser2userSummary(tmp)[0], nil
}

func (u *userInfoRepo) GetUserSummaryInfoByUserId(ctx context.Context, userid int64) (*biz.UserSummaryInfo, error) {
	user := query.User
	userdo := user.WithContext(ctx)
	tmp, err := userdo.Where(user.UserID.Eq(userid)).
		Select(user.UserID, user.Username, user.IsAdmin, user.Email, user.Password,
			user.Gender, user.AvatarPath, user.Birthday, user.Signature).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdef.ErrUserNotFound
		}
		return nil, err
	}
	return modelUser2userSummary(tmp)[0], nil
}

func (u *userInfoRepo) GetUserInfo(ctx context.Context, userId int64) (*model.User, error) {
	user := query.User
	userdo := user.WithContext(ctx)
	userInfo, err := userdo.Omit(user.Version).Where(user.UserID.Eq(userId)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdef.ErrUserNotFound
		}
		return nil, err
	}
	return userInfo, nil
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

func (u *userInfoRepo) GetUserFansByUserID(ctx context.Context, userId int64) ([]*biz.UserSummaryInfo, error) {
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
	return getUsersSummaryInfoByIDs(ctx, fansIds)
}

func (u *userInfoRepo) GetUserFollowersByUserIDFavoriteID(ctx context.Context, userId, followListId int64) ([]*biz.UserSummaryInfo, error) {
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
	return getUsersSummaryInfoByIDs(ctx, followersIds)
}

func modelUser2userSummary(users ...*model.User) []*biz.UserSummaryInfo {
	if len(users) == 0 {
		return nil
	}
	size := len(users)
	result := make([]*biz.UserSummaryInfo, size)
	for i, user := range users {
		result[i] = new(biz.UserSummaryInfo)
		result[i].Padding(user)
	}
	return result
}

func getUsersSummaryInfoByIDs(ctx context.Context, userIds []int64) ([]*biz.UserSummaryInfo, error) {
	user := query.User
	userdo := user.WithContext(ctx)
	userInfo, err := userdo.Where(user.UserID.In(userIds...)).Find()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errdef.ErrUserNotFound
		}
		return nil, err
	}
	return modelUser2userSummary(userInfo...), nil
}
