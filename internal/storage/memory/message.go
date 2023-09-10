package memory

import (
	"botyard/internal/entities/message"
	"botyard/pkg/extlib"
)

func (s *Storage) AddMessage(msg *message.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.messages = append(s.messages, msg)

	return nil
}

func (s *Storage) GetMessagesByChat(chatId string, page, limit int) (int, []*message.Message, error) {
	if page <= 0 || limit <= 0 {
		return 0, nil, nil
	}

	chatMsgs := extlib.SliceFilter(extlib.SliceReverse(s.messages), 10, func(m *message.Message) bool {
		return m.ChatId == chatId
	})

	if page == 1 && limit >= len(chatMsgs) {
		msgs := make([]message.Message, len(chatMsgs))
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

	filtered := extlib.SliceFilter(s.messages, 0, func(m *message.Message) bool {
		return m.ChatId != chatId
	})

	s.messages = filtered
	return nil
}
