package ecode

const (
	// Intertal errors
	INTERNAL_ErrInternal = 5001 + iota

	// identity errors
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

	// file errors
	FILE_UploadAvatarFailed
	FILE_UpdateAvatarFailed
)

var ()
