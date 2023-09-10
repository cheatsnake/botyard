package memory

import (
	"botyard/internal/entities/bot"
	"botyard/pkg/extlib"

	"golang.org/x/exp/slices"
)

func (s *Storage) AddBot(bot *bot.Bot) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	candidate, _ := s.GetBot(bot.Name)
	if candidate != nil {
		return extlib.ErrorBadRequest("bot with this name already exists")
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

	return nil, extlib.ErrorNotFound("bot not found")
}

func (s *Storage) GetAllBots() ([]*bot.Bot, error) {
	return s.bots, nil
}

func (s *Storage) EditBot(bot *bot.Bot) error {
	return nil
}

func (s *Storage) DeleteBot(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delIndex := slices.IndexFunc(s.bots, func(b *bot.Bot) bool {
		return b.Id == id
	})

	if delIndex == -1 {
		return extlib.ErrorNotFound("bot not found")
	}

	s.bots = extlib.SliceRemoveElement(s.bots, delIndex)
	return nil
}

func (s *Storage) GetKey(botId string) (*bot.Key, error) {
	for _, bk := range s.botKeys {
		if bk.BotId == botId {
			return bk, nil
		}
	}

	return nil, extlib.ErrorNotFound("bot key not found")
}

func (s *Storage) SaveKey(botKey *bot.Key) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	existKey, _ := s.GetKey(botKey.BotId)

	if existKey != nil {
		existKey.Token = botKey.Token
		return nil
	}

	s.botKeys = append(s.botKeys, botKey)
	return nil
}

func (s *Storage) GetWebhook(botId string) (*bot.Webhook, error) {
	for _, wh := range s.botWebhooks {
		if wh.BotId == botId {
			return wh, nil
		}
	}

	return nil, extlib.ErrorNotFound("webhook not found")
}

func (s *Storage) SaveWebhook(webhook *bot.Webhook) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	existWebhook, _ := s.GetWebhook(webhook.BotId)

	if existWebhook != nil {
		existWebhook.Url = webhook.Url
		existWebhook.Secret = webhook.Secret
		return nil
	}

	s.botWebhooks = append(s.botWebhooks, webhook)
	return nil
}

func (s *Storage) DeleteWebhook(botId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delIndex := slices.IndexFunc(s.botWebhooks, func(w *bot.Webhook) bool {
		return w.BotId == botId
	})

	if delIndex == -1 {
		return extlib.ErrorNotFound("webhook not found")
	}

	s.botWebhooks = extlib.SliceRemoveElement(s.botWebhooks, delIndex)
	return nil
}
