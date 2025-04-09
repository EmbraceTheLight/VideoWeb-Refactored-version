package verr

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_gateway/internal/pkg/ecode"
)

func init() {
	ErrGetVideoInfoFailed = kerr.New(ecode.VIDEO_GetVideoInfoFailed, "获取视频失败", "获取视频信息失败，请稍后再试")
	ErrGetVideoListFailed = kerr.New(ecode.VIDEO_GetVideoListFailed, "获取视频列表失败", "获取视频列表失败，请稍后再试")
	ErrGetVideoMpdFailed = kerr.New(ecode.VIDEOINFO_GetVideoMpdFailed, "获取视频mpd失败", "获取视频mpd失败，请稍后再试")
	ErrGetVideoCoverFailed = kerr.New(ecode.VIDEOINFO_GetVideoCoverFailed, "获取视频封面失败", "获取视频封面失败，请稍后再试")
	ErrGetVideoSegmentFailed = kerr.New(ecode.VIDEOINFO_GetVideoSegmentFailed, "获取视频服务器失败", "获取视频 segment 失败，请稍后再试")
}

var (
	ErrGetVideoListFailed    *kerr.Error
	ErrGetVideoInfoFailed    *kerr.Error
	ErrGetVideoMpdFailed     *kerr.Error
	ErrGetVideoCoverFailed   *kerr.Error
	ErrGetVideoSegmentFailed *kerr.Error
)
