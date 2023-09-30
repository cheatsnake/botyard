package memory

import (
	"botyard/internal/entities/chat"
	"botyard/pkg/exterr"
	"botyard/pkg/extlib"

	"golang.org/x/exp/slices"
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

	return nil, exterr.ErrorNotFound("chat not found")
}

func (s *Storage) GetChats(userId, botId string) ([]*chat.Chat, error) {
	chats := []*chat.Chat{}

	for _, chat := range s.chats {
		if chat.UserId == userId && chat.BotId == botId {
			chats = append(chats, chat)
		}
	}

	return chats, nil
}

func (s *Storage) DeleteChat(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delIndex := slices.IndexFunc(s.chats, func(c *chat.Chat) bool {
		return c.Id == id
	})

	if delIndex == -1 {
		return exterr.ErrorNotFound("chat not found")
	}

	s.chats = extlib.SliceRemoveElement(s.chats, delIndex)
	return nil
}

func (s *Storage) AddMessage(msg *chat.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.messages = append(s.messages, msg)

	return nil
}

func (s *Storage) GetMessagesByChat(chatId string, page, limit int) (int, []*chat.Message, error) {
	if page <= 0 || limit <= 0 {
		return 0, nil, nil
	}

	chatMsgs := extlib.SliceFilter(extlib.SliceReverse(s.messages), 10, func(m *chat.Message) bool {
		return m.ChatId == chatId
	})

	if page == 1 && limit >= len(chatMsgs) {
		msgs := make([]chat.Message, len(chatMsgs))
		for i, msg := range chatMsgs {
			msgs[i] = *msg
		}

		return len(chatMsgs), chatMsgs, nil
	}

	return len(chatMsgs), extlib.SliceReverse(extlib.SlicePaginate(chatMsgs, page, limit)), nil
}

func (s *Storage) DeleteMessagesByChat(chatId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	filtered := extlib.SliceFilter(s.messages, 0, func(m *chat.Message) bool {
		return m.ChatId != chatId
	})

	s.messages = filtered
	return nil
}
