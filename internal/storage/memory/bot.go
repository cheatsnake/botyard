package memory

import (
	"botyard/internal/entities/bot"
	"botyard/pkg/extlib"
	"errors"
)

func (s *Storage) AddBot(bot *bot.Bot) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	candidate, _ := s.GetBot(bot.Name)
	if candidate != nil {
		// TODO consts for errors
		return errors.New("bot with this name already exists")
	}

	s.bots = append(s.bots, bot)
	return nil
}

func (s *Storage) GetBot(id string) (*bot.Bot, error) {
	for _, bot := range s.bots {
		if bot.Id == id {
			return bot, nil
		}
	}

	return nil, errors.New("bot not found")
}

func (s *Storage) GetAllBots() ([]*bot.Bot, error) {
	return s.bots, nil
}

func (s *Storage) EditBot(bot *bot.Bot) error {
	return nil
}

func (s *Storage) DeleteBot(id string) error {
	for i, bot := range s.bots {
		if bot.Id == id {
			s.bots = extlib.SliceRemoveElement(s.bots, i)
			return nil
		}
	}

	return errors.New("bot not found")
}

func (s *Storage) GetKey(id string) (*bot.Key, error) {
	for _, bk := range s.botKeys {
		if bk.BotId == id {
			return bk, nil
		}
	}

	return nil, nil
}

func (s *Storage) SaveKey(botKey *bot.Key) error {
	existKey, err := s.GetKey(botKey.BotId)
	if err != nil {
		return err
	}

	if existKey != nil {
		existKey.Token = botKey.Token
		return nil
	}

	s.botKeys = append(s.botKeys, botKey)
	return nil
}

func (s *Storage) GetWebhook(id string) (*bot.Webhook, error) {
	for _, wh := range s.botWebhooks {
		if wh.BotId == id {
			return wh, nil
		}
	}

	return nil, nil
}

func (s *Storage) SaveWebhook(webhook *bot.Webhook) error {
	existWebhook, err := s.GetWebhook(webhook.BotId)
	if err != nil {
		return err
	}

	if existWebhook != nil {
		existWebhook.Url = webhook.Url
		existWebhook.Secret = webhook.Secret
		return nil
	}

	s.botWebhooks = append(s.botWebhooks, webhook)
	return nil
}
