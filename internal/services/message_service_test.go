package services

import (
	"botyard/internal/entities/message"
	"botyard/internal/tools/ulid"
	"testing"
)

func TestMessageService(t *testing.T) {
	mockFileStore := &mockFileStore{}
	mockMessageStore := &mockMessageStore{}
	testFileService := NewFileService(mockFileStore)
	messageService := NewMessageService(mockMessageStore, testFileService)

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

	t.Run("get messages by chat", func(t *testing.T) {
		_, err := messageService.GetMessagesByChat(ulid.New(), 1, 1)
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("delete messages by chat", func(t *testing.T) {
		err := messageService.DeleteMessagesByChat(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

type mockMessageStore struct{}

func (mms *mockMessageStore) AddMessage(msg *message.Message) error {
	return nil
}

func (mms *mockMessageStore) GetMessagesByChat(chatId string, page, limit int) (int, []*message.Message, error) {
	return 0, []*message.Message{}, nil
}

func (mms *mockMessageStore) DeleteMessagesByChat(chatId string) error {
	return nil
}
