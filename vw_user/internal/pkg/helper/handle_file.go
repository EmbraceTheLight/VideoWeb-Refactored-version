package helper

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
)

const defaultMultipartMemory = 32 << 20 // 32 MB
func CreateDir(dest string, fileMode fs.FileMode) (err error) {
	if err = os.MkdirAll(dest, fileMode); err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = os.RemoveAll(dest)
		}
	}()
	return nil
}

// FormFile imitate the behavior of github.com/gin-gonic/gin
func FormFile(req *http.Request, name string) (*multipart.FileHeader, error) {
	if req.MultipartForm == nil {
		if err := req.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := req.FormFile(name)
	if err != nil {
		return nil, err
	}
	_ = f.Close()
	return fh, nil
}

// ReadFileContent read the content of file
func ReadFileContent(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	ret, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// WriteToNewFile write the content of src to a new file at dst
func WriteToNewFile(src *multipart.FileHeader, dst string) error {
	srcFile, err := src.Open()
	defer srcFile.Close()
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	defer dstFile.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}
