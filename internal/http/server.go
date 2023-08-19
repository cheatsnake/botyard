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

	botService := services.NewBotService(s.store)
	bot := handlers.NewBotHandlers(botService)
	api.Get("/bot/:id", middlewares.Auth, bot.GetCommands)
	api.Post("/bot", middlewares.Admin, bot.Create)
	api.Post("/bot/commands/:id", middlewares.Admin, bot.AddCommands)
	api.Put("/bot/:id", middlewares.Admin, bot.Edit)
	api.Delete("/bot/command/:id", middlewares.Admin, bot.RemoveCommand)

	userService := services.NewUserService(s.store)
	user := handlers.NewUserHandlers(userService)
	api.Post("/user", user.Create)

	fileService := services.NewFileService(s.store)
	file := handlers.NewFileHandlers(fileService)
	api.Post("/files", middlewares.Auth, file.LoadMany)

	messageService := services.NewMessageService(s.store, fileService)
	message := handlers.NewMessageHandlers(messageService)
	api.Get("/messages/:chatId", middlewares.Auth, message.GetByChat)
	api.Post("/message", middlewares.Auth, message.Send)

	chatService := services.NewChatService(s.store, messageService)
	chat := handlers.NewChatHandlers(chatService)
	api.Get("/chats/:botId", middlewares.Auth, chat.GetMany)
	api.Post("/chat", middlewares.Auth, chat.Create)
	api.Delete("/chat/:id", middlewares.Auth, chat.Delete)
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
