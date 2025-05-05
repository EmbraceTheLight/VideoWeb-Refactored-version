package data

import (
	"context"
	stderr "errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"util/dbutil/mgutil"
	"vw_video/internal/data/dal/model"
)

const (
	uv_status  = "user_video_status"    // user-video status Collection
	ub_status  = "user_barrage_status"  // user-barrage status Collection
	uv_history = "user_video_history"   // user-video history Collection
	uc_status  = "user_comment_upvoted" // user-comment status Collection
)

type MongoDB struct {
	db          string
	mongoClient *mongo.Client
}

//// FindOne finds one document in the specified Collection that matches the filter.
//// result is a pointer to the struct that will hold the result.
//func FindOne(ctx context.Context, Collection *mongo.Collection, filter any, result any, opts ...*options.FindOneOptions) (any, error) {
//	res := Collection.FindOne(ctx, filter, opts...)
//	if err := res.Err(); err != nil {
//		return nil, err
//	}
//
//	err := res.Decode(result)
//	if err != nil {
//		return nil, err
//	}
//	return result, nil
//}
//
//// UpdateOne finds one document in the specified Collection that matches the filter.
//// result is a pointer to the struct that will hold the result.
//func UpdateOne(ctx context.Context, Collection *mongo.Collection, filter any, data any, opts ...*options.UpdateOptions) error {
//	_, err := Collection.UpdateOne(ctx, filter, bson.D{{"$set", data}}, opts...)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func (m *MongoDB) Database() *mongo.Database {
	return m.mongoClient.Database(m.db)
}

func (m *MongoDB) Collection(c string) *mongo.Collection {
	return m.Database().Collection(c)
}

func (m *MongoDB) InsertOne(ctx context.Context, collection string, data mgutil.MongoData, opts ...*options.InsertOneOptions) error {
	return mgutil.InsertOne(ctx, m.Collection(collection), data.GetInsertData(), opts...)
}

func (m *MongoDB) UpsertOne(ctx context.Context, collection string, filter any, data mgutil.MongoData) error {
	updateOptions := options.Update().SetUpsert(true)
	return mgutil.UpdateOne(ctx, m.Collection(collection), filter, data.GetUpsertData(), updateOptions)
}

func (m *MongoDB) UpdateOne(ctx context.Context, collection string, filter, data mgutil.MongoData, opts ...*options.UpdateOptions) error {
	return mgutil.UpdateOne(ctx, m.Collection(collection), filter, data.GetUpdateData(), opts...)
}

func (m *MongoDB) DeleteOne(ctx context.Context, filter any, opts ...*options.DeleteOptions) error {
	_, err := mgutil.DeleteOne(ctx, m.Collection(uc_status), filter, opts...)
	return err
}

func (m *MongoDB) FindOne(ctx context.Context, collection string, filter any, opts ...*options.FindOneOptions) (*model.UserVideoStatus, error) {
	res, err := mgutil.FindOne(ctx, m.Collection(collection), filter, &model.UserVideoStatus{}, opts...)
	if err != nil {
		return nil, err
	}
	userVideo, ok := res.(*model.UserVideoStatus)
	if !ok {
		return nil, stderr.New("FindOne: result is not a UserVideoStatus")
	}
	return userVideo, err
}
