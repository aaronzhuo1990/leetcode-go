package string

import (
	"fmt"
)

/*************************************************************************/
// Error Types

// ErrType - error type
type ErrType string

// Bad request errors
const (
	// ErrTypeBadRequest - bad request
	ErrTypeBadRequest ErrType = "bad_request"
	// ErrTypeInternalServerErr - internal server error
	ErrTypeInternalServerErr ErrType = "internal_server_error"
	// ErrTypeUnknown - Unknown error
	ErrTypeUnknown ErrType = "unknown"
)

/************************************************************************/
// Error definition

// Error interface defines the errors used in this package
type Error interface {
	error
	Type() ErrType
}

// errorImpl - the implementation of Error interface
type errImpl struct {
	msg     string
	errType ErrType
}

// Error returns error message
func (e *errImpl) Error() string {
	if e != nil {
		return e.msg
	}
	return ""
}

// Type returns error type
func (e *errImpl) Type() ErrType {
	if e != nil {
		return e.errType
	}
	return ErrTypeUnknown
}

// newError - return an error with given error type
func newError(errType ErrType, format string, a ...interface{}) Error {
	return &errImpl{
		msg:     fmt.Sprintf(format, a...),
		errType: errType,
	}
}

// ConvertError - try converting an `error` interface to an `Error` interface
func ConvertError(err error) (Error, bool) {
	if e, ok := err.(Error); ok {
		return e, ok
	}

	return nil, false
}
