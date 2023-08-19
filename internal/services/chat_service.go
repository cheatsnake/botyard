package services

import (
	"botyard/internal/entities/chat"
	"botyard/internal/storage"
	"botyard/pkg/extlib"
)

type ChatService struct {
	msg   *MessageService
	store storage.ChatStore
}

func NewChatService(s storage.ChatStore, ms *MessageService) *ChatService {
	return &ChatService{
		msg:   ms,
		store: s,
	}
}

func (s *ChatService) Create(userId string, botId string) (*chat.Chat, error) {
	chat := chat.New(userId, botId)
	err := s.store.AddChat(chat)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	return chat, nil
}

func (s *ChatService) GetByBot(userId string, botId string) ([]*chat.Chat, error) {
	chats, err := s.store.GetChats(userId, botId)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	return chats, nil
}

func (s *ChatService) Delete(id string) error {
	err := s.store.DeleteChat(id)
	if err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	return nil
}
