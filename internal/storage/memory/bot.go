package memory

import (
	"github.com/cheatsnake/botyard/internal/entities/bot"
	"github.com/cheatsnake/botyard/pkg/exterr"
	"github.com/cheatsnake/botyard/pkg/extlib"

	"golang.org/x/exp/slices"
)

func (s *Storage) AddBot(newBot *bot.Bot) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, b := range s.bots {
		if b.Name == newBot.Name {
			return exterr.ErrorBadRequest("bot with this name already exists")
		}
	}

	s.bots = append(s.bots, newBot)
	return nil
}

func (s *Storage) GetBot(id string) (*bot.Bot, error) {
	for _, bot := range s.bots {
		if bot.Id == id {
			return bot, nil
		}
	}

	return nil, exterr.ErrorNotFound("bot not found")
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
		return exterr.ErrorNotFound("bot not found")
	}

	s.bots = extlib.SliceRemoveElement(s.bots, delIndex)
	return nil
}

func (s *Storage) SaveCommand(newCmd *bot.Command) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, c := range s.botCommands {
		if c.BotId == newCmd.BotId && c.Alias == newCmd.Alias {
			c.Description = newCmd.Description
			return nil
		}
	}

	s.botCommands = append(s.botCommands, newCmd)
	return nil
}

func (s *Storage) GetCommands(botId string) ([]*bot.Command, error) {
	cmds := []*bot.Command{}

	for _, cmd := range s.botCommands {
		if cmd.BotId == botId {
			cmds = append(cmds, cmd)
		}
	}

	return cmds, nil
}

func (s *Storage) DeleteCommand(botId, alias string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delIndex := slices.IndexFunc(s.botCommands, func(cmd *bot.Command) bool {
		return cmd.BotId == botId && cmd.Alias == alias
	})

	if delIndex == -1 {
		return exterr.ErrorNotFound("command not found")
	}

	s.botCommands = extlib.SliceRemoveElement(s.botCommands, delIndex)
	return nil
}

func (s *Storage) DeleteCommandsByBot(botId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	filtered := extlib.SliceFilter(s.botCommands, 0, func(cmd *bot.Command) bool {
		return cmd.BotId != botId
	})

	s.botCommands = filtered
	return nil
}

func (s *Storage) GetKey(botId string) (*bot.Key, error) {
	for _, bk := range s.botKeys {
		if bk.BotId == botId {
			return bk, nil
		}
	}

	return nil, exterr.ErrorNotFound("bot key not found")
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

func (s *Storage) DeleteKey(botId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delIndex := slices.IndexFunc(s.botKeys, func(k *bot.Key) bool {
		return k.BotId == botId
	})

	if delIndex == -1 {
		return exterr.ErrorNotFound("bot key not found")
	}

	s.botKeys = extlib.SliceRemoveElement(s.botKeys, delIndex)
	return nil
}

func (s *Storage) GetWebhook(botId string) (*bot.Webhook, error) {
	for _, wh := range s.botWebhooks {
		if wh.BotId == botId {
			return wh, nil
		}
	}

	return nil, exterr.ErrorNotFound("webhook not found")
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
		return exterr.ErrorNotFound("webhook not found")
	}

	s.botWebhooks = extlib.SliceRemoveElement(s.botWebhooks, delIndex)
	return nil
}
