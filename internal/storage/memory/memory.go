// Package memory represents a build-in in-memory database for development
// and testing purposes. Don't use it in production!
package memory

import (
	"botyard/internal/bot"
	"botyard/internal/chat"
	"botyard/internal/user"
	"sync"
)

type Storage struct {
	bots     []*bot.Bot
	users    []*user.User
	chats    []*chat.Chat
	messages []*chat.Message
	files    []*chat.File
	mu       sync.Mutex
}

func New() *Storage {
	return &Storage{
		bots:     make([]*bot.Bot, 0, 10),
		users:    make([]*user.User, 0, 10),
		chats:    make([]*chat.Chat, 0, 10),
		messages: make([]*chat.Message, 0, 10),
		files:    make([]*chat.File, 0, 10),
	}
}
