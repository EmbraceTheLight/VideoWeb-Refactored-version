package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBarrageStatus struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    int64              `bson:"user_id,omitempty"`
	BarrageID int64              `bson:"barrage_id,omitempty"`
	Status    int64              `bson:"status"`
}

func (u *UserBarrageStatus) GetInsertData() any {
	return u
}

func (u *UserBarrageStatus) GetUpdateData() any {
	return bson.D{{"$set", u}}
}

func (u *UserBarrageStatus) GetUpsertData() any {
	return bson.D{
		{"$set", bson.D{
			{"status", u.Status}}},
		{"$setOnInsert", bson.D{
			{"user_id", u.UserID},
			{"barrage_id", u.BarrageID},
		}},
	}
}
