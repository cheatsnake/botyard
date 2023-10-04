package middlewares

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cheatsnake/botyard/internal/http/helpers"

	"github.com/gofiber/fiber/v2"
)

func TestAdminAuth(t *testing.T) {
	testApp := fiber.New(fiber.Config{
		ErrorHandler: helpers.CursomErrorHandler,
	})
	testPath := "/test"
	testApp.Get(testPath, AdminAuth, func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	t.Run("valid admin key", func(t *testing.T) {
		req := httptest.NewRequest("GET", testPath, nil)
		req.Header.Set("Authorization", os.Getenv("ADMIN_SECRET_KEY"))

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

	t.Run("invalid admin key", func(t *testing.T) {
		req := httptest.NewRequest("GET", testPath, nil)
		req.Header.Set("Authorization", "test")

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
}
