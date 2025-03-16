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
	ID          int64          `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                                    // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`                                    // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);index:deleted_at,priority:1;comment:删除时间（用于软删除）" json:"deleted_at"` // 删除时间（用于软删除）
	FavoritesID int64          `gorm:"column:favorites_id;type:bigint;not null;index:fav_id_video_id,priority:1" json:"favorites_id"`
	VideoID     int64          `gorm:"column:video_id;type:bigint;not null;index:fav_id_video_id,priority:2" json:"video_id"`
}

// TableName FavoriteVideo's table name
func (*FavoriteVideo) TableName() string {
	return TableNameFavoriteVideo
}
