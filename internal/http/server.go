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
	api.Post("/bot", middlwr.Admin, bot.CreateBot)
	api.Get("/bot/:id", middlwr.Auth, bot.GetBotCommands)
	// --------------------------------------------------

	// User's handlers ---------------
	user := handlers.NewUser(s.store)
	api.Post("/user", user.CreateUser)
	// -------------------------------

	// Chat's handlers --------------------------------------
	chat := handlers.NewChat(s.store)
	api.Post("/chat", middlwr.Auth, chat.CreateChat)
	api.Post("/chat/message", middlwr.Auth, chat.SendMessage)
	api.Post("/chat/files", middlwr.Auth, chat.LoadFiles)
	api.Get("/chat/:id", middlwr.Auth, chat.GetMessages)
	api.Delete("/chat/:id", middlwr.Auth, chat.ClearChat)
	// ------------------------------------------------------
}
