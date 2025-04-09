package errdef

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_video/internal/pkg/ecode"
)

func init() {
	ErrGetVideoInfoFailed = kerr.New(ecode.VIDEOINFO_GetVideoInfoFailed, "获取视频失败", "获取视频信息失败，请稍后再试")
	ErrVideoNotFound = kerr.New(ecode.VIDEOINFO_VideoNotFound, "视频不存在", "获取视频信息失败，请稍后再试")
	ErrGetVideoListFailed = kerr.New(ecode.VIDEOINFO_GetVideoListFailed, "获取视频列表失败", "获取视频列表失败，请稍后再试")

	ErrUploadVideoInfoFailed = kerr.New(ecode.VIDEOINFO_UploadVideoInfoFailed, "上传视频信息失败", "上传视频信息失败，请稍后再试")
	ErrUploadVideoFileFailed = kerr.New(ecode.VIDEOINFO_UploadVideoFileFailed, "上传视频文件失败", "上传视频文件失败，请稍后再试")
	ErrUploadVideoCoverFailed = kerr.New(ecode.VIDEOINFO_UploadVideoCoverFailed, "上传视频封面失败", "上传视频封面失败，请稍后再试")
	ErrGetVideoFileFailed = kerr.New(ecode.VIDEOINFO_GetVideoFileFailed, "获取视频文件失败", "获取视频文件失败，请稍后再试")
	ErrGetVideoMpdFailed = kerr.New(ecode.VIDEOINFO_GetVideoMpdFailed, "获取视频mpd失败", "获取视频mpd失败，请稍后再试")
	ErrGetVideoCoverFailed = kerr.New(ecode.VIDEOINFO_GetVideoCoverFailed, "获取视频封面失败", "获取视频封面失败，请稍后再试")
	ErrGetVideoSegmentFailed = kerr.New(ecode.VIDEOINFO_GetVideoSegmentFailed, "获取视频服务器失败", "获取视频 segment 失败，请稍后再试")
}

var (
	ErrGetVideoListFailed     *kerr.Error
	ErrGetVideoInfoFailed     *kerr.Error
	ErrVideoNotFound          *kerr.Error
	ErrUploadVideoInfoFailed  *kerr.Error
	ErrUploadVideoFileFailed  *kerr.Error
	ErrUploadVideoCoverFailed *kerr.Error
	ErrGetVideoFileFailed     *kerr.Error
	ErrGetVideoMpdFailed      *kerr.Error
	ErrGetVideoCoverFailed    *kerr.Error
	ErrGetVideoSegmentFailed  *kerr.Error
)
