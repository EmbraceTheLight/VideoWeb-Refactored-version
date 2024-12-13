package codecs

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdhttp "net/http"
)

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		return
	}
	code := int(se.Code)
	if 0 < code && code <= 600 {
		w.WriteHeader(code)
	} else {
		w.WriteHeader(stdhttp.StatusOK)
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	_, _ = w.Write(body)
}
