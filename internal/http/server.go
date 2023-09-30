package http

import (
	"botyard/internal/config"
	"botyard/internal/http/handlers/adminhandlers"
	"botyard/internal/http/handlers/bothandlers"
	"botyard/internal/http/handlers/chathandlers"
	"botyard/internal/http/handlers/filehandlers"
	"botyard/internal/http/handlers/userhandlers"
	"botyard/internal/http/helpers"
	"botyard/internal/http/middlewares"
	"botyard/internal/services/botservice"
	"botyard/internal/services/chatservice"
	"botyard/internal/services/fileservice"
	"botyard/internal/services/userservice"
	"botyard/internal/storage"
	"os"
	"path"

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
			BodyLimit: max(config.Global.Limits.File.MaxImageSize,
				config.Global.Limits.File.MaxAudioSize,
				config.Global.Limits.File.MaxVideoSize,
				config.Global.Limits.File.MaxFileSize,
			),
		}),
		store: store,
	}
}

func (s *Server) InitRoutes() {
	botService := botservice.New(s.store)
	userService := userservice.New(s.store)
	fileService := fileservice.New(s.store)
	chatService := chatservice.New(s.store, fileService)
	botMiddlewares := middlewares.NewBotMiddlewares(botService)
	botHands := bothandlers.New(botService)
	userHands := userhandlers.New(userService)
	fileHands := filehandlers.New(fileService)
	chatHands := chathandlers.New(chatService, botService)

	// Admin API --------------------------------------------------------------
	adminApiV1 := s.App.Group(adminApiV1Prefix)

	adminApiV1.Post("/bot", middlewares.AdminAuth, botHands.CreateBot)
	adminApiV1.Delete("/bot/:id", middlewares.AdminAuth)

	adminApiV1.Get("/bot/:id/key", middlewares.AdminAuth, botHands.GetKey)
	adminApiV1.Put("/bot/:id/key", middlewares.AdminAuth, botHands.RefreshKey)

	adminApiV1.Put("/config", middlewares.AdminAuth, adminhandlers.ReloadGlobalConfig)
	// ------------------------------------------------------------------------

	// Bot API ----------------------------------------------------------------
	botApiV1 := s.App.Group(botApiV1Prefix)

	botApiV1.Get("/bot", botMiddlewares.Auth, botHands.GetCurrentBot)
	botApiV1.Put("/bot", botMiddlewares.Auth, botHands.EditBot)

	botApiV1.Get("/bot/commands", botMiddlewares.Auth, botHands.GetCommands)
	botApiV1.Post("/bot/commands", botMiddlewares.Auth, botHands.AddCommands)
	botApiV1.Delete("/bot/command", botMiddlewares.Auth, botHands.RemoveCommand)

	botApiV1.Get("/bot/webhook", botMiddlewares.Auth, botHands.GetWebhook)
	botApiV1.Post("/bot/webhook", botMiddlewares.Auth, botHands.CreateWebhook)
	botApiV1.Put("/bot/webhook", botMiddlewares.Auth, botHands.EditWebhook)
	botApiV1.Delete("/bot/webhook", botMiddlewares.Auth, botHands.DeleteWebhook)

	botApiV1.Get("/chats", middlewares.UserAuth, chatHands.GetChatsByUser)
	botApiV1.Get("/chat/:id/messages", botMiddlewares.Auth, chatHands.GetMessagesByChat)
	botApiV1.Post("/chat/message", botMiddlewares.Auth, chatHands.SendBotMessage)

	botApiV1.Post("/files", botMiddlewares.Auth, fileHands.LoadMany)
	// ------------------------------------------------------------------------

	// Client API -------------------------------------------------------------
	clientApiV1 := s.App.Group(clientApiV1Prefix)

	clientApiV1.Get("/service-info", func(c *fiber.Ctx) error {
		return c.JSON(config.Global.Service)
	})

	clientApiV1.Get("/user", middlewares.UserAuth, userHands.GetCurrentUser)
	clientApiV1.Post("/user", userHands.CreateUser)

	clientApiV1.Get("/bots", botHands.GetAllBots)
	clientApiV1.Get("/bot/:id", botHands.GetBotById)
	clientApiV1.Get("/bot/:id/commands", middlewares.UserAuth, botHands.GetCommands)

	clientApiV1.Get("/chats", middlewares.UserAuth, chatHands.GetChatsByBot)
	clientApiV1.Get("/chat/:id/messages", middlewares.UserAuth, chatHands.GetMessagesByChat)
	clientApiV1.Post("/chat", middlewares.UserAuth, chatHands.CreateChat)
	clientApiV1.Post("/chat/message", middlewares.UserAuth, chatHands.SendUserMessage)
	clientApiV1.Delete("/chat/:id", middlewares.UserAuth, chatHands.DeleteChat)

	clientApiV1.Post("/files", middlewares.UserAuth, fileHands.LoadMany)
	//-------------------------------------------------------------------------

	s.App.Static("/static", path.Join(".", os.Getenv("FILES_FOLDER")))
}
