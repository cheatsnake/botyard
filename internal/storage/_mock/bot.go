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

func (mbs *botStore) GetAllBots() ([]*bot.Bot, error) {
	return []*bot.Bot{}, nil
}

func (mbs *botStore) DeleteBot(id string) error {
	return nil
}

func (mbs *botStore) GetKeyData(id string) (*bot.KeyData, error) {
	return &bot.KeyData{
		Key: "test",
	}, nil
}

func (mbs *botStore) SaveKeyData(bkd *bot.KeyData) error {
	return nil
}

func BotStore() *botStore {
	return &botStore{}
}
