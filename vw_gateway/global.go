package vw_gateway

import (
	"path/filepath"
	"runtime"
)

var ResourcePath string

func init() {
	setResourcePath()
}

func setResourcePath() {
	_, filename, _, _ := runtime.Caller(1)
	ResourcePath = filepath.Dir(filepath.Dir(filename)) + "/resources"
}
