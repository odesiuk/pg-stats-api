package rest

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ErrorHandler custom error handler for fiber.
func ErrorHandler(c *fiber.Ctx, err error) error {
	// gorm error,
	if errors.Is(err, gorm.ErrRecordNotFound) {
		e := fiber.ErrNotFound

		return c.Status(e.Code).
			JSON(ErrorResponse{Message: fmt.Sprint(e.Message)})
	}

	// errors with a hint
	var errHint *ErrWithHint
	if errors.As(err, &errHint) {
		if errHint.Err != nil {
			log.Println(errHint.Err)
		}

		return c.Status(errHint.Code).
			JSON(ErrorResponse{
				Message: errHint.Message,
				Field:   errHint.Field,
			})
	}

	// fiber errors.
	var errHTTP *fiber.Error
	if errors.As(err, &errHTTP) {
		return c.Status(errHTTP.Code).
			JSON(ErrorResponse{Message: fmt.Sprint(errHTTP.Message)})
	}

	// unknown errors.
	return c.Status(fiber.ErrInternalServerError.Code).
		JSON(ErrorResponse{Message: fiber.ErrInternalServerError.Message})
}
