package mock

import "botyard/internal/entities/message"

type messageStore struct{}

func (mms *messageStore) AddMessage(msg *message.Message) error {
	return nil
}

func (mms *messageStore) GetMessagesByChat(chatId string, page, limit int) (int, []*message.Message, error) {
	return 0, []*message.Message{}, nil
}

func (mms *messageStore) DeleteMessagesByChat(chatId string) error {
	return nil
}

func MessageStore() *messageStore {
	return &messageStore{}
}
