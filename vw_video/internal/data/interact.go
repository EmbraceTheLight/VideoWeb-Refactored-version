package data

import (
	"context"
	"database/sql"
	stderr "errors"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"strings"
	"util/dbutil"
	"util/dbutil/mgutil"
	interactv1 "vw_video/api/v1/interact"
	"vw_video/internal/biz"
	"vw_video/internal/data/dal/model"
	"vw_video/internal/data/dal/query"
)

type interactRepo struct {
	data *Data
	log  *log.Helper
}

func NewInteractRepo(data *Data, logger log.Logger) biz.InteractRepo {
	return &interactRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (i *interactRepo) AddVideoUpvote(ctx context.Context, tx *sql.Tx, videoID int64) error {
	db, videoModel, err := getGormDBAndVideoModel(tx, videoID)
	if err != nil {
		return err
	}

	err = db.Model(videoModel).
		Where("video_id = ?", videoID).
		// * Use Updates, instead of Update().Update()....,
		// * because multi Update will assemble multi sql statement,
		// * which uses the SAME optimistic-lock version as where condition,
		// * which will cause optimistic-lock conflict and could only update one column.
		// * So, use Updates to avoid this problem.
		// * Besides, I use map as argument of Updates, rather than struct, because when the struct filed is zero-value, gorm will not update it.
		// * Use map[string]interface{} can update zero-value field.
		// * In this method, use struct is okay, because AddVideoUpvote will not meet zero-value field problem.
		Updates(map[string]interface{}{
			"likes": videoModel.Likes + 1,
			"hot":   videoModel.Hot + HotEachUpvote,
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *interactRepo) DecrementVideoUpvote(ctx context.Context, tx *sql.Tx, videoID int64) error {
	db, videoModel, err := getGormDBAndVideoModel(tx, videoID)
	if err != nil {
		return err
	}

	err = db.Model(videoModel).
		Where("video_id = ?", videoID).
		Updates(map[string]interface{}{
			"likes": videoModel.Likes - 1,
			"hot":   videoModel.Hot - HotEachUpvote,
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *interactRepo) AddVideoCntFavorite(ctx context.Context, tx *sql.Tx, videoID int64) error {
	db, videoModel, err := getGormDBAndVideoModel(tx, videoID)
	if err != nil {
		return err
	}

	err = db.Model(videoModel).Debug().
		Where("video_id = ?", videoID).
		Updates(map[string]interface{}{
			"cnt_favorited": videoModel.CntFavorited + 1,
			"hot":           videoModel.Hot + HotEachFavorite,
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *interactRepo) DecrementVideoCntFavorite(ctx context.Context, tx *sql.Tx, videoID int64) error {
	db, videoModel, err := getGormDBAndVideoModel(tx, videoID)
	if err != nil {
		return err
	}

	err = db.Model(videoModel).
		Where("video_id = ?", videoID).
		Updates(map[string]interface{}{
			"cnt_favorited": videoModel.CntFavorited - 1,
			"hot":           videoModel.Hot - HotEachFavorite,
		}).Error

	if err != nil {
		return err
	}
	return nil
}

func (i *interactRepo) AddUserFavoriteVideo(ctx context.Context, tx *sql.Tx, favoritesID, videoID int64) error {
	db, _, err := getGormDBAndVideoModel(tx, videoID)
	if err != nil {
		return err
	}
	return db.Model(&model.FavoriteVideo{}).Create(&model.FavoriteVideo{
		FavoritesID: favoritesID,
		VideoID:     videoID,
	}).Error
}

func (i *interactRepo) DeleteUserFavoriteVideo(ctx context.Context, tx *sql.Tx, favoritesID int64, videoID int64) error {
	favoriteVideo := query.FavoriteVideo
	db, _, err := getGormDBAndVideoModel(tx, videoID)
	if err != nil {
		return err
	}
	return db.Model(&model.FavoriteVideo{}).Where(favoriteVideo.FavoritesID.Eq(favoritesID)).
		Unscoped().
		Delete(&model.FavoriteVideo{
			FavoritesID: favoritesID,
			VideoID:     videoID,
		}).Error
}

func (i *interactRepo) UpdateVideoShells(ctx context.Context, tx *sql.Tx, videoID int64, countShells int32) error {
	db, videoModel, err := getGormDBAndVideoModel(tx, videoID)
	if err != nil {
		return err
	}

	err = db.Model(videoModel).
		Where("video_id = ?", videoID).
		Updates(map[string]interface{}{
			"shells": videoModel.Shells + int64(countShells),
			"hot":    videoModel.Hot + int64(HotEachShell*countShells),
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *interactRepo) ShareVideo(ctx context.Context, videoID int64) string {
	return ""
}

func (i *interactRepo) AddVideoCntBarrage(ctx context.Context /*tx *sql.Tx,*/, videoID int64) error {
	//db, videoModel, err := getGormDBAndVideoModel(tx, videoID)
	//if err != nil {
	//	return err
	//}
	//err = db.Model(videoModel).
	//	Where("video_id = ?", videoID).
	//	Update("cnt_barrages", videoModel.CntBarrages+1).Update("hot", videoModel.Hot+HotEachBarrage).Error
	//if err != nil {
	//	return err
	//}
	//return nil
	video := query.Video
	videoDo, videoModel, err := addVideoModel(ctx, videoID)
	if err != nil {
		return err
	}
	_, err = videoDo.Where(video.VideoID.Eq(videoID)).Updates(&model.Video{
		CntBarrages: videoModel.CntBarrages + 1,
		Hot:         videoModel.Hot + HotEachBarrage,
	})
	if err != nil {
		return err
	}
	return nil
}

func (i *interactRepo) AddVideoBarrageUpvote(ctx context.Context, tx *sql.Tx, barrageID int64) error {
	db, err := dbutil.SqlTxToGormDB(tx)
	if err != nil {
		return err
	}

	var barrage model.Barrage
	err = db.Where("barrage_id = ?", barrageID).First(&barrage).Error
	if err != nil {
		return err
	}

	err = db.Model(&barrage).
		Where("barrage_id = ?", barrageID).
		Update("likes", barrage.Likes+1).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *interactRepo) DecrementVideoBarrageUpvote(ctx context.Context, tx *sql.Tx, barrageID int64) error {
	db, err := dbutil.SqlTxToGormDB(tx)
	if err != nil {
		return err
	}

	var barrage model.Barrage
	err = db.Where("barrage_id = ?", barrageID).First(&barrage).Error
	if err != nil {
		return err
	}

	err = db.Model(&barrage).
		Where("barrage_id = ?", barrageID).
		Update("likes", barrage.Likes-1).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *interactRepo) AddBarrage(ctx context.Context /*tx *sql.Tx,*/, info *interactv1.VideoSendBarrageReq) error {
	//db, err := dbutil.SqlTxToGormDB(tx)
	//if err != nil {
	//	return err
	//}
	//
	//return db.Model(&model.Barrage{}).Create(&model.Barrage{
	//	UserID:  info.PublisherId,
	//	VideoID: info.VideoId,
	//	Content: info.Content,
	//	Hour:    info.Hour,
	//	Minute:  info.Minute,
	//	Second:  info.Second,
	//	Color:   info.Color,
	//}).Error
	hms := strings.Split(info.Time, ":")
	if len(hms) != 3 {
		return stderr.New("invalid time format")
	}
	barrage := getQuery(ctx).Barrage
	barrageDo := barrage.WithContext(ctx)
	return barrageDo.Create(&model.Barrage{
		PublisherID: info.PublisherId,
		VideoID:     info.VideoId,
		Content:     info.Content,
		Hour:        hms[0],
		Minute:      hms[1],
		Second:      hms[2],
		Color:       info.Color,
	})
}

func (i *interactRepo) GetPublisherIdByVideoId(ctx context.Context, videoID int64) (int64, error) {
	_, video, err := addVideoModel(ctx, videoID)
	if err != nil {
		return 0, err
	}
	return video.PublisherID, nil
}

func (i *interactRepo) GetUserVideoStatus(ctx context.Context, userID, videoID int64) (status int64, shellsCount int32, err error) {
	filter := mgutil.NewBsonM("user_id", userID, "video_id", videoID)
	res, err := i.data.mongo.FindOne(ctx, uv_status, filter)
	if err != nil {
		if stderr.Is(err, mongo.ErrNoDocuments) {
			err = i.data.mongo.InsertOne(ctx, ub_status, &model.UserVideoStatus{
				UserID:  userID,
				VideoID: videoID,
				Status:  0,
			})
			if err != nil {
				return -1, -1, err
			}
			return 0, 0, nil
		}
		return -1, -1, err
	}
	return res.Status, int32(res.ShellsCount), nil
}

func (i *interactRepo) GetUserBarrageStatus(ctx context.Context, userID int64, barrageID int64) (int64, error) {
	filter := mgutil.NewBsonM("user_id", userID, "barrage_id", barrageID)
	res, err := i.data.mongo.FindOne(ctx, ub_status, filter)
	if err != nil {
		if stderr.Is(err, mongo.ErrNoDocuments) {
			err = i.data.mongo.InsertOne(ctx, ub_status, &model.UserBarrageStatus{
				UserID:    userID,
				BarrageID: barrageID,
				Status:    0,
			})
			if err != nil {
				return -1, err
			}
			return 0, nil
		}
		return -1, err
	}
	return res.Status, nil
}

func (i *interactRepo) SetUserVideoStatus(ctx context.Context, userID, videoID int64, status int64) error {
	filter := mgutil.NewBsonM("user_id", userID, "video_id", videoID)
	data := &model.UserVideoStatus{
		UserID:  userID,
		VideoID: videoID,
		Status:  status,
	}
	return i.data.mongo.UpsertOne(ctx, uv_status, filter, data)
}

// GetSqlTx is a helper function to get *sql.Tx from *gorm.DB.
// This is used to be the DTM barrier's Call method's argument.
func (i *interactRepo) GetSqlTx() *sql.Tx {
	tx := i.data.mysql.Begin()
	return tx.Statement.ConnPool.(*sql.Tx)
}

// GetMongoClient is a helper function to get *mongo.Client from *mongo.Data.
// This is used to be the DTM barrier's MongoCall method's argument.
func (i *interactRepo) GetMongoClient() *mongo.Client {
	return i.data.mongo.mongoClient
}

func (i *interactRepo) GetDB() (db string) {
	return i.data.mongo.db
}

func getGormDBAndVideoModel(tx *sql.Tx, videoID int64) (*gorm.DB, *model.Video, error) {
	db, err := dbutil.SqlTxToGormDB(tx)
	if err != nil {
		return nil, nil, err
	}

	var videoModel model.Video
	err = db.Where("video_id = ?", videoID).First(&videoModel).Error
	if err != nil {
		return nil, nil, err
	}
	return db, &videoModel, nil
}
