package bot

import (
	"botyard/internal/entities/bot"
	"botyard/internal/http/helpers"
	"botyard/internal/storage"
	"net/http"
)

type Service struct {
	store storage.Storage
}

func NewService(s storage.Storage) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) Create(body *createBody) (*bot.Bot, error) {
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
		return nil, helpers.NewHttpError(http.StatusBadRequest, err.Error())
	}

	return newBot, nil
}

func (s *Service) FindById(id string) (*bot.Bot, error) {
	foundBot, err := s.store.GetBot(id)
	if err != nil {
		return nil, helpers.NewHttpError(http.StatusNotFound, err.Error())
	}

	return foundBot, nil
}

func (s *Service) Edit(id string, body *editBody) (*bot.Bot, error) {
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
		return nil, helpers.NewHttpError(http.StatusBadRequest, err.Error())
	}

	return foundBot, nil
}

func (s *Service) AddCommands(id string, body *commandsBody) error {
	foundBot, err := s.FindById(id)
	if err != nil {
		return err
	}

	for _, cmd := range body.Commands {
		foundBot.AddCommand(cmd.Alias, cmd.Description)
	}

	err = s.store.EditBot(foundBot)
	if err != nil {
		return helpers.NewHttpError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (s *Service) RemoveCommand(id string, body *commandBody) error {
	newBot, err := s.FindById(id)
	if err != nil {
		return err
	}

	err = newBot.RemoveCommand(body.Alias)
	if err != nil {
		return helpers.NewHttpError(http.StatusNotFound, err.Error())
	}

	err = s.store.EditBot(newBot)
	if err != nil {
		return helpers.NewHttpError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (s *Service) GetCommands(id string) ([]bot.Command, error) {
	foundBot, err := s.FindById(id)
	if err != nil {
		return nil, err
	}

	return foundBot.GetCommands(), nil
}
