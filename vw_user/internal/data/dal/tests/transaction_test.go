package tests

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"util/snowflake"
	"vw_user/internal/data/dal/model"
	"vw_user/internal/data/dal/query"
)

func TestTransactionManually(t *testing.T) {
	var err error

	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			fmt.Println("rollback")
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	require.NoError(t, tx.Error)
	// create user's default follow_list and favorites
	userID := int64(987654321)
	defaultFavorites := &model.UserFavorite{
		FavoriteID:   snowflake.GetID(),
		UserID:       userID,
		FavoriteName: "默认收藏夹",
		Description:  "",
		IsPrivate:    1,
	}
	privateFavorites := &model.UserFavorite{
		FavoriteID:   snowflake.GetID(),
		UserID:       userID,
		FavoriteName: "私密收藏夹",
		Description:  "",
		IsPrivate:    -1,
	}
	userLevel := &model.UserLevel{
		UserID: userID,
	}
	defaultFollowList := &model.FollowList{
		ListID:   snowflake.GetID(),
		UserID:   userID,
		ListName: "默认关注列表",
	}
	err = tx.UserLevel.Create(userLevel)
	if err != nil {
		return
	}
	err = tx.UserFavorite.Create(defaultFavorites, privateFavorites)
	if err != nil {
		return
	}
	err = tx.FollowList.Create(defaultFollowList)
	if err != nil {
		return
	}
	tx.Commit()
}

func TestTransactionAutomatically(t *testing.T) {
	userID := int64(123456789)
	defaultFavorites := &model.UserFavorite{
		FavoriteID:   snowflake.GetID(),
		UserID:       userID,
		FavoriteName: "默认收藏夹",
		Description:  "",
		IsPrivate:    1,
	}
	privateFavorites := &model.UserFavorite{
		FavoriteID:   snowflake.GetID(),
		UserID:       userID,
		FavoriteName: "私密收藏夹",
		Description:  "",
		IsPrivate:    -1,
	}
	userLevel := &model.UserLevel{
		UserID: userID,
	}
	defaultFollowList := &model.FollowList{
		ListID:   snowflake.GetID(),
		UserID:   userID,
		ListName: "默认关注列表",
	}
	err := query.Q.Transaction(func(tx *query.Query) error {
		err := tx.UserLevel.Create(userLevel)
		if err != nil {
			return err
		}
		err = tx.UserFavorite.Create(defaultFavorites, privateFavorites)
		if err != nil {
			return err
		}
		err = tx.FollowList.Create(defaultFollowList)
		if err != nil {
			return err
		}
		return nil
	})
	require.NoError(t, err)
}
