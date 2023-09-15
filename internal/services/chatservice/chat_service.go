package chatservice

import (
	"botyard/internal/entities/chat"
	"botyard/internal/services/messageservice"
	"botyard/internal/storage"
	"botyard/pkg/exterr"
)

type Service struct {
	msg   *messageservice.Service
	store storage.ChatStore
}

func New(s storage.ChatStore, ms *messageservice.Service) *Service {
	return &Service{
		msg:   ms,
		store: s,
	}
}

func (s *Service) Create(userId string, botId string) (*chat.Chat, error) {
	chat, err := chat.New(userId, botId)
	if err != nil {
		return nil, exterr.ErrorBadRequest(err.Error())
	}

	err = s.store.AddChat(chat)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (s *Service) GetByBot(userId string, botId string) ([]*chat.Chat, error) {
	chats, err := s.store.GetChats(userId, botId)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (s *Service) Delete(id string) error {
	err := s.store.DeleteChat(id)
	if err != nil {
		return err
	}

	return nil
}
