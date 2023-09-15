package middlewares

import (
	"botyard/pkg/exterr"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	key := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	if key != os.Getenv("ADMIN_SECRET_KEY") {
		return exterr.ErrorForbidden("access denied")
	}

	return c.Next()
}
