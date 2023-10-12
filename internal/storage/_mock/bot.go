package mock

import "github.com/cheatsnake/botyard/internal/entities/bot"

type botStore struct{}

func (mbs *botStore) AddBot(bot *bot.Bot) error {
	return nil
}

func (mbs *botStore) EditBot(bot *bot.Bot) error {
	return nil
}

func (mbs *botStore) GetBot(id string) (*bot.Bot, error) {
	return &bot.Bot{}, nil
}

func (mbs *botStore) GetAllBots() ([]*bot.Bot, error) {
	return []*bot.Bot{}, nil
}

func (mbs *botStore) DeleteBot(id string) error {
	return nil
}

func (mbs *botStore) SaveCommand(cmd *bot.Command) error {
	return nil
}

func (mbs *botStore) GetCommands(botId string) ([]*bot.Command, error) {
	return []*bot.Command{}, nil
}

func (mbs *botStore) GetCommand(id string) (*bot.Command, error) {
	return &bot.Command{}, nil
}

func (mbs *botStore) DeleteCommand(id string) error {
	return nil
}

func (mbs *botStore) DeleteCommandsByBot(botId string) error {
	return nil
}

func (mbs *botStore) GetKey(id string) (*bot.Key, error) {
	return &bot.Key{
		Token: "test",
	}, nil
}

func (mbs *botStore) SaveKey(bkd *bot.Key) error {
	return nil
}

func (mbs *botStore) DeleteKey(botId string) error {
	return nil
}

func (mbs *botStore) GetWebhook(id string) (*bot.Webhook, error) {
	return &bot.Webhook{}, nil
}

func (mbs *botStore) SaveWebhook(webhook *bot.Webhook) error {
	return nil
}

func (mbs *botStore) DeleteWebhook(id string) error {
	return nil
}

func BotStore() *botStore {
	return &botStore{}
}
