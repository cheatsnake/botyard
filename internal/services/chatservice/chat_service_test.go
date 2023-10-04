package chatservice

import (
	"testing"

	"github.com/cheatsnake/botyard/internal/entities/chat"
	"github.com/cheatsnake/botyard/internal/services/fileservice"
	mock "github.com/cheatsnake/botyard/internal/storage/_mock"
	"github.com/cheatsnake/botyard/internal/tools/ulid"
)

func TestCreateChat(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	cs := New(mock.ChatStore(), fs)

	t.Run("create a new chat", func(t *testing.T) {
		_, err := cs.CreateChat(ulid.New(), ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("create a new chat with empty values", func(t *testing.T) {
		_, err := cs.CreateChat("", "")
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})
}

func TestGetChats(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	cs := New(mock.ChatStore(), fs)

	t.Run("get chats by bot", func(t *testing.T) {
		_, err := cs.GetChats(ulid.New(), ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestDeleteChat(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	cs := New(mock.ChatStore(), fs)

	t.Run("delete chat", func(t *testing.T) {
		err := cs.DeleteChat(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestCheckChatAccess(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	cs := New(mock.ChatStore(), fs)

	t.Run("check chat access", func(t *testing.T) {
		_, err := cs.CheckChatAccess(ulid.New(), "", "")
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})
}

func TestAddMessage(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	cs := New(mock.ChatStore(), fs)

	t.Run("add message", func(t *testing.T) {
		_, err := cs.AddMessage(&CreateBody{
			Message: chat.Message{
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
		_, err := cs.AddMessage(&CreateBody{})
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", nil, err.Error())
		}
	})

}

func TestGetMessagesByChat(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	cs := New(mock.ChatStore(), fs)

	t.Run("get messages by chat", func(t *testing.T) {
		_, err := cs.GetMessagesByChat(ulid.New(), "", 1, 1, 0)
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestDeleteMessagesByChat(t *testing.T) {
	fs := fileservice.New(mock.FileStore())
	cs := New(mock.ChatStore(), fs)

	t.Run("delete messages by chat", func(t *testing.T) {
		err := cs.DeleteMessagesByChat(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}
