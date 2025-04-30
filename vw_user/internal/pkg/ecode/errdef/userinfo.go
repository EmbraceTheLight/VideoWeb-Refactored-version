package errdef

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_user/internal/pkg/ecode"
)

func init() {
	ErrForgetPasswordFailed = kerr.New(ecode.USERINFO_ForgetPasswordFailed, "用户重置密码失败", "重置密码失败，请稍后再试")
	ErrModifyPasswordFailed = kerr.New(ecode.USERINFO_ModifyPasswordFailed, "用户修改密码失败", "修改密码失败，请稍后再试")
	ErrModifyUsernameFailed = kerr.New(ecode.USERINFO_ModifyUsernameFailed, "用户修改用户名失败", "修改用户名失败，请稍后再试")
	ErrAddUserLikeFailed = kerr.New(ecode.USERINFO_AddUserLikeFailed, "添加用户点赞数失败", "服务器内部错误，请稍后再试")

	// 获取用户信息相关错误
	{
		ErrGetUserInfoFailed = kerr.New(ecode.USERINFO_GetUserInfoFailed, "获取用户信息失败", "获取用户信息失败，请稍后再试")
	}

	// 修改邮箱相关错误
	{
		ErrModifyEmailFailed = kerr.New(ecode.USERINFO_ModifyEmailFailed, "用户修改邮箱失败", "修改邮箱失败，请稍后再试")

	}
	ErrModifySignatureFailed = kerr.New(ecode.USERINFO_ModifySignatureFailed, "用户修改签名失败", "修改签名失败，请稍后再试")
}

var (
	ErrForgetPasswordFailed  *kerr.Error
	ErrModifyPasswordFailed  *kerr.Error
	ErrModifyUsernameFailed  *kerr.Error
	ErrGetUserInfoFailed     *kerr.Error
	ErrModifyEmailFailed     *kerr.Error
	ErrModifySignatureFailed *kerr.Error
	ErrAddUserLikeFailed     *kerr.Error
)
