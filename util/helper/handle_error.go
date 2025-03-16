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

	// if additionalErr is not a kratos error,this is a standard-library error, just take it as the reason of kratos error.
	if !ok {
		retErr.Reason = retErr.Reason + ": " + additionalErr.Error()
		return retErr.WithMetadata(map[string]string{"err_reason": additionalErr.Error()})

		// The additionalErr is another kratos error, merge it with the original kratos error.
		// merge their metadata, append the additional error's message to the original error's message,
		// and append the additional error's reason to the original error's reason.

	} else {
		for k, v := range aErr.Metadata {
			retErr.Metadata[k] = v
		}
		retErr.Message = retErr.Message + ": " + aErr.Message

		retErr.Reason = retErr.Reason + ": " + aErr.Reason
		return retErr.WithMetadata(map[string]string{"err_reason": additionalErr.Error()})
	}
	//return retErr.WithCause(additionalErr).WithMetadata(map[string]string{"err_msg": additionalErr.Error()})
}
