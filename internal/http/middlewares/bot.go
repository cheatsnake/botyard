package middlewares

import (
	"botyard/internal/services"
	"botyard/pkg/extlib"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type BotMiddlewares struct {
	service *services.BotService
}

func NewBotMiddlewares(s *services.BotService) *BotMiddlewares {
	return &BotMiddlewares{
		service: s,
	}
}

func (bm *BotMiddlewares) Auth(c *fiber.Ctx) error {
	key := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	if key == "" {
		return extlib.ErrorForbidden("auth key is required")
	}

	if !strings.Contains(key, ":") {
		return extlib.ErrorForbidden("invalid auth key")
	}

	botId := strings.Split(key, ":")[0]
	botKey := strings.Split(key, ":")[1]

	err := bm.service.VerifyKey(botId, botKey)
	if err != nil {
		return err
	}

	c.Locals("botId", botId)
	c.Locals("botKey", botKey)

	return c.Next()
}
