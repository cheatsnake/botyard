package http

import (
	"botyard/internal/http/bot"
	"botyard/internal/http/chat"
	"botyard/internal/http/file"
	"botyard/internal/http/helpers"
	"botyard/internal/http/message"
	"botyard/internal/http/middlewares"
	"botyard/internal/http/user"
	"botyard/internal/storage"
	"botyard/pkg/extlib"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	App   *fiber.App
	store storage.Storage
}

func New(store storage.Storage) *server {
	return &server{
		App: fiber.New(fiber.Config{
			ErrorHandler: errHandler,
			BodyLimit:    25 * 1024 * 1024, // 25 MB
		}),
		store: store,
	}
}

func (s *server) InitRoutes(prefix string) {
	api := s.App.Group(prefix)
	middlwr := middlewares.New(s.store)
	botServ := bot.NewService(s.store)
	userServ := user.NewService(s.store)
	fileServ := file.NewService(s.store)
	msgServ := message.NewService(s.store, fileServ)
	chatServ := chat.NewService(s.store, botServ, msgServ)

	bot := bot.Handlers(botServ)
	api.Get("/bot/:id", middlwr.Auth, bot.GetCommands)
	api.Post("/bot", middlwr.Admin, bot.Create)
	api.Post("/bot/commands/:id", middlwr.Admin, bot.AddCommands)
	api.Put("/bot/:id", middlwr.Admin, bot.Edit)
	api.Delete("/bot/command/:id", middlwr.Admin, bot.RemoveCommand)

	user := user.Handlers(userServ)
	api.Post("/user", user.Create)

	chat := chat.Handlers(chatServ)
	api.Post("/chat", middlwr.Auth, chat.Create)
	api.Delete("/chat/:id", middlwr.Auth, chat.Delete)

	file := file.Handlers(fileServ)
	api.Post("/files", middlwr.Auth, file.LoadFiles)

	// message := message.Handlers(s.store)
	// api.Get("/messages", middlwr.Auth, message.GetMessages)
	// api.Post("/message", middlwr.Auth, message.SendMessage)
}

func errHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *extlib.ExtendedError
	if errors.As(err, &e) {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(code).JSON(helpers.JsonMessage(err.Error()))
}
