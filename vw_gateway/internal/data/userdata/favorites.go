package userdata

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vw_gateway/internal/biz/userbiz"
	favorv1 "vw_user/api/v1/favorites"
)

type favoritesRepo struct {
	data   *Data
	logger *log.Helper
}

func NewFavoritesRepo(data *Data, favoritesClient favorv1.FavoriteClient, logger log.Logger) userbiz.FavoritesRepo {
	return &favoritesRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}

}

func (f *favoritesRepo) CreateFavorites(ctx context.Context, userId int64, favoritesName string, isPrivate int32, description string) error {
	_, err := f.data.favoritesClient.CreateFavorites(ctx, &favorv1.CreateFavoritesReq{
		UserId:        userId,
		FavoritesName: favoritesName,
		IsPrivate:     isPrivate,
		Description:   description,
	})
	return err
}

func (f *favoritesRepo) DeleteFavorites(ctx context.Context, favoritesId int64) error {
	_, err := f.data.favoritesClient.DeleteFavorites(ctx, &favorv1.DeleteFavoritesReq{
		FavoritesId: favoritesId,
	})
	return err
}

func (f *favoritesRepo) ModifyFavorites(ctx context.Context, userId int64, favoritesName *string, isPrivate *int32, description *string) error {
	_, err := f.data.favoritesClient.ModifyFavorites(ctx, &favorv1.ModifyFavoritesReq{
		FavoritesId:   userId,
		FavoritesName: favoritesName,
		IsPrivate:     isPrivate,
		Description:   description,
	})
	return err
}
