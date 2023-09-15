package middlewares

import (
	"botyard/internal/services/botservice"
	"botyard/pkg/exterr"
	"strings"

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
	key := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	if key == "" {
		return exterr.ErrorForbidden("auth key is required")
	}

	if !strings.Contains(key, ":") {
		return exterr.ErrorForbidden("invalid auth key")
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
