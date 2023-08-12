package handlers

import (
	"botyard/internal/bot"
	"botyard/internal/storage"

	"github.com/gofiber/fiber/v2"
)

type Bot struct {
	store storage.Storage
}

func NewBot(store storage.Storage) *Bot {
	return &Bot{
		store: store,
	}
}

type createBotBody struct {
	bot.Bot
	Id struct{} `json:"-"`
}

type editBotBody struct {
	bot.Bot
	Commands struct{} `json:"-"`
	Id       struct{} `json:"-"`
}

func (b *Bot) CreateBot(c *fiber.Ctx) error {
	body := new(createBotBody)

	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	botCopy := bot.New(body.Name)

	if body.Description != "" {
		botCopy.SetDescription(body.Description)
	}

	if body.Avatar != "" {
		botCopy.SetAvatar(body.Avatar)
	}

	if len(body.Commands) != 0 {
		for _, cmd := range body.Commands {
			botCopy.AddCommand(cmd.Alias, cmd.Description)
		}
	}

	b.store.AddBot(botCopy)

	return c.Status(fiber.StatusCreated).JSON(botCopy)
}

func (b *Bot) EditBot(c *fiber.Ctx) error {
	botCopy, fiberErr := b.findBot(c)
	if fiberErr != nil {
		return fiberErr
	}

	body := new(editBotBody)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if body.Name != "" {
		botCopy.SetName(body.Name)
	}

	if body.Description != "" {
		botCopy.SetDescription(body.Description)
	}

	if body.Avatar != "" {
		botCopy.SetAvatar(body.Avatar)
	}

	err := b.store.EditBot(botCopy)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(response{Message: "bot updated"})
}

func (b *Bot) AddBotCommands(c *fiber.Ctx) error {
	botCopy, fiberErr := b.findBot(c)
	if fiberErr != nil {
		return fiberErr
	}

	body := new(struct{ Commands []bot.Command })
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	for _, cmd := range body.Commands {
		botCopy.AddCommand(cmd.Alias, cmd.Description)
	}

	err := b.store.EditBot(botCopy)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(response{Message: "new commands added"})
}

func (b *Bot) RemoveBotCommand(c *fiber.Ctx) error {
	botCopy, fiberErr := b.findBot(c)
	if fiberErr != nil {
		return fiberErr
	}

	body := new(struct{ Alias string })
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	botCopy.RemoveCommand(body.Alias)

	err := b.store.EditBot(botCopy)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(response{Message: "command removed"})
}

func (b *Bot) GetBotCommands(c *fiber.Ctx) error {
	botCopy, fiberErr := b.findBot(c)
	if fiberErr != nil {
		return fiberErr
	}

	return c.JSON(botCopy.GetCommands())
}

func (b *Bot) findBot(c *fiber.Ctx) (*bot.Bot, *fiber.Error) {
	botId := c.Params("id", "")
	if botId == "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	botCopy, err := b.store.GetBot(botId)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return botCopy, nil
}
