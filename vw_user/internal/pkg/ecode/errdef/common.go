package errdef

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_user/internal/pkg/ecode"
)

func init() {
	INTERNAL_ERROR = kerr.New(ecode.INTERNAL_ErrInternal, "internal error", "服务内部错误，请稍后再试")
}

var (
	INTERNAL_ERROR *kerr.Error
)
