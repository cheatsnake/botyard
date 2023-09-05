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
		return nil, extlib.ErrorBadRequest(err.Error())
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
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	return foundBot, nil
}

func (s *BotService) GetAllBots() ([]*bot.Bot, error) {
	bots, err := s.store.GetAllBots()
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	return bots, nil
}

func (s *BotService) EditBot(id string, body *BotEditBody) (*bot.Bot, error) {
	foundBot, err := s.GetBotById(id)
	if err != nil {
		return nil, err
	}

	if body.Name != "" {
		foundBot.SetName(body.Name)
	}

	if body.Description != "" {
		foundBot.SetDescription(body.Description)
	}

	if body.Avatar != "" {
		foundBot.SetAvatar(body.Avatar)
	}

	err = s.store.EditBot(foundBot)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	return foundBot, nil
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
		return extlib.ErrorBadRequest(err.Error())
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
		return extlib.ErrorNotFound(err.Error())
	}

	return nil
}

func (s *BotService) GetKey(id string) (*BotKeyResult, error) {
	keyData, err := s.store.GetKeyData(id)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	keyValue := keyData.BotId + ":" + keyData.Token
	return &BotKeyResult{
		Value: keyValue,
	}, nil
}

func (s *BotService) GenerateKey(id string) (*BotKeyResult, error) {
	token, err := extlib.RandomToken(ulid.Length)
	if err != nil {
		return nil, extlib.ErrorBadRequest("bot key generation failed: " + err.Error())
	}

	err = s.store.SaveKeyData(&bot.KeyData{
		BotId: id,
		Token: token,
	})

	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	key := id + ":" + token
	return &BotKeyResult{
		Value: key,
	}, nil
}

func (s *BotService) VerifyKeyData(id, token string) error {
	kd, err := s.store.GetKeyData(id)
	if err != nil {
		return nil
	}

	if kd == nil || kd.Token != token {
		return extlib.ErrorForbidden("invalid bot key")
	}

	return nil
}
