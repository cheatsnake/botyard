package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func (m *Middlewares) Auth(c *fiber.Ctx) error {
	userId := c.Cookies("userId", "")
	if userId == "" {
		return fiber.NewError(fiber.StatusUnauthorized, ("user is unauthorized"))
	}

	_, err := m.store.GetUser(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	c.Locals("userId", userId)
	return c.Next()
}
