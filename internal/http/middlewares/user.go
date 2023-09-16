package middlewares

import (
	"botyard/pkg/exterr"

	"github.com/gofiber/fiber/v2"
)

func UserAuth(c *fiber.Ctx) error {
	// TODO JWT auth
	userId := c.Cookies("userId", "")
	if userId == "" {
		return exterr.ErrorUnauthorized("user is unauthorized")
	}

	c.Locals("userId", userId)
	return c.Next()
}
