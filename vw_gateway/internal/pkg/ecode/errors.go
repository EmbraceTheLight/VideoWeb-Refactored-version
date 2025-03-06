package ecode

const (
	// identity errors
	IDENTITY_LoginFailed = 4001 + iota
	IDENTITY_UserNameEmpty
	IDENTITY_UserPasswordEmpty
	IDENTITY_UserNotFound
	IDENTITY_UserPasswordError
	IDENTITY_UserAlreadyExist
	IDENTITY_PasswordTooShort
	IDENTITY_PasswordNotMatch
	IDENTITY_SignatureTooLong
	IDENTITY_VerifyCodeExpired
	IDENTITY_VerifyCodeNotMatch
	IDENTITY_EncryptPasswordFailed
	IDENTITY_CreateUserDirFailed
	IDENTITY_ReadAvatarFailed
	IDENTITY_SaveAvatarFailed
	IDENTITY_ErrUserLoggedOut

	// auth errors
	AUTH_TokenMissing
	AUTH_TokenInvalid
	AUTH_AccessTokenExpired
	AUTH_RefreshTokenExpired
	AUTH_ParseTokenFailed
	AUTH_CreateTokenFailed

	// parse HTTP request's file errors
	HTTP_FormFileFailed
	HTTP_UploadAvatarFailed
)

var ()
