package services

import (
	"botyard/internal/entities/bot"
	"botyard/internal/storage"
	"botyard/internal/tools/ulid"
	"botyard/pkg/extlib"
)

type BotService struct {
	store storage.BotStore
}

type BotCreateBody struct {
	bot.Bot
	Id struct{} `json:"-"`
}

type BotCreateResult struct {
	Bot *bot.Bot      `json:"bot"`
	Key *BotKeyResult `json:"key"`
}

type BotKeyResult struct {
	Value string `json:"value"`
}

type BotEditBody struct {
	bot.Bot
	Id struct{} `json:"-"`
}

type BotCommandsBody struct {
	Commands []struct {
		bot.Command
		BotId struct{} `json:"-"`
	} `json:"commands"`
}

func NewBotService(s storage.BotStore) *BotService {
	return &BotService{
		store: s,
	}
}

func (s *BotService) CreateBot(body *BotCreateBody) (*BotCreateResult, error) {
	newBot, err := bot.New(body.Name)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
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
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	result := &BotCreateResult{
		Bot: newBot,
		Key: botKey,
	}
	return result, nil
}

func (s *BotService) GetBotById(id string) (*bot.Bot, error) {
	foundBot, err := s.store.GetBot(id)
	if err != nil {
		return nil, err
	}

	return foundBot, nil
}

func (s *BotService) GetAllBots() ([]*bot.Bot, error) {
	bots, err := s.store.GetAllBots()
	if err != nil {
		return nil, err
	}

	return bots, nil
}

func (s *BotService) EditBot(id string, body *BotEditBody) (*bot.Bot, error) {
	foundBot, err := s.GetBotById(id)
	if err != nil {
		return nil, err
	}

	if body.Name != "" {
		err := foundBot.SetName(body.Name)
		if err != nil {
			return nil, extlib.ErrorBadRequest(err.Error())
		}
	}

	if body.Description != "" {
		err := foundBot.SetDescription(body.Description)
		if err != nil {
			return nil, extlib.ErrorBadRequest(err.Error())
		}
	}

	if body.Avatar != "" {
		err := foundBot.SetAvatar(body.Avatar)
		if err != nil {
			return nil, extlib.ErrorBadRequest(err.Error())
		}
	}

	err = s.store.EditBot(foundBot)
	if err != nil {
		return nil, err
	}

	return foundBot, nil
}

func (s *BotService) DeleteBot(id string) error {
	err := s.store.DeleteBot(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *BotService) AddCommands(botId string, body *BotCommandsBody) error {
	_, err := s.GetBotById(botId)
	if err != nil {
		return err
	}

	for _, c := range body.Commands {
		newCmd, err := bot.NewCommand(botId, c.Alias, c.Description)
		if err != nil {
			return extlib.ErrorBadRequest(err.Error())
		}

		err = s.store.SaveCommand(newCmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *BotService) RemoveCommand(botId string, alias string) error {
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

func (s *BotService) GetKey(id string) (*BotKeyResult, error) {
	keyData, err := s.store.GetKey(id)
	if err != nil {
		return nil, err
	}

	return &BotKeyResult{
		Value: keyData.Assemble(),
	}, nil
}

func (s *BotService) GenerateKey(id string) (*BotKeyResult, error) {
	token, err := extlib.RandomToken(ulid.Length)
	if err != nil {
		return nil, extlib.ErrorBadRequest("bot key generation failed: " + err.Error())
	}

	botKey, err := bot.NewKey(id, token)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	err = s.store.SaveKey(botKey)
	if err != nil {
		return nil, err
	}

	return &BotKeyResult{
		Value: botKey.Assemble(),
	}, nil
}

func (s *BotService) VerifyKeyData(id, token string) error {
	kd, err := s.store.GetKey(id)
	if err != nil {
		return nil
	}

	if kd == nil || kd.Token != token {
		return extlib.ErrorForbidden("invalid bot key")
	}

	return nil
}

func (s *BotService) CreateWebhook(botId, url, secret string) error {
	webhook, err := bot.NewWebhook(botId, url, secret)
	if err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	err = s.store.SaveWebhook(webhook)
	if err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	return nil
}

func (s *BotService) GetWebhook(botId string) (*bot.Webhook, error) {
	wh, err := s.store.GetWebhook(botId)
	if err != nil {
		return nil, err
	}

	return wh, nil
}

func (s *BotService) DeleteWebhook(botId string) error {
	err := s.store.DeleteWebhook(botId)
	if err != nil {
		return err
	}

	return nil
}
