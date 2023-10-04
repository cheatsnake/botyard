package chat

import (
	"testing"

	"github.com/cheatsnake/botyard/internal/tools/ulid"
)

func TestChat(t *testing.T) {
	testUserId := ulid.New()
	testBotId := ulid.New()

	t.Run("check user id", func(t *testing.T) {
		testChat, err := New(testUserId, testBotId)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testChat, err.Error(), nil)
		}

		expect := true
		got := testChat.UserId == testUserId

		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testChat, got, expect)
		}
	})

	t.Run("check empty user id", func(t *testing.T) {
		expect := errUserIdIsEmpty
		testChat, err := New("", testBotId)
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testChat, nil, expect)
		}

		got := err.Error()
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testChat, got, expect)
		}
	})

	t.Run("check bot id", func(t *testing.T) {
		testChat, err := New(testUserId, testBotId)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testChat, err.Error(), nil)
		}

		expect := true
		got := testChat.BotId == testBotId

		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testChat, got, expect)
		}
	})

	t.Run("check empty bot id", func(t *testing.T) {
		expect := errBotIdIsEmpty
		testChat, err := New(testUserId, "")
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testChat, nil, expect)
		}

		got := err.Error()
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testChat, got, expect)
		}
	})
}
