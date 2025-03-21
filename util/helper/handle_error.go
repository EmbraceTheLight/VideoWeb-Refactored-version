package helper

import (
	"errors"
	kerr "github.com/go-kratos/kratos/v2/errors"
)

func HandleError(kerror *kerr.Error, additionalErrs ...error) *kerr.Error {
	var retErr = kerr.Clone(kerror)

	if additionalErrs == nil {
		return retErr
	}

	for _, additionalErr := range additionalErrs {
		var aErr = &kerr.Error{}
		ok := errors.As(additionalErr, &aErr)

		// if additionalErr is not a kratos error,this is a standard-library error, just take it as the reason of kratos error.
		if !ok {
			retErr.Reason = retErr.Reason + ": " + additionalErr.Error()
		} else {
			// The additionalErr is another kratos error, merge it with the original kratos error.
			// merge their metadata, append the additional error's message to the original error's message,
			// and append the additional error's reason to the original error's reason.
			for k, v := range aErr.Metadata {
				retErr.Metadata[k] = v
			}
			retErr.Message = retErr.Message + ": " + aErr.Message
			retErr.Reason = retErr.Reason + ": " + aErr.Reason
		}
	}
	return retErr
}
