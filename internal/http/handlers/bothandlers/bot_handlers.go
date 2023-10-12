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
		return exterr.ErrorBadRequest("ID is required.")
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

func (bh *Handlers) DeleteBot(c *fiber.Ctx) error {
	id := c.Params("id", "")

	err := bh.service.DeleteBot(id)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("Bot deleted."))
}

func (bh *Handlers) AddCommands(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	body := new([]botservice.PreparedCommand)
	if err := c.BodyParser(body); err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	err := bh.service.AddCommands(botId, *body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(helpers.JsonMessage("Commands added."))
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

	return c.JSON(cmds)
}

func (bh *Handlers) RemoveCommand(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
	id := c.Params("id", "")

	if id == "" {
		return exterr.ErrorBadRequest("ID is required.")
	}

	cmd, err := bh.service.GetCommand(id)
	if err != nil {
		return err
	}

	if cmd.BotId != botId {
		return exterr.ErrorForbidden("Can't delete a command that belongs to another bot.")
	}

	err = bh.service.RemoveCommand(id)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("Command removed."))
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

	return c.JSON(fiber.Map{"key": botKey})
}

func (bh *Handlers) SaveWebhook(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))
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

func (bh *Handlers) GetWebhook(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))

	wh, err := bh.service.GetWebhook(botId)
	if err != nil {
		return err
	}

	return c.JSON(wh)
}

func (bh *Handlers) DeleteWebhook(c *fiber.Ctx) error {
	botId := fmt.Sprintf("%s", c.Locals("botId"))

	err := bh.service.DeleteWebhook(botId)
	if err != nil {
		return err
	}

	return c.JSON(helpers.JsonMessage("Webhook deleted."))
}
