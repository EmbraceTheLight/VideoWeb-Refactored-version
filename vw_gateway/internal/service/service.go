package service

import (
	"github.com/google/wire"
	"vw_gateway/internal/service/user_service"
	video "vw_gateway/internal/service/video_service"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	user.NewUserIdentityService,
	user.NewCaptchaService,
	user.NewUserFileService,
	user.NewUserinfoService,
	user.NewFavoritesService,
	user.NewFollowService,

	video.NewVideoInfoService,
	video.NewInteractService,
)
