package userbiz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewUserIdentityUsecase,
	NewCaptchaUsecase,
	NewUserFileUsecase,
	NewUserinfoUsecase,
	NewFavoritesUsecase,
	NewFollowUsecase,
)
