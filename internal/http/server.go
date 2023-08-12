package http

import (
	"botyard/internal/http/handlers"
	"botyard/internal/http/middlewares"
	"botyard/internal/storage"

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
			ErrorHandler: handlers.ErrHandler,
			BodyLimit:    bodySizeLimit,
		}),
		store: store,
	}
}

func (s *Server) InitRoutes(prefix string) {
	api := s.App.Group(prefix)
	middlwr := middlewares.New(s.store)

	// Bot's handlers -----------------------------------
	bot := handlers.NewBot(s.store)
	api.Get("/bot/:id", middlwr.Auth, bot.GetBotCommands)
	api.Post("/bot", middlwr.Admin, bot.CreateBot)
	api.Post("/bot/commands/:id", middlwr.Admin, bot.AddBotCommands)
	api.Put("/bot/:id", middlwr.Admin, bot.EditBot)
	api.Delete("/bot/command/:id", middlwr.Admin, bot.RemoveBotCommand)
	// --------------------------------------------------

	// User's handlers ---------------
	user := handlers.NewUser(s.store)
	api.Post("/user", user.CreateUser)
	// -------------------------------

	// Chat's handlers --------------------------------------
	chat := handlers.NewChat(s.store)
	api.Get("/chat/:id", middlwr.Auth, chat.GetMessages)
	api.Post("/chat", middlwr.Auth, chat.CreateChat)
	api.Post("/chat/message", middlwr.Auth, chat.SendMessage)
	api.Post("/chat/files", middlwr.Auth, chat.LoadFiles)
	api.Delete("/chat/:id", middlwr.Auth, chat.ClearChat)
	// ------------------------------------------------------
}
