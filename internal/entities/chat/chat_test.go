package chat

import (
	"botyard/internal/entities/bot"
	"botyard/internal/entities/user"
	"botyard/internal/tools/ulid"
	"testing"
)

func TestChat(t *testing.T) {
	testUser, _ := user.New("user")
	testBot := bot.New("bot")

	t.Run("check chat id", func(t *testing.T) {
		testChat := New(testUser.Id, testBot.Id)
		expect := ulid.Length
		got := len(testChat.Id)

		if got != expect {
			t.Errorf("%#v\ngot: %d,\nexpect: %d\n", testChat, got, expect)
		}
	})

	t.Run("check user id", func(t *testing.T) {
		testChat := New(testUser.Id, testBot.Id)
		expect := true
		got := testChat.UserId == testUser.Id

		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testChat, got, expect)
		}
	})

	t.Run("check bot id", func(t *testing.T) {
		testChat := New(testUser.Id, testBot.Id)
		expect := true
		got := testChat.BotId == testBot.Id

		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testChat, got, expect)
		}
	})
}
