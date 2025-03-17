package data

import (
	"context"
	stderr "errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"vw_user/internal/biz"
	"vw_user/internal/data/dal/model"
)

type favoritesRepo struct {
	logger *log.Helper
	data   *Data
}

func NewFavoritesRepo(data *Data, logger log.Logger) biz.FavoritesRepo {
	return &favoritesRepo{
		logger: log.NewHelper(logger),
		data:   data,
	}
}

func (f *favoritesRepo) CheckFavoritesName(ctx context.Context, userId int64, favoritesName string) (bool, error) {
	favorites := getQuery(ctx).Favorite
	favoritesDo := favorites.WithContext(ctx)
	count, err := favoritesDo.Where(favorites.UserID.Eq(userId), favorites.FavoritesName.Eq(favoritesName)).Count()

	// if error is "record not found", set err to nil
	if err != nil && stderr.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return count == 0, err
}

func (f *favoritesRepo) CreateFavorites(ctx context.Context, favorite *model.Favorite) error {
	favorites := getQuery(ctx).Favorite
	favoritesDo := favorites.WithContext(ctx)
	return favoritesDo.Create(favorite)
}

func (f *favoritesRepo) DeleteFavorites(ctx context.Context, favoritesId int64) error {
	favorites := getQuery(ctx).Favorite
	favoritesDo := favorites.WithContext(ctx)
	_, err := favoritesDo.Where(favorites.FavoritesID.Eq(favoritesId)).Delete(&model.Favorite{})
	return err
}

func (f *favoritesRepo) CheckIfFavoritesEmpty(ctx context.Context, favoritesId int64) (bool, error) {
	favoritesVideo := getQuery(ctx).FavoriteVideo
	favoritesVideoDo := favoritesVideo.WithContext(ctx)
	count, err := favoritesVideoDo.Where(favoritesVideo.FavoritesID.Eq(favoritesId)).Count()

	// if error is "record not found", set err to nil
	if err != nil && stderr.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return count == 0, err
}

func (f *favoritesRepo) UpdateFavorites(ctx context.Context, favorite map[string]any) error {
	favorites := getQuery(ctx).Favorite
	favoritesDo := favorites.WithContext(ctx)
	_, err := favoritesDo.Debug().Where(favorites.FavoritesID.Eq(favorite["favorites_id"].(int64))).Updates(favorite)
	return err
}

func (f *favoritesRepo) GetFavoritesMetadata(ctx context.Context, favoritesId int64) (*model.Favorite, error) {
	favorites := getQuery(ctx).Favorite
	favoritesDo := favorites.WithContext(ctx)
	ret, err := favoritesDo.Where(favorites.FavoritesID.Eq(favoritesId)).First()
	if err != nil {
		return nil, err
	}
	return ret, nil
}
