package captcha

import "github.com/google/wire"

var Provider = wire.NewSet(NewEmail)
