package bothandlers

import (
	"fmt"

	"github.com/cheatsnake/botyard/internal/http/helpers"
	"github.com/cheatsnake/botyard/internal/services/botservice"
	"github.com/cheatsnake/botyard/pkg/exterr"

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
	body := new(botservice.CommandsBody)
	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	err := bh.service.AddCommands(botId, body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(helpers.JsonMessage("commands added"))
}

func (bh *Handlers) GetCommands(c *fiber.Ctx) error {
	var botId string

	if c.Locals("botId") != nil {
		botId = fmt.Sprintf("%s", c.Locals("botId"))
	}

	if botId == "" {
		botId = c.Params("id", "")
	}

	cmds, err := bh.service.GetCommands(botId)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"commands": cmds,
	})
}

func (bh *Handlers) RemoveCommand(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	alias := c.Query("alias", "")

	err := bh.service.RemoveCommand(botId, alias)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("command removed"))
}

func (bh *Handlers) GetKey(c *fiber.Ctx) error {
	botId := c.Params("id", "")

	key, err := bh.service.GetKey(botId)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"key": key,
	})
}

func (bh *Handlers) RefreshKey(c *fiber.Ctx) error {
	botId := c.Params("id", "")

	botKey, err := bh.service.GenerateKey(botId)
	if err != nil {
		return err
	}

	return c.JSON(botKey)
}

func (bh *Handlers) CreateWebhook(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	body := new(botservice.WebhookBody)
	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	result, err := bh.service.SaveWebhook(botId, body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (bh *Handlers) GetWebhook(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))

	wh, err := bh.service.GetWebhook(botId)
	if err != nil {
		return err
	}

	return c.JSON(wh)
}

func (bh *Handlers) EditWebhook(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))

	_, err := bh.service.GetWebhook(botId)
	if err != nil {
		return err
	}

	body := new(botservice.WebhookBody)
	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	result, err := bh.service.SaveWebhook(botId, body)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (bh *Handlers) DeleteWebhook(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))

	err := bh.service.DeleteWebhook(botId)
	if err != nil {
		return err
	}

	return c.JSON("webhook deleted")
}
