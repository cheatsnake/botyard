package messageservice

import (
	"botyard/internal/entities/message"
	"botyard/internal/services/fileservice"
	mock "botyard/internal/storage/_mock"
	"botyard/internal/tools/ulid"
	"testing"
)

func TestAddMessage(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	ms := New(mock.MessageStore(), fs)

	t.Run("add message", func(t *testing.T) {
		err := ms.AddMessage(&CreateBody{
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
		err := ms.AddMessage(&CreateBody{})
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", nil, err.Error())
		}
	})

}

func TestGetMessagesByChat(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	ms := New(mock.MessageStore(), fs)

	t.Run("get messages by chat", func(t *testing.T) {
		_, err := ms.GetMessagesByChat(ulid.New(), 1, 1)
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestDeleteMessagesByChat(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	ms := New(mock.MessageStore(), fs)

	t.Run("delete messages by chat", func(t *testing.T) {
		err := ms.DeleteMessagesByChat(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}
