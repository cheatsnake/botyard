package middlewares

import (
	"strings"

	"github.com/cheatsnake/botyard/internal/services/botservice"
	"github.com/cheatsnake/botyard/pkg/exterr"

	"github.com/gofiber/fiber/v2"
)

type BotMiddlewares struct {
	service *botservice.Service
}

func NewBotMiddlewares(s *botservice.Service) *BotMiddlewares {
	return &BotMiddlewares{
		service: s,
	}
}

func (bm *BotMiddlewares) Auth(c *fiber.Ctx) error {
	key := c.Get(fiber.HeaderAuthorization, "")

	if key == "" {
		return exterr.ErrorForbidden("Bot key is required.")
	}

	if !strings.Contains(key, ":") {
		return exterr.ErrorForbidden("Invalid bot key.")
	}

	botId := strings.Split(key, ":")[0]
	token := strings.Split(key, ":")[1]

	err := bm.service.VerifyKeyData(botId, token)
	if err != nil {
		return err
	}

	c.Locals("botId", botId)

	return c.Next()
}
