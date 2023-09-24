package middlewares

import (
	"botyard/internal/http/helpers"
	"botyard/internal/services/userservice"
	"botyard/internal/tools/ulid"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestUserAuth(t *testing.T) {
	testApp := fiber.New(fiber.Config{
		ErrorHandler: helpers.CursomErrorHandler,
	})
	testPath := "/test"
	testApp.Get(testPath, UserAuth, func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	t.Run("with valid cookie", func(t *testing.T) {
		userId := ulid.New()
		token, _, err := userservice.GenerateUserToken(userId)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", token, err.Error(), nil)
		}

		req := httptest.NewRequest("GET", testPath, nil)
		req.AddCookie(&http.Cookie{Name: "token", Value: token})

		resp, err := testApp.Test(req)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", req, err.Error(), nil)
		}

		expect := fiber.StatusOK
		got := resp.StatusCode
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", resp, got, expect)
		}
	})

	t.Run("with invalid cookie", func(t *testing.T) {
		token := "jwt malformed"

		req := httptest.NewRequest("GET", testPath, nil)
		req.AddCookie(&http.Cookie{Name: "token", Value: token})

		resp, err := testApp.Test(req)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", req, err.Error(), nil)
		}

		expect := fiber.StatusUnauthorized
		got := resp.StatusCode
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", resp, got, expect)
		}
	})

	t.Run("without cookie", func(t *testing.T) {
		req := httptest.NewRequest("GET", testPath, nil)

		resp, err := testApp.Test(req)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", req, err.Error(), nil)
		}

		expect := fiber.StatusUnauthorized
		got := resp.StatusCode
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", resp, got, expect)
		}
	})
}
