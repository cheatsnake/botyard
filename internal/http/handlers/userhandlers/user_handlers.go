package userhandlers

import (
	"botyard/internal/services/userservice"
	"botyard/pkg/exterr"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	service *userservice.Service
}

func New(s *userservice.Service) *Handlers {
	return &Handlers{
		service: s,
	}
}

func (h *Handlers) Create(c *fiber.Ctx) error {
	body := new(userservice.CreateBody)

	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
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
