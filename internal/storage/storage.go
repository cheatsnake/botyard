package storage

import (
	"botyard/internal/entities/bot"
	"botyard/internal/entities/chat"
	"botyard/internal/entities/file"
	"botyard/internal/entities/message"
	"botyard/internal/entities/user"
)

type BotStore interface {
	AddBot(bot *bot.Bot) error
	EditBot(bot *bot.Bot) error
	GetBot(id string) (*bot.Bot, error)
	GetAllBots() ([]*bot.Bot, error)
	DeleteBot(id string) error
	GetKeyData(id string) (*bot.KeyData, error)
	SaveKeyData(bkd *bot.KeyData) error
}

type UserStore interface {
	AddUser(user *user.User) error
	GetUser(id string) (*user.User, error)
}

type ChatStore interface {
	AddChat(chat *chat.Chat) error
	GetChat(id string) (*chat.Chat, error)
	GetChats(userId, botId string) ([]*chat.Chat, error)
	DeleteChat(id string) error
}

type MessageStore interface {
	AddMessage(msg *message.Message) error
	GetMessagesByChat(chatId string, page, limit int) (int, []*message.Message, error)
	DeleteMessagesByChat(chatId string) error
}

type FileStore interface {
	AddFile(file *file.File) error
	GetFile(id string) (*file.File, error)
	GetFiles(ids []string) ([]*file.File, error)
	DeleteFile(id string) error
}

type Storage interface {
	BotStore
	UserStore
	ChatStore
	MessageStore
	FileStore
}
