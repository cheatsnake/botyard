package chat

import (
	"botyard/internal/entities/chat"
	"botyard/internal/entities/message"
	"botyard/internal/storage"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type handlers struct {
	service *service
}

type response struct {
	Message string `json:"message"`
}

func Handlers(s storage.Storage) *handlers {
	return &handlers{
		service: &service{
			store: s,
		},
	}
}

func (h *handlers) CreateChat(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%v", c.Locals("userId"))
	b := new(struct {
		BotId string `json:"botId"`
	})

	if err := c.BodyParser(b); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err := h.service.store.GetBot(b.BotId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	chat := chat.New(userId, b.BotId)
	err = h.service.store.AddChat(chat)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(chat)
}

func (h *handlers) SendMessage(c *fiber.Ctx) error {
	b := new(struct {
		message.Message
		Id        struct{} `json:"-"`
		Timestamp struct{} `json:"-"`
	})

	if err := c.BodyParser(b); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// chat, err := ch.store.GetChat(b.ChatId)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusNotFound, err.Error())
	// }

	// err = chat.SendMessage(b.SenderId, b.Body, b.FileIds)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, err.Error())
	// }

	return c.JSON(response{Message: "message sended"})
}

func (h *handlers) GetMessages(c *fiber.Ctx) error {
	// chatId := c.Params("id", "")
	// page := c.QueryInt("page", 1)
	// limit := c.QueryInt("limit", 20)

	// chat, err := ch.store.GetChat(chatId)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusNotFound, err.Error())
	// }

	// result, err := chat.GetMessages(page, limit)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, err.Error())
	// }

	// return c.JSON(result)
	return nil
}

func (h *handlers) ClearChat(c *fiber.Ctx) error {
	// chatId := c.Params("id", "")

	// chat, err := h.service.store.GetChat(chatId)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusNotFound, err.Error())
	// }

	// // err = chat.Clear()
	// // if err != nil {
	// // 	return fiber.NewError(fiber.StatusBadRequest, err.Error())
	// // }

	return c.JSON(response{Message: "chat cleared"})
}
