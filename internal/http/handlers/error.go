package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type response struct {
	Message string `json:"message"`
}

func ErrHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(code).JSON(response{Message: err.Error()})
}

func newErr(err error, code int) *fiber.Error {
	return &fiber.Error{
		Code:    code,
		Message: err.Error(),
	}
}
