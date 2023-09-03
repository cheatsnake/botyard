// Package memory represents a build-in in-memory database
// for development and testing purposes. Don't use it in production!
package memory

import (
	"botyard/internal/entities/bot"
	"botyard/internal/entities/chat"
	"botyard/internal/entities/file"
	"botyard/internal/entities/message"
	"botyard/internal/entities/user"
	"sync"
)

type Storage struct {
	bots        []*bot.Bot
	botAuthKeys []*bot.AuthKeyData
	users       []*user.User
	chats       []*chat.Chat
	messages    []*message.Message
	files       []*file.File
	mu          sync.Mutex
}

func New() *Storage {
	return &Storage{
		bots:        make([]*bot.Bot, 0, 10),
		botAuthKeys: make([]*bot.AuthKeyData, 0, 10),
		users:       make([]*user.User, 0, 10),
		chats:       make([]*chat.Chat, 0, 10),
		messages:    make([]*message.Message, 0, 10),
		files:       make([]*file.File, 0, 10),
	}
}
