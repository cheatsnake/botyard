package http

import (
	"botyard/internal/storage/memory"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App     *fiber.App
	Storage *memory.Storage
}

func New(app *fiber.App, store *memory.Storage) *Server {
	return &Server{
		App:     app,
		Storage: store,
	}
}

func (s *Server) InitRoutes(prefix string) {
	// Bot
	s.App.Post(prefix+"/bot", s.createBot)
	s.App.Get(prefix+"/bot/:id", s.getBotCommands)

	// User
	s.App.Post(prefix+"/user", s.createUser)

	// Chat
	s.App.Post(prefix+"/chat", s.createChat)
	s.App.Post(prefix+"/chat/message", s.sendMessage)
	s.App.Get(prefix+"/chat/:id", s.getMessages)
	s.App.Delete(prefix+"/chat/:id", s.clearChat)
}
