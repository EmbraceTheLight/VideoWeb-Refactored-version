package file

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"mime/multipart"
	"os"
	"path/filepath"
)

const (
	kb = 1024
	mb = 1024 * kb

	picMaxWidth = 10 * mb
)

var (
	errUploadFileTooLarge = kerr.New(0, "上传文件过大, 超过了10M", "上传文件不能超过10M")
	errUploadFileType     = kerr.New(0, "上传文件类型错误", "不支持该文件格式")
)

// picExtCheck 验证图片后缀名
var picExtCheck = make(map[string]struct{})

// videoExtCheck 验证视频后缀名
var videoExtCheck = make(map[string]struct{})

func init() {
	//支持的图片的格式
	picExtCheck[".jpg"] = struct{}{}
	picExtCheck[".jpeg"] = struct{}{}
	picExtCheck[".png"] = struct{}{}
	picExtCheck[".jfif"] = struct{}{}

	//支持的视频的格式
	videoExtCheck[".mp4"] = struct{}{}
	videoExtCheck[".mov"] = struct{}{}
	videoExtCheck[".avi"] = struct{}{}
	videoExtCheck[".mkv"] = struct{}{}
	videoExtCheck[".m4v"] = struct{}{}
	videoExtCheck[".3gp"] = struct{}{}
	videoExtCheck[".3g2"] = struct{}{}
}

// CheckIfPictureValid checks if the picture is valid
func CheckIfPictureValid(fh *multipart.FileHeader) error {
	ext := filepath.Ext(fh.Filename)
	if !checkPicExt(ext) {
		return errUploadFileType
	}
	if fh.Size > picMaxWidth {
		return errUploadFileTooLarge
	}
	return nil
}

// CheckIfVideoValid checks if the video is valid
func CheckIfVideoValid(fh *multipart.FileHeader) error {
	ext := filepath.Ext(fh.Filename)
	if !checkVideoExt(ext) {
		return errUploadFileType
	}
	return nil
}

func checkPicExt(ext string) bool {
	_, ok := picExtCheck[ext]
	return ok
}

func checkVideoExt(ext string) bool {
	_, ok := videoExtCheck[ext]
	return ok
}

// CheckIfFileExist checks if the file path exists
func CheckIfFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}
