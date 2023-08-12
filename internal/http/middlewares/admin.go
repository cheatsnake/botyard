package middlewares

import "github.com/gofiber/fiber/v2"

func (m *Middlewares) Admin(c *fiber.Ctx) error {
	return c.Next()
}
