package handlers

import (
	"botyard/internal/http/helpers"
	"botyard/internal/services"
	"botyard/pkg/extlib"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ChatHandlers struct {
	service *services.ChatService
}

func NewChatHandlers(s *services.ChatService) *ChatHandlers {
	return &ChatHandlers{
		service: s,
	}
}

func (h *ChatHandlers) Create(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%v", c.Locals("userId"))
	body := new(services.ChatCreateBody)

	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	chat, err := h.service.Create(userId, body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(chat)
}

func (h *ChatHandlers) Delete(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return extlib.ErrorBadRequest("id is required")
	}

	err := h.service.Delete(id)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("chat cleared"))
}
