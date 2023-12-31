package helpers

import (
	"errors"

	"github.com/cheatsnake/botyard/pkg/exterr"

	"github.com/gofiber/fiber/v2"
)

func CursomErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *exterr.ExtendedError
	if errors.As(err, &e) {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(code).JSON(JsonMessage(err.Error()))
}
