package memory

import (
	"botyard/internal/chat"
	"errors"
)

func (s *Storage) AddChat(chat *chat.Chat) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.chats = append(s.chats, chat)
	return nil
}

func (s *Storage) GetChat(id string) (*chat.Chat, error) {
	for _, chat := range s.chats {
		if chat.Id == id {
			return chat, nil
		}
	}

	return nil, errors.New("chat not found")
}

func (s *Storage) FindChat(userId, botId string) (*chat.Chat, error) {
	for _, chat := range s.chats {
		if chat.UserId == userId && chat.BotId == botId {
			return chat, nil
		}
	}

	return nil, errors.New("chat not found")
}
