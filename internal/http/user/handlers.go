package user

import (
	"botyard/internal/http/helpers"
	"time"

	"github.com/gofiber/fiber/v2"
)

type handlers struct {
	service *Service
}

func Handlers(s *Service) *handlers {
	return &handlers{
		service: s,
	}
}

func (h *handlers) Create(c *fiber.Ctx) error {
	body := new(createBody)

	if err := c.BodyParser(body); err != nil {
		return helpers.NewHttpError(fiber.StatusBadRequest, err.Error())
	}

	newUser, err := h.service.Create(body)
	if err != nil {
		return err
	}

	cookie := &fiber.Cookie{
		Name:    "userId",
		Value:   newUser.Id,
		Expires: time.Now().Add(24 * time.Hour),
	}

	c.Cookie(cookie)
	return c.Status(fiber.StatusCreated).JSON(newUser)
}
