package handlers

import (
	"botyard/internal/http/helpers"
	"botyard/internal/services"
	"botyard/pkg/extlib"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MessageHandlers struct {
	service *services.MessageService
}

func NewMessageHandlers(s *services.MessageService) *MessageHandlers {
	return &MessageHandlers{
		service: s,
	}
}

func (mh *MessageHandlers) SendUserMessage(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%s", c.Locals("userId"))
	if userId == "" {
		return extlib.ErrorUnauthorized("user is not authorized")
	}

	body := new(services.CreateMessageBody)

	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	body.SenderId = userId
	err := mh.service.AddMessage(body)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("message sended"))
}

func (mh *MessageHandlers) SendBotMessage(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return extlib.ErrorUnauthorized("bot is not authorized")
	}

	body := new(services.CreateMessageBody)

	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	body.SenderId = botId
	err := mh.service.AddMessage(body)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("message sended"))
}

func (mh *MessageHandlers) GetByChat(c *fiber.Ctx) error {
	chatId := c.Params("chatId", "")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 20)

	result, err := mh.service.GetMessagesByChat(chatId, page, limit)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
