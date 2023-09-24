package middlewares

import (
	"botyard/internal/services/userservice"
	"botyard/pkg/exterr"

	"github.com/gofiber/fiber/v2"
)

func UserAuth(c *fiber.Ctx) error {
	token := c.Cookies("token", "")
	if token == "" {
		return exterr.ErrorUnauthorized("user is unauthorized")
	}

	userId, err := userservice.VerifyUserToken(token)
	if err != nil {
		return err
	}

	// Refresh token
	token, expires, err := userservice.GenerateUserToken(userId)
	if err != nil {
		return err
	}

	cookie := &fiber.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expires,
	}

	c.Cookie(cookie)
	c.Locals("userId", userId)

	return c.Next()
}
