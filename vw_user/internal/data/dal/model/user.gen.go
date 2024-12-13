// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	CreatedAt     time.Time              `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                                    // 创建时间
	UpdatedAt     time.Time              `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`                                    // 更新时间
	DeletedAt     gorm.DeletedAt         `gorm:"column:deleted_at;type:datetime(3);index:deleted_at,priority:1;comment:删除时间（用于软删除）" json:"deleted_at"` // 删除时间（用于软删除）
	UserID        int64                  `gorm:"column:user_id;type:bigint;primaryKey" json:"user_id"`
	Username      string                 `gorm:"column:username;type:varchar(40);not null;index:username,priority:1;comment:用户名" json:"username"` // 用户名
	Password      string                 `gorm:"column:password;type:varchar(72);not null;comment:用户密码（已加密）" json:"password"`                     // 用户密码（已加密）
	Email         string                 `gorm:"column:email;type:varchar(100);not null" json:"email"`
	Signature     string                 `gorm:"column:signature;type:varchar(25);comment:用户签名，不多于25字" json:"signature"`                   // 用户签名，不多于25字
	Shells        int64                  `gorm:"column:shells;type:int unsigned;not null;default:1000;comment:拥有的贝壳数" json:"shells"`       // 拥有的贝壳数
	IsAdmin       bool                   `gorm:"column:is_admin;type:tinyint(1);comment:是否是管理员" json:"is_admin"`                           // 是否是管理员
	CntMsgNotRead int64                  `gorm:"column:cnt_msg_not_read;type:int unsigned;not null;comment:未读消息数" json:"cnt_msg_not_read"` // 未读消息数
	CntFans       int64                  `gorm:"column:cnt_fans;type:int unsigned;not null;comment:粉丝数" json:"cnt_fans"`                   // 粉丝数
	CntComments   int64                  `gorm:"column:cnt_comments;type:int unsigned;not null;comment:发表评论数" json:"cnt_comments"`         // 发表评论数
	CntLikes      int64                  `gorm:"column:cnt_likes;type:int unsigned;not null;comment:获赞数" json:"cnt_likes"`                 // 获赞数
	AvatarPath    string                 `gorm:"column:avatar_path;type:varchar(256);not null;comment:用户头像路径" json:"avatar_path"`          // 用户头像路径
	Version       optimisticlock.Version `gorm:"column:version;type:int unsigned;default:1;comment:版本号，用于乐观锁" json:"version"`              // 版本号，用于乐观锁
	Gender        int                    `gorm:"column:gender;type:int;default:3;comment:用户性别 1 - 男  2 - 女  3- 保密" json:"gender"`          // 用户性别 1 - 男  2 - 女  3- 保密
	Birthday      time.Time              `gorm:"column:birthday;type:date;not null;default:2001-01-01" json:"birthday"`
	CntFollows    int64                  `gorm:"column:cnt_follows;type:int unsigned;not null;comment:关注数" json:"cnt_follows"` // 关注数
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}