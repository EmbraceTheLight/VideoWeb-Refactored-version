package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type FavoritesRepo interface {
	CreateFavorites(ctx context.Context, userId int64, favoritesName string, isPublic int32, description string) error
	DeleteFavorites(ctx context.Context, favoritesId int64) error
	ModifyFavorites(ctx context.Context, userId int64, favoritesName *string, isPrivate *int32, description *string) error
}

type FavoritesUsecase struct {
	repo   FavoritesRepo
	logger *log.Helper
}

func NewFavoritesUsecase(repo FavoritesRepo, logger log.Logger) *FavoritesUsecase {
	return &FavoritesUsecase{
		repo:   repo,
		logger: log.NewHelper(logger),
	}
}

func (uc *FavoritesUsecase) CreateFavorites(ctx context.Context, userId int64, favoritesName string, isPublic int32, description string) error {
	return uc.repo.CreateFavorites(ctx, userId, favoritesName, isPublic, description)
}

func (uc *FavoritesUsecase) DeleteFavorites(ctx context.Context, favoritesId int64) error {
	return uc.repo.DeleteFavorites(ctx, favoritesId)
}

func (uc *FavoritesUsecase) ModifyFavorites(ctx context.Context, userId int64, favoritesName *string, isPrivate *int32, description *string) error {
	return uc.repo.ModifyFavorites(ctx, userId, favoritesName, isPrivate, description)
}
