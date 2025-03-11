package errdef

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_gateway/internal/pkg/ecode"
)

func init() {
	ErrFormFileFiled = kerr.New(ecode.HTTP_FormFileFailed, "读取文件失败", "服务器内部错误,请稍后再试")
	ErrUploadAvatarFailed = kerr.New(ecode.HTTP_UploadAvatarFailed, "上传头像失败", "服务器内部错误,请稍后再试")
}

var (
	ErrFormFileFiled      *kerr.Error
	ErrUploadAvatarFailed *kerr.Error
)
