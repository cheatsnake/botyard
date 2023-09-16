package bothandlers

import (
	"botyard/internal/http/helpers"
	"botyard/internal/services/botservice"
	"botyard/pkg/exterr"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	service *botservice.Service
}

func New(s *botservice.Service) *Handlers {
	return &Handlers{
		service: s,
	}
}

func (bh *Handlers) CreateBot(c *fiber.Ctx) error {
	body := new(botservice.CreateBody)
	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	result, err := bh.service.CreateBot(body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (bh *Handlers) GetBotById(c *fiber.Ctx) error {
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

func (bh *Handlers) GetCurrentBot(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return exterr.ErrorUnauthorized("bot is not authorized")
	}

	bot, err := bh.service.GetBotById(botId)
	if err != nil {
		return err
	}

	return c.JSON((bot))
}

func (bh *Handlers) GetAllBots(c *fiber.Ctx) error {
	bots, err := bh.service.GetAllBots()
	if err != nil {
		return err
	}

	return c.JSON(bots)
}

func (bh *Handlers) EditBot(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return exterr.ErrorUnauthorized("bot is not authorized")
	}

	body := new(botservice.EditBody)

	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	editedBot, err := bh.service.EditBot(botId, body)
	if err != nil {
		return err
	}

	return c.JSON(editedBot)
}

func (bh *Handlers) AddCommands(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return exterr.ErrorUnauthorized("bot is not authorized")
	}

	body := new(botservice.CommandsBody)
	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	err := bh.service.AddCommands(botId, body)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("commands added"))
}

func (bh *Handlers) RemoveCommand(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	if botId == "" {
		return exterr.ErrorUnauthorized("bot is not authorized")
	}

	alias := c.Query("alias", "")

	err := bh.service.RemoveCommand(botId, alias)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("command removed"))
}

func (bh *Handlers) GetKey(c *fiber.Ctx) error {
	botId := c.Params("id", "")

	result, err := bh.service.GetKey(botId)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (bh *Handlers) RefreshKey(c *fiber.Ctx) error {
	botId := c.Params("id", "")

	botKey, err := bh.service.GenerateKey(botId)
	if err != nil {
		return err
	}

	return c.JSON(botKey)
}
