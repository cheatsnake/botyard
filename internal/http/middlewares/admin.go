package middlewares

import "github.com/gofiber/fiber/v2"

func Admin(c *fiber.Ctx) error {
	return c.Next()
}
