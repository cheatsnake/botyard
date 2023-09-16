package messagehandlers

import (
	"botyard/internal/http/helpers"
	"botyard/internal/services/messageservice"
	"botyard/pkg/exterr"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	service *messageservice.Service
}

func New(s *messageservice.Service) *Handlers {
	return &Handlers{
		service: s,
	}
}

func (mh *Handlers) SendUserMessage(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%s", c.Locals("userId"))
	if userId == "" {
		return exterr.ErrorUnauthorized("user is not authorized")
	}

	body := new(messageservice.CreateBody)

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

	body := new(messageservice.CreateBody)

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
