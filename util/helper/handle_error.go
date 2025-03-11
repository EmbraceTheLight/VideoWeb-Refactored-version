package helper

import (
	"errors"
	kerr "github.com/go-kratos/kratos/v2/errors"
)

func HandleError(kerror *kerr.Error, additionalErr error) *kerr.Error {
	var retErr = kerr.Clone(kerror)

	if additionalErr == nil {
		return retErr
	}

	var aErr = &kerr.Error{}
	ok := errors.As(additionalErr, &aErr)
	// if additionalErr is not a kratos error,this is a standard-library error, just take it as the cause of kratos error.
	if !ok {
		retErr.Message = retErr.Message + ": " + additionalErr.Error()
		return retErr.WithCause(additionalErr)

		// The additionalErr is another kratos error, merge it with the original kratos error.
		// merge their metadata, and append the additional error's message to the original error's message.
	} else {
		for k, v := range aErr.Metadata {
			retErr.Metadata[k] = v
		}
		retErr.Message = retErr.Message + ": " + aErr.Message
		return retErr.WithCause(additionalErr)
	}
	//return retErr.WithCause(additionalErr).WithMetadata(map[string]string{"err_msg": additionalErr.Error()})
}
