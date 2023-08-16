package http

import (
	"botyard/internal/http/handlers"
	"botyard/internal/http/helpers"
	"botyard/internal/http/middlewares"
	"botyard/internal/services"
	"botyard/internal/storage"
	"botyard/pkg/extlib"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App   *fiber.App
	store storage.Storage
}

func New(store storage.Storage) *Server {
	return &Server{
		App: fiber.New(fiber.Config{
			ErrorHandler: errHandler,
			BodyLimit:    25 * 1024 * 1024, // 25 MB
		}),
		store: store,
	}
}

func (s *Server) InitRoutes(prefix string) {
	api := s.App.Group(prefix)
	middlwr := middlewares.New(s.store)

	bot := handlers.NewBotHandlers(services.NewBotService(s.store))
	api.Get("/bot/:id", middlwr.Auth, bot.GetCommands)
	api.Post("/bot", middlwr.Admin, bot.Create)
	api.Post("/bot/commands/:id", middlwr.Admin, bot.AddCommands)
	api.Put("/bot/:id", middlwr.Admin, bot.Edit)
	api.Delete("/bot/command/:id", middlwr.Admin, bot.RemoveCommand)

	user := handlers.NewUserHandlers(services.NewUserService(s.store))
	api.Post("/user", user.Create)

	fileService := services.NewFileService(s.store)
	file := handlers.NewFileHandlers(fileService)
	api.Post("/files", middlwr.Auth, file.LoadFiles)

	messageService := services.NewMessageService(s.store, fileService)
	// message := handlers.NewMessageHandlers(messageService)
	// api.Get("/messages", middlwr.Auth, message.GetMessages)
	// api.Post("/message", middlwr.Auth, message.SendMessage)

	chat := handlers.NewChatHandlers(services.NewChatService(s.store, messageService))
	// api.Get("/chats", middlwr.Auth, chat.GetChats)
	api.Post("/chat", middlwr.Auth, chat.Create)
	api.Delete("/chat/:id", middlwr.Auth, chat.Delete)
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
