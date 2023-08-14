package bot

import (
	"botyard/internal/http/helpers"
	"botyard/pkg/extlib"

	"github.com/gofiber/fiber/v2"
)

type handlers struct {
	service *Service
}

func Handlers(s *Service) *handlers {
	return &handlers{
		service: s,
	}
}

func (h *handlers) Create(c *fiber.Ctx) error {
	body := new(createBody)
	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	newBot, err := h.service.Create(body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(newBot)
}

func (h *handlers) Edit(c *fiber.Ctx) error {
	botId := c.Params("id", "")
	if botId == "" {
		return extlib.ErrorBadRequest("id is required")
	}

	body := new(editBody)

	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	editedBot, err := h.service.Edit(botId, body)
	if err != nil {
		return err
	}

	return c.JSON(editedBot)
}

func (h *handlers) AddCommands(c *fiber.Ctx) error {
	botId := c.Params("id", "")
	if botId == "" {
		return extlib.ErrorBadRequest("id is required")
	}

	body := new(commandsBody)
	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	err := h.service.AddCommands(botId, body)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("commands added"))
}

func (h *handlers) RemoveCommand(c *fiber.Ctx) error {
	botId := c.Params("id", "")
	if botId == "" {
		return extlib.ErrorBadRequest("id is required")
	}

	body := new(commandBody)
	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	err := h.service.RemoveCommand(botId, body)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("command removed"))
}

func (h *handlers) GetCommands(c *fiber.Ctx) error {
	botId := c.Params("id", "")
	if botId == "" {
		return extlib.ErrorBadRequest("id is required")
	}

	commands, err := h.service.GetCommands(botId)
	if err != nil {
		return err
	}

	return c.JSON(commands)
}
