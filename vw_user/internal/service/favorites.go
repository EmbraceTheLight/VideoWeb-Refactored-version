package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	favorv1 "vw_user/api/v1/favorites"
	"vw_user/internal/biz"
	"vw_user/internal/domain"
)

type FavoritesService struct {
	favorv1.UnimplementedFavoriteServer
	favorites *biz.FavoritesUsecase
	logger    *log.Helper
}

func NewFavoritesService(favorites *biz.FavoritesUsecase, logger log.Logger) *FavoritesService {
	return &FavoritesService{
		favorites: favorites,
		logger:    log.NewHelper(logger),
	}
}

func (fs *FavoritesService) CreateFavorites(ctx context.Context, req *favorv1.CreateFavoritesReq) (*favorv1.CreateFavoritesResp, error) {
	err := fs.favorites.CreateFavorites(ctx, &domain.FavoritesInfo{
		UserId:        req.UserId,
		FavoritesName: &req.FavoritesName,
		IsPrivate:     &req.IsPrivate,
		Description:   &req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &favorv1.CreateFavoritesResp{
		StatusCode: 200,
		Msg:        "创建收藏夹成功",
	}, nil
}

func (fs *FavoritesService) ModifyFavorites(ctx context.Context, req *favorv1.ModifyFavoritesReq) (*favorv1.ModifyFavoritesResp, error) {
	favorites := &domain.FavoritesInfo{FavoritesId: req.FavoritesId}
	favorites.FavoritesName = req.FavoritesName
	favorites.IsPrivate = req.IsPrivate
	favorites.Description = req.Description

	err := fs.favorites.ModifyFavorites(ctx, favorites)
	if err != nil {
		return nil, err
	}
	return &favorv1.ModifyFavoritesResp{
		StatusCode: 200,
		Msg:        "修改收藏夹信息成功",
	}, nil
}

func (fs *FavoritesService) DeleteFavorites(ctx context.Context, req *favorv1.DeleteFavoritesReq) (*favorv1.DeleteFavoritesResp, error) {
	err := fs.favorites.DeleteFavorites(ctx, req.FavoritesId)
	if err != nil {
		return nil, err
	}
	return &favorv1.DeleteFavoritesResp{
		StatusCode: 200,
		Msg:        "删除收藏夹成功",
	}, nil
}
