package codecs

import (
	"fmt"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"net/http"
	"os"
	videoinfov1 "vw_gateway/api/v1/video/videoinfo"
)

const (
	baseContentType = "application"
)

type FileResponse struct {
	FilePath string
	Filename string
	Header   http.Header
}

func newFileResponse(filePath, filename string, headers map[string]*videoinfov1.FileResp_HeaderValues) *FileResponse {
	fr := &FileResponse{
		FilePath: filePath,
		Filename: filename,
	}
	header := make(http.Header)
	for k, v := range headers {
		for _, value := range v.Value {
			header.Add(k, value)
		}
	}
	fr.Header = header
	return fr
}

func (fr *FileResponse) setHeader(key, value string) {
	fr.Header.Set(key, value)
}

func (fr *FileResponse) addHeader(key, value string) {
	fr.Header.Add(key, value)
}

func (fr *FileResponse) setContentType(contentType string) {
	fr.setHeader("Content-Type", contentType)
}

func (fr *FileResponse) setContentDisposition(dispositionType, filename string) {
	fr.setHeader("Content-Disposition", fmt.Sprintf("%s; filename*=UTF-8''%s", dispositionType, filename))
}

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return baseContentType + "/" + subtype
}

// ResponseEncoder encodes the object to the HTTP response.
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if protoFr, ok := v.(*videoinfov1.FileResp); ok {
		fr := newFileResponse(protoFr.FilePath, protoFr.Filename, protoFr.Headers)
		fmt.Println(fr.Filename)
		f, err := os.Open(fr.FilePath)
		if err != nil {
			return err
		}
		defer func() {
			f.Close()
			err := os.Remove(fr.FilePath)
			//TODO: if err != nil, use message queue to retry later
			if err != nil {

			}
		}()
		for k, headerValue := range fr.Header {
			for _, val := range headerValue {
				w.Header().Add(k, val)
			}
		}

		contentType := fr.Header.Get("Content-Type")
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		w.Header().Set("Content-Type", contentType)
		http.ServeFile(w, r, fr.FilePath)
		return nil
	}

	// Below is the DefaultResponseEncoder logic
	if v == nil {
		return nil
	}
	if rd, ok := v.(khttp.Redirector); ok {
		url, code := rd.Redirect()
		http.Redirect(w, r, url, code)
		return nil
	}
	codec, _ := khttp.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", ContentType(codec.Name()))
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
