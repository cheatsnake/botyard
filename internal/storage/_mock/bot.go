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

func (mbs *botStore) GetAuthKeyData(id string) (*bot.AuthKeyData, error) {
	return &bot.AuthKeyData{
		Key: "test",
	}, nil
}

func (mbs *botStore) SaveAuthKeyData(bkd *bot.AuthKeyData) error {
	return nil
}

func BotStore() *botStore {
	return &botStore{}
}
