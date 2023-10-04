package bot

import (
	"strings"
	"testing"

	"github.com/cheatsnake/botyard/internal/tools/ulid"
	"github.com/cheatsnake/botyard/pkg/extlib"
)

func TestNewWebhook(t *testing.T) {
	testBotId := ulid.New()
	testUrl := "https://go.dev"
	testSecret := "test"

	t.Run("valid botId, url and secret", func(t *testing.T) {
		testWebhook, err := NewWebhook(testBotId, testUrl, testSecret)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err.Error(), nil)
		}

		if testWebhook.BotId != testBotId {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, testWebhook.BotId, testBotId)
		}

		if testWebhook.Url != testUrl {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, testWebhook.Url, testUrl)
		}

		if testWebhook.Secret != testSecret {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, testWebhook.Secret, testSecret)
		}
	})

	t.Run("empty botId", func(t *testing.T) {
		testWebhook, err := NewWebhook("", testUrl, testSecret)
		if err == nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err, errBotIdIsEmpty)
		}
	})

	t.Run("empty url", func(t *testing.T) {
		testWebhook, err := NewWebhook(testBotId, "", testSecret)
		if err == nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err, errWebhookUrlIsEmpty)
		}
	})

	t.Run("too long url", func(t *testing.T) {
		tooLongUrl := testUrl + strings.Repeat("a", maxWebhookUrlLen-len(testUrl)+1)
		testWebhook, err := NewWebhook(testBotId, tooLongUrl, testSecret)
		if err == nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err, errWebhookUrlTooLong)
		}
	})

	t.Run("invalid url", func(t *testing.T) {
		for _, u := range []string{"abc", "test://url.dev", "localhost:6666", "123-456-789"} {
			testWebhook, err := NewWebhook(testBotId, u, testSecret)
			if err == nil {
				t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err, extlib.ErrInvalidUrl)
			}
		}
	})

	t.Run("too secret", func(t *testing.T) {
		tooLongSecret := strings.Repeat("a", maxWebhookSecretLen+1)
		testWebhook, err := NewWebhook(testBotId, testUrl, tooLongSecret)
		if err == nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err, errWebhookSecretTooLong)
		}
	})
}

func TestSetUrl(t *testing.T) {
	testBotId := ulid.New()
	testUrl := "https://go.dev"
	testSecret := "test"

	testWebhook, err := NewWebhook(testBotId, testUrl, testSecret)
	if err != nil {
		t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err.Error(), nil)
	}

	t.Run("set valid url", func(t *testing.T) {
		err := testWebhook.SetUrl("https://google.com")
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err.Error(), nil)
		}
	})

	t.Run("set invalid url", func(t *testing.T) {
		err := testWebhook.SetUrl("test")
		if err == nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err, "error")
		}
	})
}

func TestSetSecret(t *testing.T) {
	testBotId := ulid.New()
	testUrl := "https://go.dev"
	testSecret := "test"

	testWebhook, err := NewWebhook(testBotId, testUrl, testSecret)
	if err != nil {
		t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err.Error(), nil)
	}

	t.Run("set valid secret", func(t *testing.T) {
		err := testWebhook.SetSecret("newTest")
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err.Error(), nil)
		}
	})

	t.Run("set invalid secret", func(t *testing.T) {
		err := testWebhook.SetSecret(strings.Repeat("a", maxWebhookSecretLen+1))
		if err == nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testWebhook, err, "error")
		}
	})
}
