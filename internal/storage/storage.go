package storage

import (
	"botyard/internal/entities/bot"
	"botyard/internal/entities/chat"
	"botyard/internal/entities/file"
	"botyard/internal/entities/message"
	"botyard/internal/entities/user"
)

type Storage interface {
	AddBot(bot *bot.Bot) error
	EditBot(bot *bot.Bot) error
	GetBot(id string) (*bot.Bot, error)

	AddUser(user *user.User) error
	GetUser(id string) (*user.User, error)

	AddChat(chat *chat.Chat) error
	GetChat(id string) (*chat.Chat, error)
	FindChat(userId, botId string) (*chat.Chat, error)

	AddMessage(msg *message.Message) error
	GetMessagesByChat(chatId string, page, limit int) (int, []*message.Message, error)
	DeleteMessagesByChat(chatId string) error

	AddFile(file *file.File) error
	GetFile(id string) (*file.File, error)
	GetFiles(ids []string) ([]*file.File, error)
	DeleteFile(id string) error
}
