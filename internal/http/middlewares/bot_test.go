package middlewares

import (
	"botyard/internal/http/helpers"
	"botyard/internal/services/botservice"
	mock "botyard/internal/storage/_mock"
	"botyard/internal/tools/ulid"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestBotAuth(t *testing.T) {
	botService := botservice.New(mock.BotStore())
	botMiddlewares := NewBotMiddlewares(botService)
	testApp := fiber.New(fiber.Config{
		ErrorHandler: helpers.CursomErrorHandler,
	})
	testPath := "/test"
	testApp.Get(testPath, botMiddlewares.Auth, func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	t.Run("bot auth with valid key", func(t *testing.T) {
		req := httptest.NewRequest("GET", testPath, nil)
		req.Header.Set("Authorization", "Bearer "+ulid.New()+":"+"test")

		resp, err := testApp.Test(req)
		if err != nil {
			t.Errorf("got: %v,\nexpected: %v\n", err.Error(), nil)
		}

		expect := fiber.StatusOK
		got := resp.StatusCode
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", resp, got, expect)
		}
	})

	t.Run("bot auth with empty key", func(t *testing.T) {
		req := httptest.NewRequest("GET", testPath, nil)
		req.Header.Set("Authorization", "")

		resp, err := testApp.Test(req)
		if err != nil {
			t.Errorf("got: %v,\nexpected: %v\n", err.Error(), nil)
		}

		expect := fiber.StatusForbidden
		got := resp.StatusCode
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", resp, got, expect)
		}
	})

	t.Run("bot auth with invalid key", func(t *testing.T) {
		for _, key := range []string{"invalidkey", "invalid:key"} {
			req := httptest.NewRequest("GET", testPath, nil)
			req.Header.Set("Authorization", key)

			resp, err := testApp.Test(req)
			if err != nil {
				t.Errorf("got: %v,\nexpected: %v\n", err.Error(), nil)
			}

			expect := fiber.StatusForbidden
			got := resp.StatusCode
			if got != expect {
				t.Errorf("%#v\ngot: %v,\nexpected: %v\n", resp, got, expect)
			}
		}
	})
}
