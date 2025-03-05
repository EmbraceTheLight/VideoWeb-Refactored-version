package middlewares

import (
	"github.com/google/wire"
	"vw_user/internal/pkg/middlewares/auth"
)

// ProviderSet is a Wire provider set that includes all middlewares.
var ProviderSet = wire.NewSet(auth.NewJWTAuth)
