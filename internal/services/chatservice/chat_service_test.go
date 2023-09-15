package chatservice

import (
	"botyard/internal/services/fileservice"
	"botyard/internal/services/messageservice"
	mock "botyard/internal/storage/_mock"
	"botyard/internal/tools/ulid"
	"testing"
)

func TestCreate(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	ms := messageservice.New(mock.MessageStore(), fs)
	cs := New(mock.ChatStore(), ms)

	t.Run("create a new chat", func(t *testing.T) {
		_, err := cs.Create(ulid.New(), ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("create a new chat with empty values", func(t *testing.T) {
		_, err := cs.Create("", "")
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})
}

func TestGetByBot(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	ms := messageservice.New(mock.MessageStore(), fs)
	cs := New(mock.ChatStore(), ms)

	t.Run("get chats by bot", func(t *testing.T) {
		_, err := cs.GetByBot(ulid.New(), ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestDelete(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	ms := messageservice.New(mock.MessageStore(), fs)
	cs := New(mock.ChatStore(), ms)

	t.Run("delete chat", func(t *testing.T) {
		err := cs.Delete(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}
