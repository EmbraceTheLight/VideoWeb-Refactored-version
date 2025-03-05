package biz

import (
	"github.com/google/wire"
	"path/filepath"
	"runtime"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewUserIdentityUsecase,
	NewUserInfoUsecase,
	NewCaptchaUsecase,
	NewFileUsecase,
)

var (
	// resourcePath is the root path of the service: /vw_user.
	rootPath          string
	resourcePath      string
	defaultAvatarPath string
	defaultAvatarName string = "avatar.jpg"
)

func init() {
	initPath()
}
func initPath() {
	_, filename, _, _ := runtime.Caller(1)

	// tmp = /biz
	tmp := filepath.Dir(filename)

	// rootPath = /vw_user
	rootPath = filepath.Dir(filepath.Dir(tmp))

	// resourcePath = /vw_user/resources
	resourcePath = filepath.Join(rootPath, "resources")

	// defaultAvatarPath = /vw_user/resources/default/avatar.png
	defaultAvatarPath = filepath.Join(resourcePath, "default", defaultAvatarName)
}
