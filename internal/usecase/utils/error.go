package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type JSONError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func JSONErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		code = fiber.StatusNotFound
	}

	return ctx.Status(code).JSON(JSONError{
		Error:   true,
		Message: err.Error(),
	})
}
