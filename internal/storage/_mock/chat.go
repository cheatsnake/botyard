package mock

import "botyard/internal/entities/chat"

type chatStore struct{}

func (mcs *chatStore) AddChat(chat *chat.Chat) error {
	return nil
}

func (mcs *chatStore) GetChat(id string) (*chat.Chat, error) {
	return &chat.Chat{}, nil
}

func (mcs *chatStore) GetChats(userId, botId string) ([]*chat.Chat, error) {
	return []*chat.Chat{}, nil
}

func (mcs *chatStore) DeleteChat(id string) error {
	return nil
}

func ChatStore() *chatStore {
	return &chatStore{}
}
