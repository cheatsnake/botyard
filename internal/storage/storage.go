package storage

import (
	"botyard/internal/bot"
	"botyard/internal/chat"
	"botyard/internal/user"
)

type Storage interface {
	AddBot(bot *bot.Bot) error
	GetBot(id string) (*bot.Bot, error)

	AddUser(user *user.User) error
	GetUser(id string) (*user.User, error)

	AddChat(chat *chat.Chat) error
	GetChat(id string) (*chat.Chat, error)
	FindChat(userId, botId string) (*chat.Chat, error)

	AddMessage(msg *chat.Message) error
	GetMessagesByChat(chatId string, page, limit int) (int, []*chat.Message, error)
	DeleteMessagesByChat(chatId string) error

	AddFile(file *chat.File) error
	GetFile(id string) (*chat.File, error)
	GetFiles(ids []string) ([]*chat.File, error)
	DeleteFile(id string) error
}
