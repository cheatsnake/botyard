package memory

import (
	"botyard/internal/chat"
	"botyard/pkg/extlib"
)

func (s *Storage) AddMessage(msg *chat.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.messages = append(s.messages, msg)

	return nil
}

func (s *Storage) GetMessagesByChat(chatId string, page, limit int) (int, []*chat.Message, error) {
	if page == 0 || limit == 0 {
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

	return len(chatMsgs), extlib.SlicePaginate(chatMsgs, page, limit), nil
}

func (s *Storage) DeleteMessagesByChat(chatId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	filtered := extlib.SliceFilter(s.messages, len(s.messages), func(m *chat.Message) bool {
		return m.ChatId != chatId
	})

	s.messages = filtered
	return nil
}
