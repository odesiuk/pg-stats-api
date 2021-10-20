package rest

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrBadRequestInvalidParameter(t *testing.T) {
	exp := &ErrWithHint{
		Code:    400,
		Message: "Invalid parameter 'param'. description",
		Field:   "param",
	}

	got := ErrBadRequestInvalidParameter("param", "description")

	assert.Equal(t, exp, got)

	_ = got.WithErr(errors.New("some error"))

	assert.Equal(t, exp.Error(), got.Error())
}
