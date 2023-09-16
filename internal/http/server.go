package http

import (
	"botyard/internal/http/handlers/bothandlers"
	"botyard/internal/http/handlers/chathandlers"
	"botyard/internal/http/handlers/filehandlers"
	"botyard/internal/http/handlers/messagehandlers"
	"botyard/internal/http/handlers/userhandlers"
	"botyard/internal/http/helpers"
	"botyard/internal/http/middlewares"
	"botyard/internal/services/botservice"
	"botyard/internal/services/chatservice"
	"botyard/internal/services/fileservice"
	"botyard/internal/services/messageservice"
	"botyard/internal/services/userservice"
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
	botService := botservice.New(s.store)
	userService := userservice.New(s.store)
	fileService := fileservice.New(s.store)
	messageService := messageservice.New(s.store, fileService)
	chatService := chatservice.New(s.store, messageService)
	botMiddlewares := middlewares.NewBotMiddlewares(botService)
	botHands := bothandlers.New(botService)
	userHands := userhandlers.New(userService)
	fileHands := filehandlers.New(fileService)
	messageHands := messagehandlers.New(messageService)
	chatHands := chathandlers.New(chatService)

	// Admin API --------------------------------------------------------------
	adminApiV1 := s.App.Group(adminApiV1Prefix)

	adminApiV1.Get("/bot/:id/key", middlewares.AdminAuth, botHands.GetKey)
	adminApiV1.Post("/bot", middlewares.AdminAuth, botHands.CreateBot)
	adminApiV1.Put("/bot/:id/key", middlewares.AdminAuth, botHands.RefreshKey)
	adminApiV1.Delete("/bot/:id", middlewares.AdminAuth)
	// ------------------------------------------------------------------------

	// Bot API ----------------------------------------------------------------
	botApiV1 := s.App.Group(botApiV1Prefix)

	botApiV1.Get("/bot", botMiddlewares.Auth)
	botApiV1.Put("/bot", botMiddlewares.Auth, botHands.EditBot)
	botApiV1.Post("/bot/commands", botMiddlewares.Auth, botHands.AddCommands)
	botApiV1.Delete("/bot/command", botMiddlewares.Auth, botHands.RemoveCommand)

	botApiV1.Get("/messages/:chatId", botMiddlewares.Auth, messageHands.GetByChat)
	botApiV1.Post("/message", botMiddlewares.Auth, messageHands.SendBotMessage)
	// ------------------------------------------------------------------------

	// Client API -------------------------------------------------------------
	clientApiV1 := s.App.Group(clientApiV1Prefix)

	clientApiV1.Post("/user", userHands.Create)

	clientApiV1.Get("/bots", botHands.GetAllBots)
	clientApiV1.Get("/bot/:id", botHands.GetBot)

	clientApiV1.Post("/files", middlewares.UserAuth, fileHands.LoadMany)

	clientApiV1.Get("/messages/:chatId", middlewares.UserAuth, messageHands.GetByChat)
	clientApiV1.Post("/message", middlewares.UserAuth, messageHands.SendUserMessage)

	clientApiV1.Get("/chats/:botId", middlewares.UserAuth, chatHands.GetMany)
	clientApiV1.Post("/chat", middlewares.UserAuth, chatHands.Create)
	clientApiV1.Delete("/chat/:id", middlewares.UserAuth, chatHands.Delete)
	//-------------------------------------------------------------------------
}
