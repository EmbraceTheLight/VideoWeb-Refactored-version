package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserCommentUpvoted struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    int64              `bson:"user_id,omitempty"`
	CommentID int64              `bson:"comment_id,omitempty"`
}

func (u *UserCommentUpvoted) GetInsertData() any {
	return u
}

func (u *UserCommentUpvoted) GetUpdateData() any {
	return bson.D{{"$set", u}}
}

func (u *UserCommentUpvoted) GetUpsertData() any {
	return bson.D{
		{"$setOnInsert", bson.D{
			{"user_id", u.UserID},
			{"barrage_id", u.CommentID},
		}},
	}
}
