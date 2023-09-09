package services

import (
	"botyard/internal/entities/message"
	mock "botyard/internal/storage/_mock"
	"botyard/internal/tools/ulid"
	"testing"
)

func TestMessageServiceAddMessage(t *testing.T) {
	testFileService := NewFileService(mock.FileStore())
	messageService := NewMessageService(mock.MessageStore(), testFileService)

	t.Run("add message", func(t *testing.T) {
		err := messageService.AddMessage(&CreateMessageBody{
			Message: message.Message{
				ChatId:   ulid.New(),
				SenderId: ulid.New(),
				Body:     "test",
			},
		})
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("add empty message", func(t *testing.T) {
		err := messageService.AddMessage(&CreateMessageBody{})
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", nil, err.Error())
		}
	})

}

func TestMessageServiceGetMessagesByChat(t *testing.T) {
	testFileService := NewFileService(mock.FileStore())
	messageService := NewMessageService(mock.MessageStore(), testFileService)

	t.Run("get messages by chat", func(t *testing.T) {
		_, err := messageService.GetMessagesByChat(ulid.New(), 1, 1)
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestMessageServiceDeleteMessagesByChat(t *testing.T) {
	testFileService := NewFileService(mock.FileStore())
	messageService := NewMessageService(mock.MessageStore(), testFileService)

	t.Run("delete messages by chat", func(t *testing.T) {
		err := messageService.DeleteMessagesByChat(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}
