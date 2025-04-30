package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserVideoStatus struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      int64              `bson:"user_id,omitempty"`
	VideoID     int64              `bson:"video_id,omitempty"`
	ShellsCount int64              `bson:"shells_count,omitempty"` // 贝壳最大投递数量
	Status      int64              `bson:"status"`
}

func (u *UserVideoStatus) GetInsertData() any {
	return u
}

func (u *UserVideoStatus) GetUpdateData() any {
	return bson.D{{"$set", u}}
}

func (u *UserVideoStatus) GetUpsertData() any {
	return bson.D{
		{"$set", bson.D{
			{"status", u.Status},
			{"shells_count", u.ShellsCount},
		}},
		{"$setOnInsert", bson.D{
			{"user_id", u.UserID},
			{"video_id", u.VideoID},
		}},
	}
}
