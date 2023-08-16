package handlers

import (
	"botyard/internal/services"
	"botyard/pkg/extlib"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	service *services.UserService
}

func NewUserHandlers(s *services.UserService) *UserHandlers {
	return &UserHandlers{
		service: s,
	}
}

func (h *UserHandlers) Create(c *fiber.Ctx) error {
	body := new(services.UserCreateBody)

	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
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
