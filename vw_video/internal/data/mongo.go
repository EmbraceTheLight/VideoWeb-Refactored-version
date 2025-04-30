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
	uv_status  = "user_video_status"   // user-video status collection
	ub_status  = "user_barrage_status" // user-barrage status collection
	uv_history = "user_video_history"  // user-video history collection
	uc_status  = "user_comment_status" // user-comment status collection
)

type MongoDB struct {
	db          string
	mongoClient *mongo.Client
}

//// FindOne finds one document in the specified collection that matches the filter.
//// result is a pointer to the struct that will hold the result.
//func FindOne(ctx context.Context, collection *mongo.Collection, filter any, result any, opts ...*options.FindOneOptions) (any, error) {
//	res := collection.FindOne(ctx, filter, opts...)
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
//// UpdateOne finds one document in the specified collection that matches the filter.
//// result is a pointer to the struct that will hold the result.
//func UpdateOne(ctx context.Context, collection *mongo.Collection, filter any, data any, opts ...*options.UpdateOptions) error {
//	_, err := collection.UpdateOne(ctx, filter, bson.D{{"$set", data}}, opts...)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func (m *MongoDB) database() *mongo.Database {
	return m.mongoClient.Database(m.db)
}

func (m *MongoDB) collection(c string) *mongo.Collection {
	return m.database().Collection(c)
}

func (m *MongoDB) InsertOne(ctx context.Context, collection string, data mgutil.MongoData) error {
	return mgutil.InsertOne(ctx, m.collection(collection), data.GetInsertData())
}

func (m *MongoDB) UpsertOne(ctx context.Context, collection string, filter any, data mgutil.MongoData) error {
	updateOptions := options.Update().SetUpsert(true)
	return mgutil.UpdateOne(ctx, m.collection(collection), filter, data.GetUpsertData(), updateOptions)
}

func (m *MongoDB) UpdateOne(ctx context.Context, collection string, filter, data mgutil.MongoData, opts ...*options.UpdateOptions) error {
	return mgutil.UpdateOne(ctx, m.collection(collection), filter, data.GetUpdateData(), opts...)
}

func (m *MongoDB) FindOne(ctx context.Context, collection string, filter any, opts ...*options.FindOneOptions) (*model.UserVideoStatus, error) {
	res, err := mgutil.FindOne(ctx, m.collection(collection), filter, &model.UserVideoStatus{}, opts...)
	if err != nil {
		return nil, err
	}
	userVideo, ok := res.(*model.UserVideoStatus)
	if !ok {
		return nil, stderr.New("FindOne: result is not a UserVideoStatus")
	}
	return userVideo, err
}
