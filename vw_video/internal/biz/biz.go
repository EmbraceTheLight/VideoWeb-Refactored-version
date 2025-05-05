package biz

import (
	"github.com/google/wire"
	"path/filepath"
	"runtime"
)

const (
	kb = 1024
	mb = 1024 * 1024

	coverName = "cover"
	videoPath = "video"
	dashPath  = "dash"
	dashName  = "dash.mpd"
)

const (
	uv_status  = "user_video_status"    // user-video status collection
	ub_status  = "user_barrage_status"  // user-barrage status collection
	uv_history = "user_video_history"   // user-video history collection
	uc_status  = "user_comment_upvoted" // user-comment status collection
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewVideoInfoUsecase,
	NewInteractUseCase,
	NewVideoCommentUsecase,
)

var (
	resourcePath string
)

func init() {
	initPath()
}
func initPath() {
	_, filename, _, _ := runtime.Caller(1)

	// tmp = internal/
	tmp := filepath.Dir(filepath.Dir(filename))

	// rootPath = video_web/
	rootPath := filepath.Dir(filepath.Dir(tmp))

	// resourcePath = video_web/resources
	resourcePath = filepath.Join(rootPath, "resources")
}
