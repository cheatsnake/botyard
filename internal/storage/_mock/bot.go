package mock

import "botyard/internal/entities/bot"

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

func (mbs *botStore) DeleteBot(id string) error {
	return nil
}

func (mbs *botStore) GetBotKeyData(id string) (*bot.BotKeyData, error) {
	return &bot.BotKeyData{
		Key: "test",
	}, nil
}

func (mbs *botStore) SaveBotKeyData(bkd *bot.BotKeyData) error {
	return nil
}

func BotStore() *botStore {
	return &botStore{}
}
