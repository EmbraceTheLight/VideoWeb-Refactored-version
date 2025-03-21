package errdef

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_gateway/internal/pkg/ecode"
)

func init() {
	// Register errors
	{
		ErrUserLoggedOut = kerr.New(ecode.IDENTITY_ErrUserLoggedOut, "用户已退出登录", "用户已退出登录")
	}
}

var (
	// Login errors
	ErrUserLoggedOut *kerr.Error
)
