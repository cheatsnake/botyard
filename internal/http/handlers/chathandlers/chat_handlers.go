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

func (h *Handlers) CreateChat(c *fiber.Ctx) error {
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

func (h *Handlers) GetChatsByBot(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%v", c.Locals("userId"))
	botId := c.Params("botId", "")

	chats, err := h.service.GetChats(userId, botId)
	if err != nil {
		return err
	}

	return c.JSON(chats)
}

func (h *Handlers) GetChatsByUser(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%v", c.Locals("botId"))
	userId := c.Params("userId", "")

	chats, err := h.service.GetChats(userId, botId)
	if err != nil {
		return err
	}

	return c.JSON(chats)
}

func (h *Handlers) DeleteChat(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return exterr.ErrorBadRequest("id is required")
	}

	err := h.service.Delete(id)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("chat deleted"))
}

func (mh *Handlers) SendUserMessage(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%s", c.Locals("userId"))
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

func (mh *Handlers) GetMessagesByChat(c *fiber.Ctx) error {
	id := c.Params("id", "")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 20)

	result, err := mh.service.GetMessagesByChat(id, page, limit)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
