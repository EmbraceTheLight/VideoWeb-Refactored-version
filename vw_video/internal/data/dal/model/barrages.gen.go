// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameBarrage = "barrages"

// Barrage mapped from table <barrages>
type Barrage struct {
	BarrageID   int64          `gorm:"column:barrage_id;type:bigint;primaryKey;autoIncrement:true;comment:“弹幕ID”" json:"barrage_id"`                 // “弹幕ID”
	VideoID     int64          `gorm:"column:video_id;type:bigint;not null;comment:视频id" json:"video_id"`                                            // 视频id
	PublisherID int64          `gorm:"column:publisher_id;type:bigint;comment: 发布者 id" json:"publisher_id"`                                          //  发布者 id
	Hour        string         `gorm:"column:hour;type:char(2);not null;comment:弹幕出现时间--小时" json:"hour"`                                             // 弹幕出现时间--小时
	Minute      string         `gorm:"column:minute;type:char(2);not null;comment:弹幕出现时间--分钟" json:"minute"`                                         // 弹幕出现时间--分钟
	Second      string         `gorm:"column:second;type:char(2);not null;comment:弹幕出现时间--秒" json:"second"`                                          // 弹幕出现时间--秒
	Content     string         `gorm:"column:content;type:varchar(150);not null;comment:弹幕内容" json:"content"`                                        // 弹幕内容
	Color       string         `gorm:"column:color;type:char(8);not null;comment:弹幕颜色，使用十六进制表示" json:"color"`                                        // 弹幕颜色，使用十六进制表示
	Likes       int64          `gorm:"column:likes;type:int unsigned;comment:弹幕获赞数" json:"likes"`                                                    // 弹幕获赞数
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                                            // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`                                            // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);index:barrages__index_deletedat,priority:1;comment:删除时间" json:"deleted_at"` // 删除时间
}

// TableName Barrage's table name
func (*Barrage) TableName() string {
	return TableNameBarrage
}
