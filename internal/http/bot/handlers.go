package bot

import (
	"botyard/internal/entities/bot"
	"botyard/internal/storage"

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

func (h *handlers) CreateBot(c *fiber.Ctx) error {
	body := new(struct {
		bot.Bot
		Id struct{} `json:"-"`
	})

	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	newBot := bot.New(body.Name)

	if body.Description != "" {
		newBot.SetDescription(body.Description)
	}

	if body.Avatar != "" {
		newBot.SetAvatar(body.Avatar)
	}

	if len(body.Commands) != 0 {
		for _, cmd := range body.Commands {
			newBot.AddCommand(cmd.Alias, cmd.Description)
		}
	}

	h.service.store.AddBot(newBot)

	return c.Status(fiber.StatusCreated).JSON(newBot)
}

func (h *handlers) EditBot(c *fiber.Ctx) error {
	newBot, fiberErr := h.findBot(c)
	if fiberErr != nil {
		return fiberErr
	}

	body := new(struct {
		bot.Bot
		Commands struct{} `json:"-"`
		Id       struct{} `json:"-"`
	})

	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if body.Name != "" {
		newBot.SetName(body.Name)
	}

	if body.Description != "" {
		newBot.SetDescription(body.Description)
	}

	if body.Avatar != "" {
		newBot.SetAvatar(body.Avatar)
	}

	err := h.service.store.EditBot(newBot)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(response{Message: "bot updated"})
}

func (h *handlers) AddBotCommands(c *fiber.Ctx) error {
	newBot, fiberErr := h.findBot(c)
	if fiberErr != nil {
		return fiberErr
	}

	body := new(struct{ Commands []bot.Command })
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	for _, cmd := range body.Commands {
		newBot.AddCommand(cmd.Alias, cmd.Description)
	}

	err := h.service.store.EditBot(newBot)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(response{Message: "new commands added"})
}

func (h *handlers) RemoveBotCommand(c *fiber.Ctx) error {
	newBot, fiberErr := h.findBot(c)
	if fiberErr != nil {
		return fiberErr
	}

	body := new(struct{ Alias string })
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	newBot.RemoveCommand(body.Alias)

	err := h.service.store.EditBot(newBot)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(response{Message: "command removed"})
}

func (h *handlers) GetBotCommands(c *fiber.Ctx) error {
	newBot, fiberErr := h.findBot(c)
	if fiberErr != nil {
		return fiberErr
	}

	return c.JSON(newBot.GetCommands())
}

func (h *handlers) findBot(c *fiber.Ctx) (*bot.Bot, *fiber.Error) {
	botId := c.Params("id", "")
	if botId == "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	foundBot, err := h.service.store.GetBot(botId)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return foundBot, nil
}
