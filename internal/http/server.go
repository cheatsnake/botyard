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
	botMiddlewares := middlewares.NewBotMiddlewares(botService)
	bot := handlers.NewBotHandlers(botService)
	api.Post("/bot", middlewares.AdminAuth, bot.Create)
	api.Put("/bot/key", middlewares.AdminAuth, bot.RefreshBotKey)
	api.Put("/bot", botMiddlewares.Auth, bot.Edit)
	api.Delete("/bot", botMiddlewares.Auth, bot.RemoveCommand)

	api.Get("/bot/commands/:id", middlewares.UserAuth, bot.GetCommands)
	api.Post("/bot/commands", botMiddlewares.Auth, bot.AddCommands)
	api.Delete("/bot/command", botMiddlewares.Auth, bot.RemoveCommand)

	userService := services.NewUserService(s.store)
	user := handlers.NewUserHandlers(userService)
	api.Post("/user", user.Create)

	fileService := services.NewFileService(s.store)
	file := handlers.NewFileHandlers(fileService)
	api.Post("/files", middlewares.UserAuth, file.LoadMany)

	messageService := services.NewMessageService(s.store, fileService)
	message := handlers.NewMessageHandlers(messageService)
	api.Get("/messages/:chatId", middlewares.UserAuth, message.GetByChat)
	api.Post("/message", middlewares.UserAuth, message.Send)

	chatService := services.NewChatService(s.store, messageService)
	chat := handlers.NewChatHandlers(chatService)
	api.Get("/chats/:botId", middlewares.UserAuth, chat.GetMany)
	api.Post("/chat", middlewares.UserAuth, chat.Create)
	api.Delete("/chat/:id", middlewares.UserAuth, chat.Delete)
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
