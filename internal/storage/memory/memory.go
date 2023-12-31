// Package memory represents a build-in in-memory database
// for development and testing purposes. Don't use it in production!
package memory

import (
	"sync"

	"github.com/cheatsnake/botyard/internal/entities/bot"
	"github.com/cheatsnake/botyard/internal/entities/chat"
	"github.com/cheatsnake/botyard/internal/entities/file"
	"github.com/cheatsnake/botyard/internal/entities/user"
)

type Storage struct {
	bots        []*bot.Bot
	botCommands []*bot.Command
	botKeys     []*bot.Key
	botWebhooks []*bot.Webhook
	users       []*user.User
	chats       []*chat.Chat
	messages    []*chat.Message
	files       []*file.File
	mu          sync.Mutex
}

func New() *Storage {
	return &Storage{
		bots:        make([]*bot.Bot, 0, 10),
		botCommands: make([]*bot.Command, 0, 10),
		botKeys:     make([]*bot.Key, 0, 10),
		botWebhooks: make([]*bot.Webhook, 0, 10),
		users:       make([]*user.User, 0, 10),
		chats:       make([]*chat.Chat, 0, 10),
		messages:    make([]*chat.Message, 0, 10),
		files:       make([]*file.File, 0, 10),
	}
}
