package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"util/dbutil/mysqlutil"
	"util/getid"
	"util/helper"
	"vw_user/internal/data/dal/model"
	"vw_user/internal/domain"
	"vw_user/internal/pkg/ecode/errdef"
)

type FavoritesRepo interface {
	// CheckFavoritesName Check if the userbiz has already created a favorites with the same name
	// Return FALSE if the userbiz has already created a favorites with the same name, otherwise return TRUE
	CheckFavoritesName(ctx context.Context, userId int64, favoritesName string) (bool, error)
	CreateFavorites(ctx context.Context, favorite *model.Favorite) error

	// DeleteFavorites Delete the specified favorites
	DeleteFavorites(ctx context.Context, favoritesId int64) error

	// CheckIfFavoritesEmpty check if the specified favorites is empty.
	// If the favorites is empty(Can be deleted), return TRUE, otherwise return FALSE.
	CheckIfFavoritesEmpty(ctx context.Context, favoritesId int64) (bool, error)

	// UpdateFavorites Update the specified favorites
	UpdateFavorites(ctx context.Context, favorite map[string]any) error
}

type FavoritesUsecase struct {
	repo   FavoritesRepo
	logger *log.Helper
	tx     mysqlutil.Transaction
}

func NewFavoritesUsecase(repo FavoritesRepo, tx mysqlutil.Transaction, logger log.Logger) *FavoritesUsecase {
	return &FavoritesUsecase{
		repo:   repo,
		tx:     tx,
		logger: log.NewHelper(logger),
	}
}

func (uc *FavoritesUsecase) CreateFavorites(ctx context.Context, info *domain.FavoritesInfo) error {
	ok, err := uc.repo.CheckFavoritesName(ctx, info.UserId, *info.FavoritesName)
	if err != nil {
		return helper.HandleError(errdef.ErrCreateFavoriteFailed, err)
	}
	if !ok {
		return helper.HandleError(errdef.ErrCreateFavoriteFailed, errdef.ErrFavoritesNameConflict)
	}

	err = uc.repo.CreateFavorites(ctx, &model.Favorite{
		FavoritesID:   getid.GetID(),
		UserID:        info.UserId,
		FavoritesName: *info.FavoritesName,
		Description:   *info.Description,
		IsPrivate:     int64(*info.IsPrivate),
	})
	if err != nil {
		return helper.HandleError(errdef.ErrCreateFavoriteFailed, err)
	}
	return nil
}

func (uc *FavoritesUsecase) DeleteFavorites(ctx context.Context, favoritesId int64) error {
	err := uc.tx.WithTx(ctx, func(ctx context.Context) error {
		// 1. Delete the records of the specified favorites
		ok, err := uc.repo.CheckIfFavoritesEmpty(ctx, favoritesId)
		if err != nil {
			return helper.HandleError(errdef.ErrDeleteFavoriteFailed, err)
		}
		if !ok {
			return helper.HandleError(errdef.ErrDeleteFavoriteFailed, errdef.ErrFavoritesNotEmpty)
		}

		// 2. Delete the specified favorites itself
		err = uc.repo.DeleteFavorites(ctx, favoritesId)
		if err != nil {
			return helper.HandleError(errdef.ErrDeleteFavoriteFailed, err)
		}
		return nil
	})
	return err
}

func (uc *FavoritesUsecase) ModifyFavorites(ctx context.Context, info *domain.FavoritesInfo) error {
	err := uc.repo.UpdateFavorites(ctx, info.ToMap())
	if err != nil {
		return helper.HandleError(errdef.ErrModifyFavoriteFailed, err)
	}
	return nil
}
