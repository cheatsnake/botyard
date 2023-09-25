package chathandlers

import (
	"botyard/internal/http/client"
	"botyard/internal/http/helpers"
	"botyard/internal/services/botservice"
	"botyard/internal/services/chatservice"
	"botyard/pkg/exterr"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	service    *chatservice.Service
	botService *botservice.Service
	client     *client.Client
}

func New(s *chatservice.Service, b *botservice.Service) *Handlers {
	return &Handlers{
		service:    s,
		botService: b,
		client:     client.New(),
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

	chat, err := h.service.CreateChat(userId, body.BotId)
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
	userId := fmt.Sprintf("%s", c.Locals("userId"))
	id := c.Params("id", "")

	_, err := h.service.CheckChatAccess(id, "", userId)
	if err != nil {
		return err
	}

	err = h.service.DeleteChat(id)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("chat deleted"))
}

func (h *Handlers) SendUserMessage(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%s", c.Locals("userId"))
	body := new(chatservice.CreateBody)

	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	chat, err := h.service.CheckChatAccess(body.ChatId, "", userId)
	if err != nil {
		return err
	}

	body.SenderId = userId
	msg, err := h.service.AddMessage(body)
	if err != nil {
		return err
	}

	botWh, err := h.botService.GetWebhook(chat.BotId)
	if err == nil {
		h.client.SendPost(botWh.Url, msg, botWh.Secret)
	}

	return c.Status(fiber.StatusCreated).JSON(msg)
}

func (h *Handlers) SendBotMessage(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	body := new(chatservice.CreateBody)

	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	_, err := h.service.CheckChatAccess(body.ChatId, botId, "")
	if err != nil {
		return err
	}

	body.SenderId = botId
	msg, err := h.service.AddMessage(body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(msg)
}

func (h *Handlers) GetMessagesByChat(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	userId := fmt.Sprintf("%s", c.Locals("userId"))
	id := c.Params("id", "")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 20)

	_, err := h.service.CheckChatAccess(id, botId, userId)
	if err != nil {
		return err
	}

	result, err := h.service.GetMessagesByChat(id, page, limit)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
