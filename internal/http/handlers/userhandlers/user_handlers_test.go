package userhandlers

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cheatsnake/botyard/internal/entities/user"
	"github.com/cheatsnake/botyard/internal/http/helpers"
	"github.com/cheatsnake/botyard/internal/services/userservice"
	mock "github.com/cheatsnake/botyard/internal/storage/_mock"

	"github.com/gofiber/fiber/v2"
)

func TestCreate(t *testing.T) {
	userService := userservice.New(mock.UserStore())
	userHandlers := New(userService)
	testApp := fiber.New(fiber.Config{
		ErrorHandler: helpers.CursomErrorHandler,
	})
	testPath := "/test"
	testApp.Post(testPath, userHandlers.CreateUser)

	t.Run("create new user with valid body", func(t *testing.T) {
		jsonBody, err := json.Marshal(userservice.CreateBody{
			User: user.User{
				Nickname: "test",
			},
		})
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", jsonBody, err.Error(), nil)
		}

		req := httptest.NewRequest("POST", testPath, strings.NewReader(string(jsonBody)))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		resp, err := testApp.Test(req)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", req, err.Error(), nil)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", resp, err.Error(), nil)
		}

		expect := fiber.StatusCreated
		got := resp.StatusCode
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", string(body), got, expect)
		}

		if len(resp.Cookies()) == 0 {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", resp, resp.Cookies(), "user cookie")
		}
	})

	t.Run("create new user with invalid body", func(t *testing.T) {
		jsonBody, err := json.Marshal(helpers.JsonMessage("test"))
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", jsonBody, err.Error(), nil)
		}

		req := httptest.NewRequest("POST", testPath, strings.NewReader(string(jsonBody)))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		resp, err := testApp.Test(req)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", req, err.Error(), nil)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", resp, err.Error(), nil)
		}

		expect := fiber.StatusBadRequest
		got := resp.StatusCode
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", string(body), got, expect)
		}
	})
}
