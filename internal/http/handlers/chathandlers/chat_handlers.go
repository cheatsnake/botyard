package chathandlers

import (
	"fmt"

	"github.com/cheatsnake/botyard/internal/http/client"
	"github.com/cheatsnake/botyard/internal/http/helpers"
	"github.com/cheatsnake/botyard/internal/services/botservice"
	"github.com/cheatsnake/botyard/internal/services/chatservice"
	"github.com/cheatsnake/botyard/pkg/exterr"

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
	botId := c.Query("bot_id", "")

	if botId == "" {
		return exterr.ErrorBadRequest("bot id is required")
	}

	chats, err := h.service.GetChats(userId, botId)
	if err != nil {
		return err
	}

	return c.JSON(chats)
}

func (h *Handlers) GetChatsByUser(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%v", c.Locals("botId"))
	userId := c.Query("user_id", "")

	if userId == "" {
		return exterr.ErrorBadRequest("user id is required")
	}

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
	var botId, userId string

	if c.Locals("botId") != nil {
		botId = fmt.Sprintf("%s", c.Locals("botId"))
	}

	if c.Locals("userId") != nil {
		userId = fmt.Sprintf("%s", c.Locals("userId"))
	}

	id := c.Params("id", "")
	_, err := h.service.CheckChatAccess(id, botId, userId)
	if err != nil {
		return err
	}

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 20)
	senderId := c.Query("sender_id", "")
	since := c.QueryInt("since", 0)

	result, err := h.service.GetMessagesByChat(id, senderId, page, limit, int64(since))
	if err != nil {
		return err
	}

	return c.JSON(result)
}
