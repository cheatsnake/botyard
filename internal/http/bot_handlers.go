package http

import (
	"botyard/internal/bot"

	"github.com/gofiber/fiber/v2"
)

type botBody struct {
	bot.Bot
	Id struct{} `json:"-"`
}

func (s *Server) createBot(c *fiber.Ctx) error {
	b := new(botBody)

	if err := c.BodyParser(b); err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	bot := bot.New(b.Name)
	if b.Description != "" {
		bot.SetDescription(b.Description)
	}
	if b.Avatar != "" {
		bot.SetAvatar(b.Avatar)
	}
	if len(b.Commands) != 0 {
		for _, cmd := range b.Commands {
			bot.AddCommand(cmd.Alias, cmd.Description)
		}
	}

	s.Storage.AddBot(bot)

	return c.Status(fiber.StatusCreated).JSON(bot)
}

func (s *Server) getBotCommands(c *fiber.Ctx) error {
	botId := c.Params("id")

	bot, err := s.Storage.GetBot(botId)
	if err != nil {
		return newErr(err, fiber.StatusNotFound)
	}

	return c.JSON(bot.GetCommands())
}
