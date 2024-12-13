package biz

import (
	"github.com/google/wire"
	"path/filepath"
	"runtime"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserIdentityUsecase, NewUserInfoUsecase)

var (
	// resourcePath is the root path of the service: /vw_user.
	rootPath          string
	resourcePath      string
	defaultAvatarPath string
)

func init() {
	initPath()
}
func initPath() {
	_, filename, _, _ := runtime.Caller(1)
	tmp := filepath.Dir(filename)
	rootPath = filepath.Dir(filepath.Dir(tmp))
	resourcePath = filepath.Join(rootPath, "resources")
	defaultAvatarPath = filepath.Join(resourcePath, "default", "avatar.png")
}
