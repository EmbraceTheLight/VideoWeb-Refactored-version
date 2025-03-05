package helper

import (
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
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
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return fh, nil
}

// ReadFileContentFromFile read the content of local file
func ReadFileContentFromFile(filename string) ([]byte, error) {
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

// ReadFileContentFromFileHeader read the content of file header
func ReadFileContentFromFileHeader(file *multipart.FileHeader, fileContent *[]byte) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	_, err = f.Read(*fileContent)
	if err != nil {
		return err

	}
	return nil
}

// WriteToNewFile write the content of src to a new file at dst
func WriteToNewFile(dst string, src string) error {
	srcFile, err := os.Open(src)
	defer srcFile.Close()

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
