package middlewares

import (
	"os"

	"github.com/cheatsnake/botyard/pkg/exterr"

	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	key := c.Get(fiber.HeaderAuthorization)

	if key != os.Getenv("ADMIN_SECRET_KEY") {
		return exterr.ErrorForbidden("Access denied.")
	}

	return c.Next()
}
