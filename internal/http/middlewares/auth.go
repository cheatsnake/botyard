package middlewares

import (
	"botyard/pkg/extlib"

	"github.com/gofiber/fiber/v2"
)

func (m *Middlewares) Auth(c *fiber.Ctx) error {
	userId := c.Cookies("userId", "")
	if userId == "" {
		return extlib.ErrorUnauthorized("user is unauthorized")
	}

	_, err := m.store.GetUser(userId)
	if err != nil {
		return extlib.ErrorNotFound(err.Error())
	}

	c.Locals("userId", userId)
	return c.Next()
}
