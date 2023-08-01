package http

import (
	"botyard/internal/bot"
	"botyard/internal/chat"
	"botyard/internal/storage/memory"
	"botyard/internal/user"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App     *fiber.App
	Storage Storage
}

type Storage interface {
	AddBot(bot *bot.Bot) error
	GetBot(id string) (*bot.Bot, error)
	AddUser(user *user.User) error
	GetUser(id string) (*user.User, error)
	AddChat(chat *chat.Chat) error
	GetChat(id string) (*chat.Chat, error)
	FindChat(userId, botId string) (*chat.Chat, error)
	chat.MessageStore
	chat.FileStore
} 

func New(app *fiber.App, store *memory.Storage) *Server {
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
