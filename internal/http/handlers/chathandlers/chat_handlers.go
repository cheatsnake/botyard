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

func (mh *Handlers) SendUserMessage(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%s", c.Locals("userId"))
	if userId == "" {
		return exterr.ErrorUnauthorized("user is not authorized")
	}

	body := new(chatservice.CreateBody)

	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	body.SenderId = userId
	err := mh.service.AddMessage(body)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("message sended"))
}

func (mh *Handlers) SendBotMessage(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return exterr.ErrorUnauthorized("bot is not authorized")
	}

	body := new(chatservice.CreateBody)

	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	body.SenderId = botId
	err := mh.service.AddMessage(body)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("message sended"))
}

func (mh *Handlers) GetByChat(c *fiber.Ctx) error {
	chatId := c.Params("chatId", "")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 20)

	result, err := mh.service.GetMessagesByChat(chatId, page, limit)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
