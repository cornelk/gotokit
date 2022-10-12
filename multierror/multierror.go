// Package multierror provides appendable errors.
package multierror

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
)

// Error is an error type to track multiple errors. This is used to
// accumulate errors in cases and return them as a single "error".
type Error = multierror.Error

// New creates a new multierror with a custom formatter that only breaks
// errors over multiple lines if there is more than one error.
func New() *Error {
	return &Error{
		ErrorFormat: format,
	}
}

// Append is a helper function that will append more errors
// onto an Error in order to create a larger multi-error.
//
// If err is not a multierror.Error, then it will be turned into
// one. If any of the errs are multierr.Error, they will be flattened
// one level into err.
// Any nil errors within errs will be ignored. If err is nil, a new
// *Error will be returned.
func Append(err error, errs ...error) *Error {
	return multierror.Append(err, errs...)
}

func format(es []error) string {
	switch len(es) {
	case 0:
		return ""

	case 1:
		return es[0].Error()

	default:
		points := make([]string, len(es)+1)
		for i, err := range es {
			points[i] = fmt.Sprintf("\t* %s", err)
		}

		return strings.Join(points, "\n")
	}
}
