package storage

import (
	"botyard/internal/entities/bot"
	"botyard/internal/entities/chat"
	"botyard/internal/entities/file"
	"botyard/internal/entities/user"
)

type BotStore interface {
	AddBot(bot *bot.Bot) error
	EditBot(bot *bot.Bot) error
	GetBot(id string) (*bot.Bot, error)
	GetAllBots() ([]*bot.Bot, error)
	DeleteBot(id string) error

	SaveCommand(cmd *bot.Command) error
	GetCommands(botId string) ([]*bot.Command, error)
	DeleteCommand(botId, alias string) error
	DeleteCommandsByBot(botId string) error

	GetKey(botId string) (*bot.Key, error)
	SaveKey(key *bot.Key) error
	DeleteKey(botId string) error

	GetWebhook(botId string) (*bot.Webhook, error)
	SaveWebhook(wh *bot.Webhook) error
	DeleteWebhook(botId string) error
}

type UserStore interface {
	AddUser(user *user.User) error
	GetUser(id string) (*user.User, error)
	DeleteUser(id string) error
}

type ChatStore interface {
	AddChat(chat *chat.Chat) error
	GetChat(id string) (*chat.Chat, error)
	GetChats(userId, botId string) ([]*chat.Chat, error)
	DeleteChat(id string) error

	AddMessage(msg *chat.Message) error
	GetMessagesByChat(chatId string, page, limit int) (int, []*chat.Message, error)
	DeleteMessagesByChat(chatId string) error
}

type FileStore interface {
	AddFile(file *file.File) error
	GetFiles(ids []string) ([]*file.File, error)
	DeleteFiles(ids []string) error
}

type Storage interface {
	BotStore
	UserStore
	ChatStore
	FileStore
}
