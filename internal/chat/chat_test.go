package chat

import (
	"botyard/internal/tools/ulid"
	"botyard/internal/user"
	"testing"

	"golang.org/x/exp/slices"
)

func TestChat(t *testing.T) {
	testUser, _ := user.New("user")
	testStore := new(mockStore)

	t.Run("check chat id", func(t *testing.T) {
		testChat := New([]string{testUser.Id}, testStore)
		expect := ulid.Length
		got := len(testChat.Id)

		if got != expect {
			t.Errorf("%#v got: %d, expect: %d", testChat, got, expect)
		}
	})

	t.Run("check member ids", func(t *testing.T) {
		testChat := New([]string{testUser.Id}, testStore)
		expect := true
		got := slices.Contains(testChat.MemberIds, testUser.Id)

		if got != expect {
			t.Errorf("%#v got: %v, expect: %v", testChat, got, expect)
		}
	})

	t.Run("send message", func(t *testing.T) {
		testChat := New([]string{testUser.Id}, testStore)
		got := testChat.SendMessage(testUser.Id, "hello", nil)

		if got != nil {
			t.Errorf("%#v got: %v, expect: %v", testChat, got, nil)
		}
	})

	t.Run("send message by not allowed member", func(t *testing.T) {
		testChat := New([]string{testUser.Id}, testStore)
		expect := errSenderNotMember
		got := testChat.SendMessage(ulid.New(), "hello", nil).Error()

		if got != expect {
			t.Errorf("%#v got: %v, expect: %v", testChat, got, expect)
		}
	})

	t.Run("get messages", func(t *testing.T) {
		testChat := New([]string{testUser.Id}, testStore)
		_, got := testChat.GetMessages(1, 1)

		if got != nil {
			t.Errorf("%#v got: %v, expect: %v", testChat, got, nil)
		}
	})

	t.Run("clear chat", func(t *testing.T) {
		testChat := New([]string{testUser.Id}, testStore)
		got := testChat.Clear()

		if got != nil {
			t.Errorf("%#v got: %v, expect: %v", testChat, got, nil)
		}
	})
}

type mockStore struct{}

func (ms *mockStore) AddMessage(msg *Message) error {
	return nil
}

func (ms *mockStore) GetMessagesByChat(chatId string, page, limit int) (int, []*Message, error) {
	return 0, []*Message{newMessage("", "", "", nil)}, nil
}

func (ms *mockStore) DeleteMessagesByChat(chatId string) error {
	return nil
}

func (ms *mockStore) AddFile(file *File) error {
	return nil
}

func (ms *mockStore) GetFile(id string) (*File, error) {
	return NewFile("", ""), nil
}

func (ms *mockStore) GetFiles(ids []string) ([]*File, error) {
	return []*File{NewFile("", "")}, nil
}

func (ms *mockStore) DeleteFile(id string) error {
	return nil
}
