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
	Commands struct{} `json:"-"`
	Id       struct{} `json:"-"`
}

type BotCommandsBody struct {
	Commands []bot.Command `json:"commands"`
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

	for _, cmd := range body.Commands {
		newBot.AddCommand(cmd.Alias, cmd.Description)
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

func (s *BotService) AddCommands(id string, body *BotCommandsBody) error {
	foundBot, err := s.GetBotById(id)
	if err != nil {
		return err
	}

	for _, cmd := range body.Commands {
		foundBot.AddCommand(cmd.Alias, cmd.Description)
	}

	err = s.store.EditBot(foundBot)
	if err != nil {
		return err
	}

	return nil
}

func (s *BotService) RemoveCommand(id string, alias string) error {
	newBot, err := s.GetBotById(id)
	if err != nil {
		return err
	}

	err = newBot.RemoveCommand(alias)
	if err != nil {
		return extlib.ErrorNotFound(err.Error())
	}

	err = s.store.EditBot(newBot)
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
