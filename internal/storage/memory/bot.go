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

func (s *Storage) GetBotKeyData(id string) (*bot.BotKeyData, error) {
	for _, bkd := range s.botKeysData {
		if bkd.BotId == id {
			return bkd, nil
		}
	}

	return nil, nil
}

func (s *Storage) SaveBotKeyData(newBkd *bot.BotKeyData) error {
	bkd, err := s.GetBotKeyData(newBkd.BotId)
	if err != nil {
		return err
	}

	if bkd != nil {
		bkd.Key = newBkd.Key
		return nil
	}

	s.botKeysData = append(s.botKeysData, newBkd)
	return nil
}
