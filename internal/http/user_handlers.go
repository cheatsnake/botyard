package http

import (
	"botyard/internal/user"

	"github.com/gofiber/fiber/v2"
)

type userBody struct {
	user.User
	Id struct{} `json:"-"`
}

func (s *Server) createUser(c *fiber.Ctx) error {
	b := new(userBody)

	if err := c.BodyParser(b); err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	user := user.New(b.Name)
	return c.Status(fiber.StatusCreated).JSON(user)
}
