package memory

import (
	"botyard/internal/entities/chat"
	"botyard/pkg/extlib"
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

func (s *Storage) GetChats(userId, botId string) ([]*chat.Chat, error) {
	var chats []*chat.Chat

	for _, chat := range s.chats {
		if chat.UserId == userId && chat.BotId == botId {
			chats = append(chats, chat)
		}
	}

	return chats, nil
}

func (s *Storage) DeleteChat(id string) error {
	for i, chat := range s.chats {
		if chat.Id == id {
			s.chats = extlib.SliceRemoveElement(s.chats, i)
			return nil
		}
	}

	return errors.New("chat not found")
}
