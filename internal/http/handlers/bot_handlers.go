package handlers

import (
	"botyard/internal/http/helpers"
	"botyard/internal/services"
	"botyard/pkg/extlib"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type BotHandlers struct {
	service *services.BotService
}

func NewBotHandlers(s *services.BotService) *BotHandlers {
	return &BotHandlers{
		service: s,
	}
}

func (h *BotHandlers) Create(c *fiber.Ctx) error {
	body := new(services.BotCreateBody)
	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	result, err := h.service.Create(body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *BotHandlers) Edit(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return extlib.ErrorBadRequest("id is required")
	}

	body := new(services.BotEditBody)

	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	editedBot, err := h.service.Edit(botId, body)
	if err != nil {
		return err
	}

	return c.JSON(editedBot)
}

func (h *BotHandlers) AddCommands(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return extlib.ErrorBadRequest("id is required")
	}

	body := new(services.BotCommandsBody)
	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	err := h.service.AddCommands(botId, body)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("commands added"))
}

func (h *BotHandlers) RemoveCommand(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return extlib.ErrorBadRequest("id is required")
	}

	alias := c.Query("alias", "")

	err := h.service.RemoveCommand(botId, alias)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("command removed"))
}

func (h *BotHandlers) GetInfo(c *fiber.Ctx) error {
	botId := c.Params("id", "")
	if botId == "" {
		return extlib.ErrorBadRequest("id is required")
	}

	bot, err := h.service.FindById(botId)
	if err != nil {
		return err
	}

	return c.JSON(bot)
}

func (h *BotHandlers) GetCommands(c *fiber.Ctx) error {
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

func (h *BotHandlers) RefreshBotKey(c *fiber.Ctx) error {
	body := new(struct {
		BotId string `json:"botId"`
	})

	if err := c.BodyParser(body); err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	botKeyRes, err := h.service.GenerateBotKey(body.BotId)
	if err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	return c.JSON(botKeyRes)
}
