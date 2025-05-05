package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"util/dbutil/mysqlutil"
	"util/helper"
	"vw_user/internal/data/dal/model"
	"vw_user/internal/domain"
	"vw_user/internal/pkg/ecode/errdef"
)

type FollowRepo interface {
	// CreateFollowRecord inserts a new follow record into the db.
	CreateFollowRecord(ctx context.Context, follow *model.UserFollow) error

	// DeleteFollowRecord deletes a follow record from the db.
	DeleteFollowRecord(ctx context.Context, follow *model.UserFollow) error
}

type FollowUsecase struct {
	repo     FollowRepo
	infoRepo UserInfoRepo
	tx       mysqlutil.Transaction
	log      *log.Helper
}

func NewFollowUseCase(repo FollowRepo, infoRepo UserInfoRepo, tx mysqlutil.Transaction, logger log.Logger) *FollowUsecase {
	return &FollowUsecase{
		repo:     repo,
		infoRepo: infoRepo,
		tx:       tx,
		log:      log.NewHelper(logger),
	}
}

// Follow follows a userbiz
func (uc *FollowUsecase) Follow(ctx context.Context, followerID, followeeID, followListId int64) error {
	err := uc.tx.WithTx(ctx, func(context.Context) error {
		/* Follow logic */
		// 1.1 Insert a new follow record into the user_follows table.

		err := uc.repo.CreateFollowRecord(ctx, &model.UserFollow{
			FollowlistID: followListId,
			FollowUserID: followeeID,
			UserID:       followerID,
		})
		if err != nil {
			return err
		}

		// 1.2 Update the follower's followee count.
		err = uc.infoRepo.UpdateCntFollows(ctx, followerID, 1)
		if err != nil {
			return err
		}

		// 1.3 Update the followee's follower count.
		err = uc.infoRepo.UpdateCntFans(ctx, followeeID, 1)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return helper.HandleError(errdef.ErrFollowOtherUserFailed, err)
	}
	return nil
}

func (uc *FollowUsecase) Unfollow(ctx context.Context, followerID, followeeID int64) error {
	err := uc.tx.WithTx(ctx, func(context.Context) error {
		/* Follow logic */
		// 1.1 Insert a new follow record into the user_follows table.
		err := uc.repo.DeleteFollowRecord(ctx, &model.UserFollow{
			FollowUserID: followeeID,
			UserID:       followerID,
		})
		if err != nil {
			return err
		}

		// 1.2 Update the follower's followee count.
		err = uc.infoRepo.UpdateCntFollows(ctx, followerID, -1)
		if err != nil {
			return err
		}

		// 1.3 Update the followee's follower count.
		err = uc.infoRepo.UpdateCntFans(ctx, followeeID, -1)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return helper.HandleError(errdef.ErrUnfollowOtherUserFailed, err)
	}
	return nil
}

// GetFolloweesInfo Gets followees' summary info by userbiz id and follow list id.
func (uc *FollowUsecase) GetFolloweesInfo(ctx context.Context, userId int64, followListId int64, pageNum int32, pageSize int32) ([]*domain.UserSummary, error) {
	return uc.infoRepo.GetUserFolloweesByUserIDFollowListID(ctx, userId, followListId, pageNum, pageSize)
}
