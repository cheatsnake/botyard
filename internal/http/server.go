package http

import (
	"os"
	"path"

	"github.com/cheatsnake/botyard/internal/config"
	"github.com/cheatsnake/botyard/internal/http/handlers/adminhandlers"
	"github.com/cheatsnake/botyard/internal/http/handlers/bothandlers"
	"github.com/cheatsnake/botyard/internal/http/handlers/chathandlers"
	"github.com/cheatsnake/botyard/internal/http/handlers/filehandlers"
	"github.com/cheatsnake/botyard/internal/http/handlers/userhandlers"
	"github.com/cheatsnake/botyard/internal/http/helpers"
	"github.com/cheatsnake/botyard/internal/http/middlewares"
	"github.com/cheatsnake/botyard/internal/services/botservice"
	"github.com/cheatsnake/botyard/internal/services/chatservice"
	"github.com/cheatsnake/botyard/internal/services/fileservice"
	"github.com/cheatsnake/botyard/internal/services/userservice"
	"github.com/cheatsnake/botyard/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	clientApiPrefix = "/client-api"
	adminApiPrefix  = "/admin-api"
	botApiPrefix    = "/bot-api"
	apiV1Prefix     = "/v1"
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

	// Middlewares ------------------------------------------------------------
	s.App.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost, http://127.0.0.1",
	}))
	// ------------------------------------------------------------------------

	// Admin API --------------------------------------------------------------
	adminApiV1 := s.App.Group(apiV1Prefix + adminApiPrefix)

	adminApiV1.Post("/bot", middlewares.AdminAuth, botHands.CreateBot)
	adminApiV1.Delete("/bot/:id", middlewares.AdminAuth, botHands.DeleteBot)

	adminApiV1.Get("/bot/:id/key", middlewares.AdminAuth, botHands.GetKey)
	adminApiV1.Put("/bot/:id/key", middlewares.AdminAuth, botHands.RefreshKey)

	adminApiV1.Put("/config", middlewares.AdminAuth, adminhandlers.ReloadGlobalConfig)
	// ------------------------------------------------------------------------

	// Bot API ----------------------------------------------------------------
	botApiV1 := s.App.Group(apiV1Prefix + botApiPrefix)

	botApiV1.Get("/bot", botMiddlewares.Auth, botHands.GetCurrentBot)
	botApiV1.Put("/bot", botMiddlewares.Auth, botHands.EditBot)

	botApiV1.Get("/bot/commands", botMiddlewares.Auth, botHands.GetCommands)
	botApiV1.Post("/bot/commands", botMiddlewares.Auth, botHands.AddCommands)
	botApiV1.Delete("/bot/command/:id", botMiddlewares.Auth, botHands.RemoveCommand)

	botApiV1.Get("/bot/webhook", botMiddlewares.Auth, botHands.GetWebhook)
	botApiV1.Post("/bot/webhook", botMiddlewares.Auth, botHands.SaveWebhook)
	botApiV1.Delete("/bot/webhook", botMiddlewares.Auth, botHands.DeleteWebhook)

	botApiV1.Get("/chats", botMiddlewares.Auth, chatHands.GetChatsByUser)
	botApiV1.Get("/chat/:id/messages", botMiddlewares.Auth, chatHands.GetMessagesByChat)
	botApiV1.Get("/chat/message/:id", botMiddlewares.Auth, chatHands.GetMessage)
	botApiV1.Post("/chat/message", botMiddlewares.Auth, chatHands.SendBotMessage)

	botApiV1.Post("/files", botMiddlewares.Auth, fileHands.UploadFiles)
	// ------------------------------------------------------------------------

	// Client API -------------------------------------------------------------
	clientApiV1 := s.App.Group(apiV1Prefix + clientApiPrefix)

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

	clientApiV1.Post("/files", middlewares.UserAuth, fileHands.UploadFiles)
	//-------------------------------------------------------------------------

	s.App.Static("/static", path.Join(".", os.Getenv("FILES_FOLDER")))
	s.App.Static("/", path.Join(".", "web", "dist"))
	s.App.Static("*", path.Join(".", "web", "dist"))
}
