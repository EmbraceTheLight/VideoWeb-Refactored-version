package biz

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/dtm/client/dtmcli"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	_ "github.com/dtm-labs/dtmdriver-kratos"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"strconv"
	"util"
	"util/dbutil/mgutil"
	"util/helper"
	interactv1 "vw_video/api/v1/interact"
	"vw_video/internal/data/dal/model"
	"vw_video/internal/pkg/ecode/errdef"
)

type InteractRepo interface {
	/* Used for dtm */
	GetSqlTx() *sql.Tx
	GetMongoClient() *mongo.Client
	GetDB() (db string)

	/* MySQL */
	GetPublisherIdByVideoId(ctx context.Context, videoID int64) (int64, error)
	AddVideoUpvote(ctx context.Context, tx *sql.Tx, videoID int64) error
	DecrementVideoUpvote(ctx context.Context, tx *sql.Tx, videoID int64) error

	AddVideoCntFavorite(ctx context.Context, tx *sql.Tx, videoID int64) error
	DecrementVideoCntFavorite(ctx context.Context, tx *sql.Tx, videoID int64) error
	AddUserFavoriteVideo(ctx context.Context, tx *sql.Tx, favoritesID, videoID int64) error
	DeleteUserFavoriteVideo(ctx context.Context, tx *sql.Tx, favoritesID int64, videoID int64) error

	UpdateVideoShells(ctx context.Context, tx *sql.Tx, videoID int64, countShells int32) error
	ShareVideo(ctx context.Context, videoID int64) string
	AddVideoCntBarrage(ctx context.Context, videoID int64) error
	AddBarrage(ctx context.Context, into *interactv1.VideoSendBarrageReq) error
	AddVideoBarrageUpvote(ctx context.Context, tx *sql.Tx, barrageID int64) error
	DecrementVideoBarrageUpvote(ctx context.Context, tx *sql.Tx, barrageID int64) error

	/* MongoDB */
	GetUserVideoStatus(ctx context.Context, userID, videoID int64) (status int64, shellsCount int32, err error)
	GetUserBarrageStatus(ctx context.Context, userID int64, barrageID int64) (int64, error)
	SetUserVideoStatus(ctx context.Context, userID, videoID int64, status int64) error
}

type InteractUsecase struct {
	repo          InteractRepo
	videoinfoRepo VideoInfoRepo
	tx            util.Transaction
	log           *log.Helper
}

func NewInteractUseCase(repo InteractRepo, infoRepo VideoInfoRepo, tx util.Transaction, logger log.Logger) *InteractUsecase {
	return &InteractUsecase{
		repo:          repo,
		videoinfoRepo: infoRepo,
		tx:            tx,
		log:           log.NewHelper(logger),
	}
}

func (u *InteractUsecase) UpvoteVideo(ctx context.Context, videoID int64, upvote int32) error {
	u.log.Info("DTM-UpvoteVideo start...")
	defer u.log.Info("DTM-UpvoteVideo end...")

	tx := u.repo.GetSqlTx()
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}

	err = barrier.Call(tx, func(tx *sql.Tx) error {
		if upvote == 1 {
			return u.repo.AddVideoUpvote(ctx, tx, videoID)
		} else {
			return u.repo.DecrementVideoUpvote(ctx, tx, videoID)
		}
	})
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}
	return nil
}

