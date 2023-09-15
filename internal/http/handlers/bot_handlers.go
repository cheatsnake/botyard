package handlers

import (
	"botyard/internal/http/helpers"
	"botyard/internal/services"
	"botyard/pkg/exterr"
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

func (bh *BotHandlers) CreateBot(c *fiber.Ctx) error {
	body := new(services.BotCreateBody)
	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	result, err := bh.service.CreateBot(body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (bh *BotHandlers) GetBot(c *fiber.Ctx) error {
	botId := c.Params("id", "")
	if botId == "" {
		return exterr.ErrorBadRequest("id is required")
	}

	bot, err := bh.service.GetBotById(botId)
	if err != nil {
		return err
	}

	return c.JSON(bot)
}

func (bh *BotHandlers) GetAllBots(c *fiber.Ctx) error {
	bots, err := bh.service.GetAllBots()
	if err != nil {
		return err
	}

	return c.JSON(bots)
}

func (bh *BotHandlers) EditBot(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return exterr.ErrorBadRequest("id is required")
	}

	body := new(services.BotEditBody)

	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	editedBot, err := bh.service.EditBot(botId, body)
	if err != nil {
		return err
	}

	return c.JSON(editedBot)
}

func (bh *BotHandlers) AddCommands(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return exterr.ErrorBadRequest("id is required")
	}

	body := new(services.BotCommandsBody)
	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	err := bh.service.AddCommands(botId, body)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("commands added"))
}

func (bh *BotHandlers) RemoveCommand(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return exterr.ErrorBadRequest("id is required")
	}

	alias := c.Query("alias", "")

	err := bh.service.RemoveCommand(botId, alias)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("command removed"))
}

func (bh *BotHandlers) GetKey(c *fiber.Ctx) error {
	botId := c.Params("id", "")

	result, err := bh.service.GetKey(botId)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (bh *BotHandlers) RefreshKey(c *fiber.Ctx) error {
	botId := c.Params("id", "")

	botKey, err := bh.service.GenerateKey(botId)
	if err != nil {
		return err
	}

	return c.JSON(botKey)
}
