package model

import (
	"go.mongodb.org/mongo-driver/bson"
)

type VideoSummary struct {
	VideoId       int64  `bson:"video_id" gorm:"column:video_id;type:bigint;primaryKey;comment:视频ID" json:"video_id"`
	CntBarrages   int64  `bson:"cnt_barrages"  gorm:"column:cnt_barrages;type:int unsigned;comment:视频弹幕数" json:"cnt_barrages"`
	Title         string `bson:"title" gorm:"column:title;type:varchar(100);not null;comment: 视频标题" json:"title"`
	Duration      string `bson:"duration" gorm:"column:duration;type:varchar(10);not null;comment:视频时长" json:"duration"`
	PublisherName string `bson:"publisher_name" gorm:"column:publisher_name;type:varchar(200);not null;comment:上传者用户名" json:"publisher_name"`
	CntViewed     int64  `bson:"cnt_viewed" gorm:"column:cnt_viewed;type:bigint;comment:视频观看数（点开就算看）" json:"cnt_viewed"`
	CoverPath     string `bson:"cover_path" gorm:"column:size;type:bigint;not null;comment:视频文件大小" json:"size"`
}
type UserVideoHistory struct {
	//ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       int64         `bson:"user_id,omitempty"`
	VideoID      int64         `bson:"video_id,omitempty"`
	Timestamp    int64         `bson:"timestamp,omitempty"`
	VideoSummary *VideoSummary `bson:"video_summary,omitempty"`
}

func (u *UserVideoHistory) GetInsertData() any {
	return u
}

func (u *UserVideoHistory) GetUpdateData() any {
	return bson.D{{"$set", u}}
}

func (u *UserVideoHistory) GetUpsertData() any {
	return bson.D{
		{"$set", bson.D{
			{"timestamp", u.Timestamp},
		}},
		{"$setOnInsert", bson.D{
			{"user_id", u.UserID},
			{"video_id", u.VideoID},
			{"video_summary", u.VideoSummary},
		}},
	}
}
