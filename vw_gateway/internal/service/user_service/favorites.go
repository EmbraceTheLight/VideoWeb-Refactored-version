package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	favorv1 "vw_gateway/api/v1/user/favorites"
	"vw_gateway/internal/biz/userbiz"
)

type FavoritesService struct {
	favorv1.UnimplementedFavoriteServer
	favorites *userbiz.FavoritesUsecase
	logger    *log.Helper
}

func NewFavoritesService(favorites *userbiz.FavoritesUsecase, logger log.Logger) *FavoritesService {
	return &FavoritesService{
		favorites: favorites,
		logger:    log.NewHelper(logger),
	}
}

func (u *FavoritesService) CreateFavorites(ctx context.Context, req *favorv1.CreateFavoritesReq) (*favorv1.CreateFavoritesResp, error) {
	err := u.favorites.CreateFavorites(ctx, req.UserId, req.FavoritesName, req.IsPrivate, req.Description)
	if err != nil {
		return nil, err
	}
	return &favorv1.CreateFavoritesResp{
		StatusCode: 200,
		Msg:        "创建收藏夹成功",
	}, nil
}

func (u *FavoritesService) DeleteFavorites(ctx context.Context, req *favorv1.DeleteFavoritesReq) (*favorv1.DeleteFavoritesResp, error) {
	err := u.favorites.DeleteFavorites(ctx, req.FavoritesId)
	if err != nil {
		return nil, err
	}
	return &favorv1.DeleteFavoritesResp{
		StatusCode: 200,
		Msg:        "删除收藏夹成功",
	}, nil
}

func (u *FavoritesService) ModifyFavorites(ctx context.Context, req *favorv1.ModifyFavoritesReq) (*favorv1.ModifyFavoritesResp, error) {
	err := u.favorites.ModifyFavorites(ctx, req.FavoritesId, req.FavoritesName, req.IsPrivate, req.Description)
	if err != nil {
		return nil, err
	}
	return &favorv1.ModifyFavoritesResp{
		StatusCode: 200,
		Msg:        "修改收藏夹成功",
	}, nil
}
