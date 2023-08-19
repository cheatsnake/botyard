package middlewares

import (
	"botyard/pkg/extlib"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	userId := c.Cookies("userId", "")
	if userId == "" {
		return extlib.ErrorUnauthorized("user is unauthorized")
	}

	c.Locals("userId", userId)
	return c.Next()
}
