package rest

import (
	"fmt"
	"net/http"
)

// ErrWithHint custom error for ErrorHandler.
type ErrWithHint struct {
	Code    int
	Message string
	Field   string
	Err     error
}

func (e ErrWithHint) Error() string {
	return fmt.Sprintf("%d: %s | Err: %v", e.Code, e.Message, e.Err)
}

func (e *ErrWithHint) WithErr(err error) *ErrWithHint {
	e.Err = err

	return e
}

func ErrBadRequestInvalidParameter(name string) *ErrWithHint {
	return &ErrWithHint{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("Invalid parameter '%s'", name),
		Field:   name,
	}
}
