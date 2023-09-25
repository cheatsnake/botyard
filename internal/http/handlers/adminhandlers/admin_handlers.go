package adminhandlers

import (
	"botyard/internal/config"
	"botyard/internal/http/helpers"
	"botyard/pkg/exterr"

	"github.com/gofiber/fiber/v2"
)

// Handler for reloading global config without stopping the application
func ReloadGlobalConfig(c *fiber.Ctx) error {
	err := config.Load()
	if err != nil {
		return exterr.ErrorInternalServer(err.Error())
	}

	return c.JSON(helpers.JsonMessage("config reloaded"))
}