func (u *InteractUsecase) FavoriteVideo(ctx context.Context, videoID, favoriteID int64, favorite int32) error {
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}

	err = barrier.Call(u.repo.GetSqlTx(), func(tx *sql.Tx) error {
		var err error

		// 1. update video's favorite count.
		// 2. update user-favorite relationship.
		if favorite == 1 {
			err = u.repo.AddVideoCntFavorite(ctx, tx, videoID)
			if err != nil {
				return err
			}

			err = u.repo.AddUserFavoriteVideo(ctx, tx, favoriteID, videoID)
			if err != nil {
				return err
			}
		} else {
			err = u.repo.DecrementVideoCntFavorite(ctx, tx, videoID)
			if err != nil {
				return err
			}

			err = u.repo.DeleteUserFavoriteVideo(ctx, tx, favoriteID, videoID)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}
	return nil
}

func (u *InteractUsecase) UpdateVideoShells(ctx context.Context, videoID int64, countShells int32) error {
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}

	err = barrier.Call(u.repo.GetSqlTx(), func(tx *sql.Tx) error {
		// update video's shells(and video's hot)
		err := u.repo.UpdateVideoShells(ctx, tx, videoID, countShells)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}
	return nil
}

func (u *InteractUsecase) ShareVideo(ctx context.Context, videoID int64) (string, error) {
	return "/api/v1/video/" + strconv.FormatInt(videoID, 10) + "/info", nil
}

func (u *InteractUsecase) SendBarrage(ctx context.Context, info *interactv1.VideoSendBarrageReq) error {
	//tx := u.repo.GetSqlTx()
	//barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	//if err != nil {
	//	return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	//}
	//err = barrier.Call(tx, func(tx *sql.Tx) error {
	//	// 1. insert barrage into barrages table
	//	err := u.repo.AddBarrage(ctx, tx, info)
	//	if err != nil {
	//		return err
	//	}
	//
	//	// 2. add video's cnt_barrage
	//	err = u.repo.AddVideoCntBarrage(ctx, tx, info.VideoId)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//})
	//if err != nil {
	//	return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	//}
	//return nil
	err := u.tx.WithTx(ctx, func(ctx context.Context) error {
		// 1. insert barrage into barrages table
		err := u.repo.AddBarrage(ctx, info)
		if err != nil {
			return err
		}

		// 2. add video's cnt_barrage
		err = u.repo.AddVideoCntBarrage(ctx, info.VideoId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return helper.HandleError(errdef.ErrSendBarrageFailed, err)
	}
	return nil
}

func (u *InteractUsecase) UpvoteBarrage(ctx context.Context, barrageID int64, upvote int32) error {
	tx := u.repo.GetSqlTx()
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}

	err = barrier.Call(tx, func(tx *sql.Tx) error {
		// upvote == 1 , which indicates that the user wants to upvote the barrage,
		// otherwise, the user wants to cancel the upvote.
		if upvote == 1 {
			return u.repo.AddVideoBarrageUpvote(ctx, tx, barrageID)
		} else {
			return u.repo.DecrementVideoBarrageUpvote(ctx, tx, barrageID)
		}
	})
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}
	return nil
}

func (u *InteractUsecase) GetUserVideoStatus(ctx context.Context, userID, videoID int64) (int64, int32, error) {
	status, shellsCount, err := u.repo.GetUserVideoStatus(ctx, userID, videoID)
	if err != nil {
		return -1, -1, err
	}
	return status, shellsCount, nil
}
func (u *InteractUsecase) GetUserBarrageStatus(ctx context.Context, userID, barrageID int64) (int64, error) {
	status, err := u.repo.GetUserBarrageStatus(ctx, userID, barrageID)
	if err != nil {
		return -1, err
	}
	return status, nil
}

func (u *InteractUsecase) SetUserVideoStatus(ctx context.Context, userID, videoID int64, req *interactv1.SetUserVideoStatusReq) error {
	status := req.OriginStatus
	statuses := req.Statuses
	setStatus(&status, UpvoteStatus, statuses.IsUpvoted)
	setStatus(&status, FavoriteStatus, statuses.IsFavorited)
	setStatus(&status, ShareStatus, statuses.IsShared)
	setStatus(&status, ThrowShellStatus, statuses.IsThrownShells)
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}

	err = barrier.MongoCall(u.repo.GetMongoClient(), func(sc mongo.SessionContext) error {
		// 1. init mongoDB collection
		cli := sc.Client()
		db, collection := u.repo.GetDB(), uv_status
		coll := mgutil.NewCollection(cli, db, collection)

		// 2. generate filter and update data
		filter := mgutil.NewBsonM("user_id", userID, "video_id", videoID)
		data := &model.UserVideoStatus{
			Status:      status,
			ShellsCount: int64(req.ShellsCount),
		}
		opts := mgutil.NewUpdateOptions().SetUpsert(true)
		err = mgutil.UpdateOne(sc, coll, filter, data.GetUpsertData(), opts)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}
	return nil
}
func (u *InteractUsecase) SetUserVideoStatusRevert(ctx context.Context, userID, videoID int64, req *interactv1.SetUserVideoStatusReq) error {
	status := req.OriginStatus
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}

	err = barrier.MongoCall(u.repo.GetMongoClient(), func(sc mongo.SessionContext) error {
		// 1. init mongoDB collection
		cli := sc.Client()
		db, collection := u.repo.GetDB(), uv_status
		coll := mgutil.NewCollection(cli, db, collection)

		// 2. generate filter, update data and `upsert` options
		filter := mgutil.NewBsonM("user_id", userID, "video_id", videoID)
		data := &model.UserVideoStatus{
			Status:      status,
			ShellsCount: int64(req.ShellsCount),
		}
		opts := mgutil.NewUpdateOptions().SetUpsert(true)
		err = mgutil.UpdateOne(sc, coll, filter, data.GetUpsertData(), opts)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}
	return nil
}

func (u *InteractUsecase) SetUserBarrageStatus(ctx context.Context, userID int64, barrageID int64, req *interactv1.SetUserBarrageStatusReq) error {
	status := req.OriginStatus
	statuses := req.Statuses
	setStatus(&status, UpvoteStatus, statuses.IsUpvoted)
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}

	err = barrier.MongoCall(u.repo.GetMongoClient(), func(sc mongo.SessionContext) error {
		// 1. init mongoDB collection
		cli := sc.Client()
		db, collection := u.repo.GetDB(), ub_status
		coll := mgutil.NewCollection(cli, db, collection)

		// 2. generate filter and update data
		filter := mgutil.NewBsonM("user_id", userID, "barrage_id", barrageID)
		data := &model.UserBarrageStatus{
			Status: status,
		}
		opts := mgutil.NewUpdateOptions().SetUpsert(true)
		err = mgutil.UpdateOne(sc, coll, filter, data.GetUpsertData(), opts)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}
	return nil
}
func (u *InteractUsecase) SetUserBarrageStatusRevert(ctx context.Context, userID int64, barrageID int64, req *interactv1.SetUserBarrageStatusReq) error {
	status := req.OriginStatus
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}

	err = barrier.MongoCall(u.repo.GetMongoClient(), func(sc mongo.SessionContext) error {
		// 1. init mongoDB collection
		cli := sc.Client()
		db, collection := u.repo.GetDB(), ub_status
		coll := mgutil.NewCollection(cli, db, collection)

		// 2. generate filter, update data and `upsert` options
		filter := mgutil.NewBsonM("user_id", userID, "barrage_id", barrageID)
		data := &model.UserBarrageStatus{
			Status: status,
		}
		opts := mgutil.NewUpdateOptions().SetUpsert(true)
		err = mgutil.UpdateOne(sc, coll, filter, bson.D{{"$set", data}}, opts)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return helper.HandleGrpcError(codes.Aborted, dtmcli.ResultFailure, err)
	}
	return nil
}

func setStatus(originStatus *int64, status int64, statusSwitch int32) {
	if statusSwitch == 1 {
		*originStatus |= status // set status
	} else if statusSwitch == -1 {
		*originStatus &= ^status // clear status
	}
}
