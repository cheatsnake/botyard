package http

import (
	"botyard/internal/storage"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App     *fiber.App
	Storage storage.Storage
}

func New(app *fiber.App, store storage.Storage) *Server {
	return &Server{
		App:     app,
		Storage: store,
	}
}

func (s *Server) InitRoutes(prefix string) {
	s.App.Post(prefix+"/bot", s.createBot)
	s.App.Get(prefix+"/bot/:id", s.getBotCommands)

	s.App.Post(prefix+"/user", s.createUser)

	s.App.Post(prefix+"/chat", s.createChat)
	s.App.Post(prefix+"/chat/message", s.sendMessage)
	s.App.Post(prefix+"/chat/files", s.loadFiles)
	s.App.Get(prefix+"/chat/:id", s.getMessages)
	s.App.Delete(prefix+"/chat/:id", s.clearChat)
}
