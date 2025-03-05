package biz_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	// resourcePath is the root path of the service: /vw_user.
	gwPath            string
	resourcePath      string
	defaultAvatarPath string
)

func init() {
	initPath()
}
func initPath() {
	_, filename, _, _ := runtime.Caller(1)
	tmp := filepath.Dir(filename)
	gwPath = filepath.Dir(filepath.Dir(tmp))
	userPath := filepath.Join(filepath.Dir(gwPath), "vw_user")
	resourcePath = filepath.Join(userPath, "resources")
	defaultAvatarPath = filepath.Join(resourcePath, "default", "avatar.png")
}
func TestBiz(t *testing.T) {
	fmt.Println(gwPath, resourcePath, defaultAvatarPath)
}
