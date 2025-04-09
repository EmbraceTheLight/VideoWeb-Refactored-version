package ecode

const (
	// Intertal errors
	INTERNAL_ErrInternal = 3001 + iota

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
	FILE_GetMpdFileFailed
	FILE_CoverFileNotFound
	FILE_SegmentFileNotFound

	// video errors
	VIDEO_GetVideoListFailed
	VIDEO_GetVideoInfoFailed
	VIDEOINFO_GetVideoFileFailed
	VIDEOINFO_GetVideoMpdFailed
	VIDEOINFO_GetVideoCoverFailed
	VIDEOINFO_GetVideoSegmentFailed
)

var ()
