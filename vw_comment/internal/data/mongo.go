package data

import (
	"context"
	stderr "errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"util/dbutil/mgutil"
	"vw_comment/internal/data/dal/model"
)

const (
	uc_status = "user_comment_upvoted" // user-comment status collection
)

type MongoDB struct {
	db          string
	mongoClient *mongo.Client
}

func (m *MongoDB) Database() *mongo.Database {
	return m.mongoClient.Database(m.db)
}

func (m *MongoDB) Collection(c string) *mongo.Collection {
	return m.Database().Collection(c)
}

func (m *MongoDB) InsertOne(ctx context.Context, data mgutil.MongoData) error {
	return mgutil.InsertOne(ctx, m.Collection(uc_status), data.GetInsertData())
}

func (m *MongoDB) UpsertOne(ctx context.Context, filter any, data mgutil.MongoData) error {
	updateOptions := options.Update().SetUpsert(true)
	return mgutil.UpdateOne(ctx, m.Collection(uc_status), filter, data.GetUpsertData(), updateOptions)
}

func (m *MongoDB) UpdateOne(ctx context.Context, filter, data mgutil.MongoData, opts ...*options.UpdateOptions) error {
	return mgutil.UpdateOne(ctx, m.Collection(uc_status), filter, data.GetUpdateData(), opts...)
}

func (m *MongoDB) FindOne(ctx context.Context, filter any, opts ...*options.FindOneOptions) (*model.UserCommentUpvoted, error) {
	res, err := mgutil.FindOne(ctx, m.Collection(uc_status), filter, &model.UserCommentUpvoted{}, opts...)
	if err != nil {
		return nil, err
	}
	userVideo, ok := res.(*model.UserCommentUpvoted)
	if !ok {
		return nil, stderr.New("FindOne: result is not a UserVideoStatus")
	}
	return userVideo, err
}

func (m *MongoDB) DeleteOne(ctx context.Context, filter any, opts ...*options.DeleteOptions) error {
	_, err := mgutil.DeleteOne(ctx, m.Collection(uc_status), filter, opts...)
	return err
}
