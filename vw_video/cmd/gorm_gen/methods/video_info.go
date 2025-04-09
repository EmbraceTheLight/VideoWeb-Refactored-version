package methods

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type Video struct {
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                                                                   // 创建时间
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`                                                                   // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);index:video_index_deletedat,priority:1;comment:删除时间" json:"deleted_at"`                            // 删除时间
	Title        string         `gorm:"column:title;type:varchar(100);not null;comment: 视频标题" json:"title"`                                                                  //  视频标题
	Description  string         `gorm:"column:description;type:text;comment:视频描述" json:"description"`                                                                        // 视频描述
	Class        string         `gorm:"column:class;type:varchar(20);not null;uniqueIndex:video_joint_idx_class_hot,priority:1;comment:视频所属类别（以英文逗号 , 分隔多个类别）" json:"class"` // 视频所属类别（以英文逗号 , 分隔多个类别）
	Hot          int64          `gorm:"column:hot;type:bigint;uniqueIndex:video_joint_idx_class_hot,priority:2;comment:视频热度" json:"hot"`                                     // 视频热度
	Tags         string         `gorm:"column:tags;type:varchar(200);comment:视频标签，以英文逗号分隔" json:"tags"`                                                                      // 视频标签，以英文逗号分隔
	VideoPath    string         `gorm:"column:video_path;type:varchar(200);not null;comment:视频文件路径" json:"video_path"`                                                       // 视频文件路径
	VideoID      int64          `gorm:"column:video_id;type:bigint;primaryKey;comment:视频ID" json:"video_id"`                                                                 // 视频ID
	UserID       int64          `gorm:"column:user_id;type:bigint;not null;comment:上传者用户ID" json:"user_id"`                                                                  // 上传者用户ID
	UserName     string         `gorm:"column:user_name;type:varchar(200);not null;comment:上传者用户名" json:"user_name"`                                                         // 上传者用户名
	Likes        int64          `gorm:"column:likes;type:int unsigned;comment:视频点赞数" json:"likes"`                                                                           // 视频点赞数
	Shells       int64          `gorm:"column:shells;type:int unsigned;comment:视频获得的贝壳数" json:"shells"`                                                                      // 视频获得的贝壳数
	CntBarrages  int64          `gorm:"column:cnt_barrages;type:int unsigned;comment:视频弹幕数" json:"cnt_barrages"`                                                             // 视频弹幕数
	CntShares    int64          `gorm:"column:cnt_shares;type:int unsigned;comment:视频分享数" json:"cnt_shares"`                                                                 // 视频分享数
	CntFavorited int64          `gorm:"column:cnt_favorited;type:int unsigned;comment:视频收藏数" json:"cnt_favorited"`                                                           // 视频收藏数
	CntViewed    int64          `gorm:"column:cnt_viewed;type:bigint;comment:视频观看数（点开就算看）" json:"cnt_viewed"`                                                                // 视频观看数（点开就算看）
	Duration     string         `gorm:"column:duration;type:varchar(10);not null;comment:视频时长" json:"duration"`                                                              // 视频时长
	Size         int64          `gorm:"column:size;type:bigint;not null;comment:视频文件大小" json:"size"`                                                                         // 视频文件大小
	CoverPath    string         `gorm:"column:cover_path;type:varchar(200);not null;comment:视频封面路径" json:"cover_path"`                                                       // 视频封面路径
}

const Separator = ","

func (v *Video) CovertTags() []string {
	return strings.Split(v.Tags, Separator)
}

func (v *Video) CovertClass() []string {
	return strings.Split(v.Class, Separator)
}
