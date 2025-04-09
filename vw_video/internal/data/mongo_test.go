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
	"vw_video/internal/conf"
)

func TestMongo(t *testing.T) {
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

	ctx := context.Background()
	idxName, err := m.Collection().Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "video_id", Value: 1},
			{Key: "user_id", Value: 1},
		},
		Options: options.Index().SetName("video_user_id_idx").SetUnique(true),
	},
	)
	require.NoError(t, err)
	fmt.Println(idxName)
	cursor, err := m.Collection().Indexes().List(ctx)
	require.NoError(t, err)
	var results []bson.M
	err = cursor.All(ctx, &results)
	require.NoError(t, err)
	spew.Dump(results)

}
