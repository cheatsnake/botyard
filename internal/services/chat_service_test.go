package services

import (
	mock "botyard/internal/storage/_mock"
	"botyard/internal/tools/ulid"
	"testing"
)

func TestChatService(t *testing.T) {
	fileService := NewFileService(mock.FileStore())
	messageService := NewMessageService(mock.MessageStore(), fileService)
	chatService := NewChatService(mock.ChatStore(), messageService)

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
