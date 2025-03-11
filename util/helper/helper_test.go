package helper_test

import (
	"errors"
	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/spewerspew/spew"
	"testing"
	"util/helper"
)

func TestHandleError(t *testing.T) {
	kerror := kerr.New(1000, "test error reason", "test error message")

	// Test case 1: additional err is nil
	spew.Dump(helper.HandleError(kerror, nil))

	// Test case 2: additional err is standard library error
	spew.Dump(helper.HandleError(kerror, errors.New("I'm a standard library error")))

	// Test case 3: additional err is another kratos error
	additionalErr := kerr.New(2000, "I'm an additional error reason", "I'm an additional error message")
	spew.Dump(helper.HandleError(kerror, additionalErr))
}
