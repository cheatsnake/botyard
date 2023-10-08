package botservice

import (
	"github.com/cheatsnake/botyard/internal/entities/bot"
	"github.com/cheatsnake/botyard/internal/storage"
	"github.com/cheatsnake/botyard/internal/tools/ulid"
	"github.com/cheatsnake/botyard/pkg/exterr"
	"github.com/cheatsnake/botyard/pkg/extlib"
)

type Service struct {
	store storage.BotStore
}

type CreateBody struct {
	bot.Bot
	Id struct{} `json:"-"`
}

type CreateResult struct {
	Bot *bot.Bot `json:"bot"`
	Key string   `json:"key"`
}

type EditBody struct {
	bot.Bot
	Id struct{} `json:"-"`
}

type PreparedCommand struct {
	Alias       string `json:"alias"`
	Description string `json:"description,omitempty"`
}

type CommandsBody struct {
	Commands []PreparedCommand `json:"commands"`
}

type CommandsList struct {
	BotId    string            `json:"botId"`
	Commands []PreparedCommand `json:"commands"`
}

type WebhookBody struct {
	bot.Webhook
	BotId struct{} `json:"-"`
}

func New(s storage.BotStore) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) CreateBot(body *CreateBody) (*CreateResult, error) {
	newBot, err := bot.New(body.Name)
	if err != nil {
		return nil, exterr.ErrorBadRequest(err.Error())
	}

	if body.Description != "" {
		newBot.SetDescription(body.Description)
	}

	if body.Avatar != "" {
		newBot.SetAvatar(body.Avatar)
	}

	err = s.store.AddBot(newBot)
	if err != nil {
		return nil, err
	}

	botKey, err := s.GenerateKey(newBot.Id)
	if err != nil {
		return nil, exterr.ErrorBadRequest(err.Error())
	}

	result := &CreateResult{
		Bot: newBot,
		Key: botKey,
	}
	return result, nil
}

func (s *Service) GetBotById(id string) (*bot.Bot, error) {
	foundBot, err := s.store.GetBot(id)
	if err != nil {
		return nil, err
	}

	return foundBot, nil
}

func (s *Service) GetAllBots() ([]*bot.Bot, error) {
	bots, err := s.store.GetAllBots()
	if err != nil {
		return nil, err
	}

	return bots, nil
}

func (s *Service) EditBot(id string, body *EditBody) (*bot.Bot, error) {
	foundBot, err := s.GetBotById(id)
	if err != nil {
		return nil, err
	}

	if body.Name != "" {
		err := foundBot.SetName(body.Name)
		if err != nil {
			return nil, exterr.ErrorBadRequest(err.Error())
		}
	}

	if body.Description != "" {
		err := foundBot.SetDescription(body.Description)
		if err != nil {
			return nil, exterr.ErrorBadRequest(err.Error())
		}
	}

	if body.Avatar != "" {
		err := foundBot.SetAvatar(body.Avatar)
		if err != nil {
			return nil, exterr.ErrorBadRequest(err.Error())
		}
	}

	err = s.store.EditBot(foundBot)
	if err != nil {
		return nil, err
	}

	return foundBot, nil
}

func (s *Service) DeleteBot(id string) error {
	_, err := s.store.GetBot(id)
	if err != nil {
		return err
	}

	err = s.store.DeleteBot(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddCommands(botId string, body *CommandsBody) error {
	_, err := s.GetBotById(botId)
	if err != nil {
		return err
	}

	for _, c := range body.Commands {
		newCmd, err := bot.NewCommand(botId, c.Alias, c.Description)
		if err != nil {
			return exterr.ErrorBadRequest(err.Error())
		}

		err = s.store.SaveCommand(newCmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) GetCommands(botId string) (*CommandsList, error) {
	_, err := s.GetBotById(botId)
	if err != nil {
		return nil, err
	}

	cmds, err := s.store.GetCommands(botId)
	if err != nil {
		return nil, err
	}

	result := CommandsList{
		BotId:    botId,
		Commands: make([]PreparedCommand, 0, len(cmds)),
	}

	for _, cmd := range cmds {
		result.Commands = append(result.Commands, PreparedCommand{
			Alias:       cmd.Alias,
			Description: cmd.Description,
		})
	}

	return &result, nil
}

func (s *Service) RemoveCommand(botId string, alias string) error {
	_, err := s.GetBotById(botId)
	if err != nil {
		return err
	}

	err = s.store.DeleteCommand(botId, alias)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetKey(id string) (string, error) {
	keyData, err := s.store.GetKey(id)
	if err != nil {
		return "", err
	}

	return keyData.Assemble(), nil
}

func (s *Service) GenerateKey(id string) (string, error) {
	token, err := extlib.RandomToken(ulid.Length)
	if err != nil {
		return "", exterr.ErrorBadRequest("Bot key generation failed: " + err.Error())
	}

	botKey, err := bot.NewKey(id, token)
	if err != nil {
		return "", exterr.ErrorBadRequest(err.Error())
	}

	err = s.store.SaveKey(botKey)
	if err != nil {
		return "", err
	}

	return botKey.Assemble(), nil
}

func (s *Service) VerifyKeyData(id, token string) error {
	kd, err := s.store.GetKey(id)
	if err != nil {
		return nil
	}

	if kd == nil || kd.Token != token {
		return exterr.ErrorForbidden("Invalid bot key.")
	}

	return nil
}

func (s *Service) SaveWebhook(botId string, body *WebhookBody) (*bot.Webhook, error) {
	_, err := s.GetBotById(botId)
	if err != nil {
		return nil, err
	}

	webhook, err := bot.NewWebhook(botId, body.Url, body.Secret)
	if err != nil {
		return nil, exterr.ErrorBadRequest(err.Error())
	}

	err = s.store.SaveWebhook(webhook)
	if err != nil {
		return nil, exterr.ErrorBadRequest(err.Error())
	}

	return webhook, nil
}

func (s *Service) GetWebhook(botId string) (*bot.Webhook, error) {
	wh, err := s.store.GetWebhook(botId)
	if err != nil {
		return nil, err
	}

	return wh, nil
}

func (s *Service) DeleteWebhook(botId string) error {
	_, err := s.GetWebhook(botId)
	if err != nil {
		return err
	}

	err = s.store.DeleteWebhook(botId)
	if err != nil {
		return err
	}

	return nil
}
