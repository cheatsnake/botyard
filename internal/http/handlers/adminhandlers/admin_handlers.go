package adminhandlers

import (
	"github.com/cheatsnake/botyard/internal/config"
	"github.com/cheatsnake/botyard/internal/http/helpers"
	"github.com/cheatsnake/botyard/pkg/exterr"

	"github.com/gofiber/fiber/v2"
)

// Handler for reloading global config without stopping the application
func ReloadGlobalConfig(c *fiber.Ctx) error {
	err := config.Load()
	if err != nil {
		return exterr.ErrorInternalServer(err.Error())
	}

	return c.JSON(helpers.JsonMessage("Config reloaded."))
}
