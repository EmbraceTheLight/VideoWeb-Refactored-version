package errdef

import (
	kerror "github.com/go-kratos/kratos/v2/errors"
	"vw_video/internal/pkg/ecode"
)

var (
	ErrInternal = kerror.New(ecode.InternalError, "internal error", "服务器内部错误，请稍后再试")
)
