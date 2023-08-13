package middlewares

import (
	"botyard/internal/http/helpers"

	"github.com/gofiber/fiber/v2"
)

func (m *Middlewares) Auth(c *fiber.Ctx) error {
	userId := c.Cookies("userId", "")
	if userId == "" {
		return helpers.NewHttpError(fiber.StatusUnauthorized, ("user is unauthorized"))
	}

	_, err := m.store.GetUser(userId)
	if err != nil {
		return helpers.NewHttpError(fiber.StatusNotFound, err.Error())
	}

	c.Locals("userId", userId)
	return c.Next()
}
