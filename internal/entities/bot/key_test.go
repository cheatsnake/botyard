package bot

import (
	"testing"

	"github.com/cheatsnake/botyard/internal/tools/ulid"
)

func TestNewKey(t *testing.T) {
	testBotId := ulid.New()
	testToken := "test"

	t.Run("valid botId and token", func(t *testing.T) {
		key, err := NewKey(testBotId, testToken)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", key, err.Error(), nil)
		}

		if key.BotId != testBotId {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", key, key.BotId, testBotId)
		}

		if key.Token != testToken {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", key, key.Token, testToken)
		}
	})

	t.Run("empty botId", func(t *testing.T) {
		key, err := NewKey("", testToken)
		if err == nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", key, err, errBotIdIsEmpty)
		}
	})

	t.Run("empty token", func(t *testing.T) {
		key, err := NewKey(testBotId, "")
		if err == nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", key, err, errKeyTokenIsEmpty)
		}
	})
}

func TestSetToken(t *testing.T) {
	t.Run("set valild token", func(t *testing.T) {
		testKey, err := NewKey(ulid.New(), "test")
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testKey, err.Error(), nil)
		}

		testToken := "newTest"
		err = testKey.SetToken(testToken)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testKey, err.Error(), nil)
		}

		if testKey.Token != testToken {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testKey, testKey.Token, testToken)
		}
	})

	t.Run("set empty token", func(t *testing.T) {
		testKey, err := NewKey(ulid.New(), "test")
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testKey, err.Error(), nil)
		}

		testToken := ""
		err = testKey.SetToken(testToken)
		if err == nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testKey, err, errKeyTokenIsEmpty)
		}
	})
}

func TestAssemble(t *testing.T) {
	t.Run("assemble key", func(t *testing.T) {
		testBotId := ulid.New()
		testToken := "test"
		testKey, err := NewKey(testBotId, testToken)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testKey, err.Error(), nil)
		}

		expect := testBotId + ":" + testToken
		got := testKey.Assemble()

		if got != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testKey, got, expect)
		}
	})
}
