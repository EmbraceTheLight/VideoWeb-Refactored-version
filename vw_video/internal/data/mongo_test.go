package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/spewerspew/spew"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"util/dbutil/mgutil"
	"vw_video/internal/biz"
	"vw_video/internal/conf"
	"vw_video/internal/data/dal/model"
)

func TestMongo(t *testing.T) {
	m := getMongoDB()

	ctx := context.Background()
	idxName, err := m.collection(uv_status).Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "video_id", Value: 1},
			{Key: "user_id", Value: 1},
		},
		Options: options.Index().SetName("video_user_id_idx").SetUnique(true),
	},
	)
	require.NoError(t, err)
	fmt.Println(idxName)
	cursor, err := m.collection(uv_status).Indexes().List(ctx)
	require.NoError(t, err)
	var results []bson.M
	err = cursor.All(ctx, &results)
	require.NoError(t, err)
	spew.Dump(results)

}

func TestMongoUpsertOne(t *testing.T) {
	m := getMongoDB()
	ctx := context.Background()

	// upsert with bson.M
	// ! NOTE: This method to update is deprecated.
	// ! Because the UpsertOne method accept an interface mgutil.MongoData as the data parameter, rather than interface{}
	//userID := "123456"
	//videoID := "789012"
	//status := biz.UpvoteStatus | biz.FavoriteStatus | biz.ShareStatus | biz.ThrowShellStatus
	//filter := mgutil.NewBsonM("user_id", userID, "video_id", videoID)
	//data := mgutil.NewBsonD("status", status)
	//err := m.UpsertOne(ctx, uv_status, filter, data)
	//require.NoError(t, err)

	// upsert with struct
	userID2 := int64(123456)
	videoID2 := int64(654321)
	uvs := &model.UserVideoStatus{
		Status: biz.UpvoteStatus | biz.FavoriteStatus,
	}
	filter2 := mgutil.NewBsonM("user_id", userID2, "video_id", videoID2)
	err := m.UpsertOne(ctx, uv_status, filter2, uvs)
	require.NoError(t, err)
}

func TestMongoFindOne(t *testing.T) {
	m := getMongoDB()
	ctx := context.Background()

	// CASE 1: find with bson.M
	userID := "123456"
	videoID := "789012"
	filter := mgutil.NewBsonM("user_id", userID, "video_id", videoID)
	ret := m.collection(uv_status).FindOne(ctx, filter)
	require.NoError(t, ret.Err())

	res1 := bson.D{}
	require.NoError(t, ret.Decode(&res1))
	spew.Dump(res1)

	// CASE 2: find with struct
	filter2 := &model.UserVideoStatus{
		UserID:  int64(123456),
		VideoID: int64(654321),
	}
	//userID2 := int64(123456)
	//videoID2 := int64(654321)
	//filter2, err := mgutil.NewBsonM("user_id", userID2, "video_id", videoID2)
	//require.NoError(t, err)

	res, err := m.FindOne(ctx, uv_status, filter2)
	require.NoError(t, err)
	spew.Dump(res)
}

func getMongoDB() *MongoDB {
	c := config.New(
		config.WithSource(
			file.NewSource("../../configs/config.yaml"),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	m := NewMongo(bc.Data)
	return m
}
