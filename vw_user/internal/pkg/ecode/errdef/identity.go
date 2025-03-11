package errdef

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_user/internal/pkg/ecode"
)

func init() {
	// Login errors
	{
		ErrUserNameEmpty = kerr.New(ecode.IDENTITY_UserNameEmpty, "用户名为空", "用户名不能为空")
		ErrUserPasswordEmpty = kerr.New(ecode.IDENTITY_UserPasswordEmpty, "密码为空", "密码不能为空")
		ErrUserNotFound = kerr.New(ecode.IDENTITY_UserNotFound, "用户不存在", "用户不存在")
		ErrUserPasswordError = kerr.New(ecode.IDENTITY_UserPasswordError, "用户密码错误", "密码错误")
	}

	// Register errors
	{
		ErrUserAlreadyExist = kerr.New(ecode.IDENTITY_UserAlreadyExist, "用户名已存在", "该用户名已存在")
		ErrPasswordTooShort = kerr.New(ecode.IDENTITY_PasswordTooShort, "密码太短，少于6位", "密码长度不能少于6位")
		ErrPasswordNotMatch = kerr.New(ecode.IDENTITY_PasswordNotMatch, "第一次输入的密码和第二次输入的密码不匹配", "两次输入的密码不匹配，请重新输入")
		ErrSignatureTooLong = kerr.New(ecode.IDENTITY_SignatureTooLong, "签名过长，超过了25个字符", "签名不能超过25个字符")
		ErrVerifyCodeExpired = kerr.New(ecode.IDENTITY_VerifyCodeExpired, "验证码已过期", "验证码已过期,请重新获取")
		ErrVerifyCodeNotMatch = kerr.New(ecode.IDENTITY_VerifyCodeNotMatch, "验证码错误", "验证码错误,请重新输入验证码")
		ErrEncryptPasswordFailed = kerr.New(ecode.IDENTITY_EncryptPasswordFailed, "用户密码加密失败", "服务器内部错误,请稍后重试")
		ErrCreateUserDirFailed = kerr.New(ecode.IDENTITY_CreateUserDirFailed, "创建用户目录失败", "服务器内部错误,请稍后重试")
		ErrCreateUserAvatarFailed = kerr.New(ecode.IDENTITY_CreateUserAvatarFailed, "创建用户头像文件失败", "服务器内部错误,请稍后重试")
		ErrReadAvatarFailed = kerr.New(ecode.IDENTITY_ReadAvatarFailed, "读取用户头像失败", "服务器内部错误,请稍后重试")
		ErrSaveAvatarFailed = kerr.New(ecode.IDENTITY_SaveAvatarFailed, "保存用户头像失败", "服务器内部错误,请稍后重试")
		ErrCreateUserRecordsFailed = kerr.New(ecode.IDENTITY_ErrCreateUserRecordsFailed, "创建用户记录失败", "服务器内部错误,请稍后重试")
	}

	// Logout errors
	{
		ErrLogoutFailed = kerr.New(ecode.IDENTITY_ErrLogoutFailed, "退出登录失败", "退出登录失败")
	}
}

var (
	// Login errors
	ErrUserNameEmpty     *kerr.Error
	ErrUserPasswordEmpty *kerr.Error
	ErrUserNotFound      *kerr.Error
	ErrUserPasswordError *kerr.Error

	// Register errors
	ErrUserAlreadyExist        *kerr.Error
	ErrPasswordTooShort        *kerr.Error
	ErrPasswordNotMatch        *kerr.Error
	ErrSignatureTooLong        *kerr.Error
	ErrVerifyCodeExpired       *kerr.Error
	ErrVerifyCodeNotMatch      *kerr.Error
	ErrEncryptPasswordFailed   *kerr.Error
	ErrCreateUserDirFailed     *kerr.Error
	ErrCreateUserAvatarFailed  *kerr.Error
	ErrReadAvatarFailed        *kerr.Error
	ErrSaveAvatarFailed        *kerr.Error
	ErrCreateUserRecordsFailed *kerr.Error

	// Logout errors
	ErrLogoutFailed *kerr.Error
)
