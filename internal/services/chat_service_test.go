package services

import (
	"botyard/internal/entities/chat"
	"botyard/internal/tools/ulid"
	"testing"
)

func TestChatService(t *testing.T) {
	mockFileStore := &mockFileStore{}
	mockMessageStore := &mockMessageStore{}
	mockChatStore := &mockChatStore{}
	fileService := NewFileService(mockFileStore)
	messageService := NewMessageService(mockMessageStore, fileService)
	chatService := NewChatService(mockChatStore, messageService)

	t.Run("create a new chat", func(t *testing.T) {
		_, err := chatService.Create(ulid.New(), ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("create a new chat with empty values", func(t *testing.T) {
		_, err := chatService.Create("", "")
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})

	t.Run("get chats by bot", func(t *testing.T) {
		_, err := chatService.GetByBot(ulid.New(), ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("delete chat", func(t *testing.T) {
		err := chatService.Delete(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

type mockChatStore struct{}

func (mcs *mockChatStore) AddChat(chat *chat.Chat) error {
	return nil
}

func (mcs *mockChatStore) GetChat(id string) (*chat.Chat, error) {
	return &chat.Chat{}, nil
}

func (mcs *mockChatStore) GetChats(userId, botId string) ([]*chat.Chat, error) {
	return []*chat.Chat{}, nil
}

func (mcs *mockChatStore) DeleteChat(id string) error {
	return nil
}
