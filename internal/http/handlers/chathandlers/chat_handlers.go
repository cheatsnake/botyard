package chathandlers

import (
	"botyard/internal/http/helpers"
	"botyard/internal/services/chatservice"
	"botyard/pkg/exterr"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	service *chatservice.Service
}

func New(s *chatservice.Service) *Handlers {
	return &Handlers{
		service: s,
	}
}

func (h *Handlers) Create(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%v", c.Locals("userId"))
	body := new(struct {
		BotId string `json:"botId"`
	})

	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	chat, err := h.service.Create(userId, body.BotId)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(chat)
}

func (h *Handlers) GetMany(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%v", c.Locals("userId"))
	botId := c.Params("botId", "")

	chats, err := h.service.GetByBot(userId, botId)
	if err != nil {
		return err
	}

	return c.JSON(chats)
}

func (h *Handlers) Delete(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return exterr.ErrorBadRequest("id is required")
	}

	err := h.service.Delete(id)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("chat cleared"))
}
