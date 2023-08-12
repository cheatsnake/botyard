package http

import (
	"botyard/internal/http/handlers"
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
	bot := handlers.NewBot(s.store)
	user := handlers.NewUser(s.store)
	chat := handlers.NewChat(s.store)

	s.App.Post(prefix+"/bot", bot.CreateBot)
	s.App.Get(prefix+"/bot/:id", bot.GetBotCommands)

	s.App.Post(prefix+"/user", user.CreateUser)

	s.App.Post(prefix+"/chat", chat.CreateChat)
	s.App.Post(prefix+"/chat/message", chat.SendMessage)
	s.App.Post(prefix+"/chat/files", chat.LoadFiles)
	s.App.Get(prefix+"/chat/:id", chat.GetMessages)
	s.App.Delete(prefix+"/chat/:id", chat.ClearChat)
}
