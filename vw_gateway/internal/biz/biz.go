package biz

import (
	"github.com/google/wire"
	"vw_gateway/internal/biz/userbiz"
	"vw_gateway/internal/biz/videobiz"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	userbiz.NewUserIdentityUsecase,
	userbiz.NewCaptchaUsecase,
	userbiz.NewUserFileUsecase,
	userbiz.NewUserinfoUsecase,
	userbiz.NewFavoritesUsecase,
	userbiz.NewFollowUsecase,

	videobiz.NewVideoInfoUsecase,
	videobiz.NewInteractUsecase,
)
