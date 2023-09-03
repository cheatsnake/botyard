package http

import (
	"botyard/internal/http/handlers"
	"botyard/internal/http/helpers"
	"botyard/internal/http/middlewares"
	"botyard/internal/services"
	"botyard/internal/storage"

	"github.com/gofiber/fiber/v2"
)

const (
	apiV1             = "/v1"
	clientApiV1Prefix = apiV1 + "/client-api"
	adminApiV1Prefix  = apiV1 + "/admin-api"
	botApiV1Prefix    = apiV1 + "/bot-api"
)

type Server struct {
	App   *fiber.App
	store storage.Storage
}

func New(store storage.Storage) *Server {
	return &Server{
		App: fiber.New(fiber.Config{
			ErrorHandler: helpers.CursomErrorHandler,
			BodyLimit:    25 * 1024 * 1024, // 25 MB
		}),
		store: store,
	}
}

func (s *Server) InitRoutes() {
	botService := services.NewBotService(s.store)
	userService := services.NewUserService(s.store)
	fileService := services.NewFileService(s.store)
	messageService := services.NewMessageService(s.store, fileService)
	chatService := services.NewChatService(s.store, messageService)
	botMiddlewares := middlewares.NewBotMiddlewares(botService)
	bot := handlers.NewBotHandlers(botService)
	user := handlers.NewUserHandlers(userService)
	file := handlers.NewFileHandlers(fileService)
	message := handlers.NewMessageHandlers(messageService)
	chat := handlers.NewChatHandlers(chatService)

	// Admin API --------------------------------------------------------------
	adminApiV1 := s.App.Group(adminApiV1Prefix)

	adminApiV1.Get("/bots", middlewares.AdminAuth)
	adminApiV1.Get("/bot/:id/key", middlewares.AdminAuth)
	adminApiV1.Post("/bot", middlewares.AdminAuth, bot.Create)
	adminApiV1.Put("/bot/:id/key", middlewares.AdminAuth, bot.RefreshBotKey)
	adminApiV1.Delete("/bot/:id", middlewares.AdminAuth)
	// ------------------------------------------------------------------------

	// Bot API ----------------------------------------------------------------
	botApiV1 := s.App.Group(botApiV1Prefix)

	botApiV1.Get("/bot", botMiddlewares.Auth)
	botApiV1.Put("/bot", botMiddlewares.Auth, bot.Edit)
	botApiV1.Post("/bot/commands", botMiddlewares.Auth, bot.AddCommands)
	botApiV1.Delete("/bot/command", botMiddlewares.Auth, bot.RemoveCommand)

	botApiV1.Get("/messages/:chatId", middlewares.UserAuth, message.GetByChat)
	botApiV1.Post("/message", middlewares.UserAuth, message.Send)
	// ------------------------------------------------------------------------

	// Client API -------------------------------------------------------------
	clientApiV1 := s.App.Group(clientApiV1Prefix)

	clientApiV1.Post("/user", user.Create)

	clientApiV1.Get("/bot/:id", middlewares.UserAuth, bot.GetInfo)
	clientApiV1.Get("/bot/:id/commands", middlewares.UserAuth, bot.GetCommands)

	clientApiV1.Post("/files", middlewares.UserAuth, file.LoadMany)

	clientApiV1.Get("/messages/:chatId", middlewares.UserAuth, message.GetByChat)
	clientApiV1.Post("/message", middlewares.UserAuth, message.Send)

	clientApiV1.Get("/chats/:botId", middlewares.UserAuth, chat.GetMany)
	clientApiV1.Post("/chat", middlewares.UserAuth, chat.Create)
	clientApiV1.Delete("/chat/:id", middlewares.UserAuth, chat.Delete)
	//-------------------------------------------------------------------------
}
