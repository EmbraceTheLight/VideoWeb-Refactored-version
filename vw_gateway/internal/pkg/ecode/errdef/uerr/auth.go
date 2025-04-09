package uerr

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_gateway/internal/pkg/ecode"
)

func init() {
	ErrTokenMissing = kerr.New(ecode.AUTH_TokenMissing, "token缺失", "token缺失")
	ErrTokenInvalid = kerr.New(ecode.AUTH_TokenInvalid, "token无效", "token无效")
	ErrAccessTokenExpired = kerr.New(ecode.AUTH_AccessTokenExpired, "access token已过期,由客户端携带已更新过的access token重新请求", "")
	ErrRefreshTokenExpired = kerr.New(ecode.AUTH_RefreshTokenExpired, "refresh token已过期", "refresh token已过期,请重新登录")
	ErrParseTokenFailed = kerr.New(ecode.AUTH_ParseTokenFailed, "解析token失败", "服务器内部错误,请稍后再试")
	ErrCreateTokenFailed = kerr.New(ecode.AUTH_CreateTokenFailed, "创建token失败", "服务器内部错误,请稍后再试")
}

var (
	ErrTokenMissing        *kerr.Error
	ErrTokenInvalid        *kerr.Error
	ErrAccessTokenExpired  *kerr.Error
	ErrRefreshTokenExpired *kerr.Error
	ErrParseTokenFailed    *kerr.Error
	ErrCreateTokenFailed   *kerr.Error
)
