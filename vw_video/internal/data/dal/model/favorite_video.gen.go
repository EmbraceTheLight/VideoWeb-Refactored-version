// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFavoriteVideo = "favorite_video"

// FavoriteVideo mapped from table <favorite_video>
type FavoriteVideo struct {
	ID          int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键" json:"id"`                                                        // 主键
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                                                               // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`                                                               // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间" json:"deleted_at"`                                                               // 删除时间
	FavoritesID int64          `gorm:"column:favorites_id;type:bigint;not null;uniqueIndex:favorite_video__index_vid_fid,priority:1;comment:收藏夹id" json:"favorites_id"` // 收藏夹id
	VideoID     int64          `gorm:"column:video_id;type:bigint;not null;uniqueIndex:favorite_video__index_vid_fid,priority:2;comment:收藏的视频id" json:"video_id"`       // 收藏的视频id
}

// TableName FavoriteVideo's table name
func (*FavoriteVideo) TableName() string {
	return TableNameFavoriteVideo
}
