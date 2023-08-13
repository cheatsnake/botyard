package chat

import (
	"botyard/internal/http/helpers"
	"fmt"

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
	userId := fmt.Sprintf("%v", c.Locals("userId"))
	b := new(createBody)

	if err := c.BodyParser(b); err != nil {
		return helpers.NewHttpError(fiber.StatusBadRequest, err.Error())
	}

	chat, err := h.service.Create(userId, b.BotId)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(chat)
}

func (h *handlers) Delete(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return helpers.NewHttpError(fiber.StatusBadRequest, "id is required")
	}

	err := h.service.Delete(id)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("chat cleared"))
}
