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

type BotKeyResult struct {
	Key string `json:"key"`
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

func (s *BotService) CreateBot(body *BotCreateBody) (*BotKeyResult, error) {
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

	botKeyRes, err := s.GenerateAuthKey(newBot.Id)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	return botKeyRes, nil
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

func (s *BotService) GenerateAuthKey(id string) (*BotKeyResult, error) {
	key, err := extlib.RandomToken(ulid.Length)
	if err != nil {
		return nil, extlib.ErrorBadRequest("bot key generation failed: " + err.Error())
	}

	err = s.store.SaveAuthKeyData(&bot.AuthKeyData{
		BotId: id,
		Key:   key,
	})

	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	botKey := id + ":" + key
	return &BotKeyResult{
		Key: botKey,
	}, nil
}

func (s *BotService) VerifyAuthKey(id, key string) error {
	bkd, err := s.store.GetAuthKeyData(id)
	if err != nil {
		return nil
	}

	if bkd == nil || bkd.Key != key {
		return extlib.ErrorForbidden("invalid bot key")
	}

	return nil
}
