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
	NewFavoritesUsecase,
	NewFollowUseCase,
)

const avatar = "avatar"

var (
	// resourcePath is the root path of the service: /vw_user.
	rootPath          string
	resourcePath      string
	defaultAvatarPath string
	defaultAvatarName string = avatar + ".jpg"
)

func init() {
	initPath()
}
func initPath() {
	_, filename, _, _ := runtime.Caller(1)

	// tmp = internal/
	tmp := filepath.Dir(filepath.Dir(filename))

	// rootPath = video_web/
	rootPath = filepath.Dir(filepath.Dir(tmp))

	// resourcePath = video_web/resources
	resourcePath = filepath.Join(rootPath, "resources")

	// defaultAvatarPath = video_web/resources/default/avatar.png
	defaultAvatarPath = filepath.Join(resourcePath, "default", defaultAvatarName)
}
