package helper

import (
	"errors"
	kerr "github.com/go-kratos/kratos/v2/errors"
)

func HandleError(kerror, err error) error {
	var retErr = &kerr.Error{}
	errors.As(kerror, &retErr)
	return retErr.WithMetadata(map[string]string{"err_msg": err.Error()})
}
