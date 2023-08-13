package http

import (
	"botyard/internal/http/bot"
	"botyard/internal/http/chat"
	"botyard/internal/http/middlewares"
	"botyard/internal/http/user"
	"botyard/internal/storage"
	"errors"

	"github.com/gofiber/fiber/v2"
)

const bodySizeLimit = 25 * 1024 * 1024 // 25 MB

type Server struct {
	App   *fiber.App
	store storage.Storage
}

func New(store storage.Storage) *Server {
	return &Server{
		App: fiber.New(fiber.Config{
			ErrorHandler: errHandler,
			BodyLimit:    bodySizeLimit,
		}),
		store: store,
	}
}

func (s *Server) InitRoutes(prefix string) {
	api := s.App.Group(prefix)
	middlwr := middlewares.New(s.store)

	// Bot's handlers -----------------------------------
	bot := bot.Handlers(s.store)
	api.Get("/bot/:id", middlwr.Auth, bot.GetBotCommands)
	api.Post("/bot", middlwr.Admin, bot.CreateBot)
	api.Post("/bot/commands/:id", middlwr.Admin, bot.AddBotCommands)
	api.Put("/bot/:id", middlwr.Admin, bot.EditBot)
	api.Delete("/bot/command/:id", middlwr.Admin, bot.RemoveBotCommand)
	// --------------------------------------------------

	// User's handlers ---------------
	user := user.Handlers(s.store)
	api.Post("/user", user.CreateUser)
	// -------------------------------

	// Chat's handlers --------------------------------------
	chat := chat.Handlers(s.store)
	api.Get("/chat/:id", middlwr.Auth, chat.GetMessages)
	api.Post("/chat", middlwr.Auth, chat.CreateChat)
	api.Post("/chat/message", middlwr.Auth, chat.SendMessage)
	api.Delete("/chat/:id", middlwr.Auth, chat.ClearChat)
	// ------------------------------------------------------
}

func errHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(code).JSON(struct{ Message string }{Message: err.Error()})
}
