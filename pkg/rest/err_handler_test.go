package rest

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

func TestErrorHandler(t *testing.T) {
	tests := map[string]struct {
		err     error
		exp     ErrorResponse
		expCode int
		expErr  error
	}{
		"ErrRecordNotFound": {
			err:     gorm.ErrRecordNotFound,
			exp:     ErrorResponse{Message: "Not Found"},
			expCode: 404,
			expErr:  nil,
		},
		"ErrWithHint": {
			err: &ErrWithHint{
				Code:    400,
				Message: "FieldError",
				Field:   "some_field",
				Err:     errors.New("err_message"),
			},
			exp:     ErrorResponse{Message: "FieldError", Field: "some_field"},
			expCode: 400,
			expErr:  nil,
		},
		"AnyFiberError": {
			err:     fiber.ErrUnauthorized,
			exp:     ErrorResponse{Message: "Unauthorized"},
			expCode: 401,
			expErr:  nil,
		},
		"InternalServerError": {
			err:     errors.New("some_error"),
			exp:     ErrorResponse{Message: "Internal Server Error"},
			expCode: 500,
			expErr:  nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var ctx fasthttp.RequestCtx
			ctx.Init(new(fasthttp.Request), nil, nil)

			err := ErrorHandler(fiber.New().AcquireCtx(&ctx), tt.err)
			if err != nil {
				assert.IsType(t, tt.expErr, err)
				assert.Equal(t, tt.expErr.Error(), err.Error())

				return
			}

			var result ErrorResponse
			err = json.Unmarshal(ctx.Response.Body(), &result)
			if err != nil {
				t.Fatal(err)
			}

			assert.EqualValues(t, tt.exp, result)
		})
	}
}
