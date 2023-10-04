package middlewares

import (
	"github.com/cheatsnake/botyard/internal/services/userservice"
	"github.com/cheatsnake/botyard/pkg/exterr"

	"github.com/gofiber/fiber/v2"
)

func UserAuth(c *fiber.Ctx) error {
	token := c.Cookies("token", "")
	if token == "" {
		return exterr.ErrorUnauthorized("user is unauthorized")
	}

	userClaims, err := userservice.VerifyUserToken(token)
	if err != nil {
		return err
	}

	// Refresh token
	token, expires, err := userservice.GenerateUserToken(userClaims.Id, userClaims.Nickname)
	if err != nil {
		return err
	}

	cookie := &fiber.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expires,
	}

	c.Cookie(cookie)
	c.Locals("userId", userClaims.Id)

	return c.Next()
}
