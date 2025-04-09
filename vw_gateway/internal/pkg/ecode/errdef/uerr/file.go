package uerr

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_gateway/internal/pkg/ecode"
)

func init() {
	ErrFormFileFiled = kerr.New(ecode.HTTP_FormFileFailed, "读取文件失败", "服务器内部错误,请稍后再试")
	ErrUploadAvatarFailed = kerr.New(ecode.HTTP_UploadAvatarFailed, "上传头像失败", "服务器内部错误,请稍后再试")
	ErrUploadFileFailed = kerr.New(ecode.FILE_UploadAvatarFailed, "上传文件失败", "服务器内部错误,请稍后再试")
	ErrUpdateAvatarFailed = kerr.New(ecode.FILE_UpdateAvatarFailed, "更新头像失败", "服务器内部错误,请稍后再试")

	ErrGetMpdFileFailed = kerr.New(ecode.FILE_GetMpdFileFailed, "获取mpd文件失败", "服务器内部错误,请稍后再试")
	ErrCoverFileNotFound = kerr.New(ecode.FILE_CoverFileNotFound, "封面文件不存在", "服务器内部错误,请稍后再试")
	ErrSegmentFileNotFound = kerr.New(ecode.FILE_SegmentFileNotFound, "分片文件不存在", "服务器内部错误,请稍后再试")
}

var (
	ErrFormFileFiled       *kerr.Error
	ErrUploadAvatarFailed  *kerr.Error
	ErrUploadFileFailed    *kerr.Error
	ErrUpdateAvatarFailed  *kerr.Error
	ErrGetMpdFileFailed    *kerr.Error
	ErrCoverFileNotFound   *kerr.Error
	ErrSegmentFileNotFound *kerr.Error
)
