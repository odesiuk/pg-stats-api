package rest

import (
	"fmt"
	"net/http"
	"strings"
)

// ErrWithHint custom error for ErrorHandler.
type ErrWithHint struct {
	Code    int
	Message string
	Field   string
	Err     error
}

func (e ErrWithHint) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func (e *ErrWithHint) WithErr(err error) *ErrWithHint {
	e.Err = err

	return e
}

func ErrBadRequestInvalidParameter(name string, description ...string) *ErrWithHint {
	return &ErrWithHint{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("Invalid parameter '%s'. %s", name, strings.Join(description, ".")),
		Field:   name,
	}
}
