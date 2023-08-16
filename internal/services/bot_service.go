package services

import (
	"botyard/internal/entities/bot"
	"botyard/internal/storage"
	"botyard/pkg/extlib"
)

type BotService struct {
	store storage.BotStore
}

type BotCreateBody struct {
	bot.Bot
	Id struct{} `json:"-"`
}

type BotEditBody struct {
	bot.Bot
	Commands struct{} `json:"-"`
	Id       struct{} `json:"-"`
}

type BotCommandsBody struct {
	Commands []bot.Command
}

type BotCommandBody struct {
	Alias string
}

func NewBotService(s storage.BotStore) *BotService {
	return &BotService{
		store: s,
	}
}

func (s *BotService) Create(body *BotCreateBody) (*bot.Bot, error) {
	newBot := bot.New(body.Name)

	if body.Description != "" {
		newBot.SetDescription(body.Description)
	}

	if body.Avatar != "" {
		newBot.SetAvatar(body.Avatar)
	}

	for _, cmd := range body.Commands {
		newBot.AddCommand(cmd.Alias, cmd.Description)
	}

	err := s.store.AddBot(newBot)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	return newBot, nil
}

func (s *BotService) FindById(id string) (*bot.Bot, error) {
	foundBot, err := s.store.GetBot(id)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	return foundBot, nil
}

func (s *BotService) Edit(id string, body *BotEditBody) (*bot.Bot, error) {
	foundBot, err := s.FindById(id)
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
	foundBot, err := s.FindById(id)
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

func (s *BotService) RemoveCommand(id string, body *BotCommandBody) error {
	newBot, err := s.FindById(id)
	if err != nil {
		return err
	}

	err = newBot.RemoveCommand(body.Alias)
	if err != nil {
		return extlib.ErrorNotFound(err.Error())
	}

	err = s.store.EditBot(newBot)
	if err != nil {
		return extlib.ErrorNotFound(err.Error())
	}

	return nil
}

func (s *BotService) GetCommands(id string) ([]bot.Command, error) {
	foundBot, err := s.FindById(id)
	if err != nil {
		return nil, err
	}

	return foundBot.GetCommands(), nil
}
