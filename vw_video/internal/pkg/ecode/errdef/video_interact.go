package errdef

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_video/internal/pkg/ecode"
)

func init() {
	ErrUpvoteVideoFailed = kerr.New(ecode.VIDEOINTERACT_UpvoteVideoFailed, "视频点赞失败", "视频点赞失败, "+ecode.RetryStr)
	ErrUpvoteVideoRevertFailed = kerr.New(ecode.VIDEOINTERACT_UpvoteVideoRevertFailed, "视频点赞撤销失败(用于SAGA事务回滚报错)", "")

	ErrCancelUpvoteVideoFailed = kerr.New(ecode.VIDEOINTERACT_CancelUpvoteVideoFailed, "视频取消点赞失败", "视频取消点赞失败, "+ecode.RetryStr)
	ErrCancelUpvoteVideoRevertFailed = kerr.New(ecode.VIDEOINTERACT_CancelUpvoteVideoRevertFailed, "视频取消点赞撤销失败(用于SAGA事务回滚报错)", "")

	ErrFavoriteVideoFailed = kerr.New(ecode.VIDEOINTERACT_FavoriteVideoFailed, "视频收藏失败", "视频收藏失败, "+ecode.RetryStr)
	ErrFavoriteVideoRevertFailed = kerr.New(ecode.VIDEOINTERACT_FavoriteVideoRevertFailed, "视频收藏撤销失败(用于SAGA事务回滚报错)", "")

	ErrCancelFavoriteVideoFailed = kerr.New(ecode.VIDEOINTERACT_CancelFavoriteVideoFailed, "视频取消收藏失败", "视频取消收藏失败, "+ecode.RetryStr)
	ErrCancelFavoriteVideoRevertFailed = kerr.New(ecode.VIDEOINTERACT_CancelFavoriteVideoRevertFailed, "视频取消收藏撤销失败(用于SAGA事务回滚报错)", "")

	ErrThrowShellsFailed = kerr.New(ecode.VIDEOINTERACT_ThrowShellsFailed, "投币失败", "投币失败, "+ecode.RetryStr)
	ErrThrowShellsRevertFailed = kerr.New(ecode.VIDEOINTERACT_ThrowShellsRevertFailed, "投币撤销失败(用于SAGA事务回滚报错)", "")

	ErrShareVideoFailed = kerr.New(ecode.VIDEOINTERACT_ShareVideoFailed, "分享视频失败", "生成视频分享链接失败, "+ecode.RetryStr)
	ErrSendBarrageFailed = kerr.New(ecode.VIDEOINTERACT_SendBarrageFailed, "发送弹幕失败", "发送弹幕失败, "+ecode.RetryStr)
	ErrUpvoteBarrageFailed = kerr.New(ecode.VIDEOINTERACT_UpvoteBarrageFailed, "点赞弹幕失败", "点赞弹幕失败, "+ecode.RetryStr)
	ErrGetUserVideoStatusFailed = kerr.New(ecode.VIDEOINTERACT_GetUserVideoStatusFailed, "获取用户视频状态失败", "获取用户视频状态失败, "+ecode.RetryStr)
	ErrSetUserVideoStatusFailed = kerr.New(ecode.VIDEOINTERACT_SetUserVideoStatusFailed, "设置用户视频状态失败", "设置用户视频状态失败, "+ecode.RetryStr)
	ErrSetUserVideoStatusRevertFailed = kerr.New(ecode.VIDEOINTERACT_SetUserVideoStatusRevertFailed, "设置用户视频状态撤销失败(用于SAGA事务回滚报错)", "")
}

var (
	ErrUpvoteVideoFailed               *kerr.Error
	ErrUpvoteVideoRevertFailed         *kerr.Error
	ErrCancelUpvoteVideoFailed         *kerr.Error
	ErrCancelUpvoteVideoRevertFailed   *kerr.Error
	ErrFavoriteVideoFailed             *kerr.Error
	ErrFavoriteVideoRevertFailed       *kerr.Error
	ErrCancelFavoriteVideoFailed       *kerr.Error
	ErrCancelFavoriteVideoRevertFailed *kerr.Error
	ErrThrowShellsFailed               *kerr.Error
	ErrThrowShellsRevertFailed         *kerr.Error
	ErrShareVideoFailed                *kerr.Error
	ErrSendBarrageFailed               *kerr.Error
	ErrUpvoteBarrageFailed             *kerr.Error
	ErrGetUserVideoStatusFailed        *kerr.Error
	ErrSetUserVideoStatusFailed        *kerr.Error
	ErrSetUserVideoStatusRevertFailed  *kerr.Error
)
