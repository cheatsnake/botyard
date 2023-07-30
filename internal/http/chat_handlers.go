package http

import (
	"botyard/internal/chat"

	"github.com/gofiber/fiber/v2"
)

type chatBody struct {
	chat.Chat
	Id struct{} `json:"-"`
}

type messageBody struct {
	chat.Message
	Id        struct{} `json:"-"`
	Timestamp struct{} `json:"-"`
}

func (s *Server) createChat(c *fiber.Ctx) error {
	b := new(chatBody)

	if err := c.BodyParser(b); err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	chat := chat.New(b.MemberIds, s.Storage)
	err := s.Storage.AddChat(chat)
	if err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusCreated).JSON(chat)
}

func (s *Server) sendMessage(c *fiber.Ctx) error {
	b := new(messageBody)

	if err := c.BodyParser(b); err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	chat, err := s.Storage.GetChat(b.ChatId)
	if err != nil {
		return newErr(err, fiber.StatusNotFound)
	}

	err = chat.SendMessage(b.SenderId, b.Body, b.FileIds)
	if err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	return c.JSON(response{Message: "message sended"})
}

func (s *Server) getMessages(c *fiber.Ctx) error {
	chatId := c.Params("id", "")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 20)

	chat, err := s.Storage.GetChat(chatId)
	if err != nil {
		return newErr(err, fiber.StatusNotFound)
	}

	result, err := chat.GetMessages(page, limit)
	if err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	return c.JSON(result)
}

// func (s *Server) loadFile(c *fiber.Ctx) error {}

func (s *Server) clearChat(c *fiber.Ctx) error {
	chatId := c.Params("id", "")

	chat, err := s.Storage.GetChat(chatId)
	if err != nil {
		return newErr(err, fiber.StatusNotFound)
	}

	err = chat.Clear()
	if err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	return c.JSON(response{Message: "chat cleared"})
}
