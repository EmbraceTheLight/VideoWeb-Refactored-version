package errdef

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_user/internal/pkg/ecode"
)

func init() {
	ErrUploadAvatarFailed = kerr.New(ecode.FILE_UploadAvatarFailed, "上传文件失败", "服务器内部错误,请稍后再试")
	ErrUpdateAvatarFailed = kerr.New(ecode.FILE_UpdateAvatarFailed, "更新文件失败", "服务器内部错误,请稍后再试")
}

var (
	ErrUploadAvatarFailed *kerr.Error
	ErrUpdateAvatarFailed *kerr.Error
)
