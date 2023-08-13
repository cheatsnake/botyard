package chat

import (
	"botyard/internal/entities/chat"
	"botyard/internal/http/bot"
	"botyard/internal/http/helpers"
	"botyard/internal/http/message"
	"botyard/internal/storage"
	"net/http"
)

type Service struct {
	bot   *bot.Service
	msg   *message.Service
	store storage.Storage
}

func NewService(s storage.Storage, bs *bot.Service, ms *message.Service) *Service {
	return &Service{
		bot:   bs,
		msg:   ms,
		store: s,
	}
}

func (s *Service) Create(userId, botId string) (*chat.Chat, error) {
	_, err := s.bot.FindById(botId)
	if err != nil {
		return nil, err
	}

	chat := chat.New(userId, botId)
	err = s.store.AddChat(chat)
	if err != nil {
		return nil, helpers.NewHttpError(http.StatusBadRequest, err.Error())
	}

	return chat, nil
}

func (s *Service) Delete(id string) error {
	err := s.store.DeleteChat(id)
	if err != nil {
		return helpers.NewHttpError(http.StatusNotFound, err.Error())
	}

	return nil
}
