package botservice

import (
	"strings"
	"testing"

	"github.com/cheatsnake/botyard/internal/entities/bot"
	mock "github.com/cheatsnake/botyard/internal/storage/_mock"
	"github.com/cheatsnake/botyard/internal/tools/ulid"
)

func TestCreateBot(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("add a new bot", func(t *testing.T) {
		res, err := bs.CreateBot(&CreateBody{
			Bot: bot.Bot{
				Name: "test",
			},
		})
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}

		if res.Key == "" {
			t.Errorf("got: %v,\nexpect: %v\n", res.Key, "key")
		}
	})

	t.Run("add a new bot without name", func(t *testing.T) {
		_, err := bs.CreateBot(&CreateBody{
			Bot: bot.Bot{},
		})
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})
}

func TestGetBotById(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("get bot by id", func(t *testing.T) {
		_, err := bs.GetBotById(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestGetAllBots(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("get all bots", func(t *testing.T) {
		_, err := bs.GetAllBots()
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestEditBot(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("edit bot", func(t *testing.T) {
		_, err := bs.EditBot(ulid.New(), &EditBody{})
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestAddCommands(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("add bot commands", func(t *testing.T) {
		err := bs.AddCommands(ulid.New(), []PreparedCommand{})
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestRemoveCommand(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("remove bot command", func(t *testing.T) {
		err := bs.RemoveCommand(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestDeleteBot(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("delete bot using valid id", func(t *testing.T) {
		err := bs.DeleteBot(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestCreateWebhook(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("create webhook", func(t *testing.T) {
		_, err := bs.SaveWebhook(ulid.New(), &WebhookBody{
			Webhook: bot.Webhook{
				Url:    "https://go.dev",
				Secret: "",
			},
		})
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("create webhook with invalid input", func(t *testing.T) {
		_, err := bs.SaveWebhook(ulid.New(), &WebhookBody{
			Webhook: bot.Webhook{
				Url:    "bad url",
				Secret: "",
			},
		})
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})
}
func TestGetWebhook(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("get webhook", func(t *testing.T) {
		_, err := bs.GetWebhook(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestDeleteWebhook(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("delete webhook", func(t *testing.T) {
		err := bs.DeleteWebhook(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestGenerateKey(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("generate bot key", func(t *testing.T) {
		botId := ulid.New()
		botKey, err := bs.GenerateKey(botId)
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}

		if !strings.Contains(botKey, botId) {
			t.Errorf("got: %v,\nexpect: %v\n", botKey, botId)
		}
	})

	t.Run("generate bot key with invalid botId", func(t *testing.T) {
		_, err := bs.GenerateKey("")
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})

}

func TestGetKey(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("get bot key", func(t *testing.T) {
		botId := ulid.New()
		_, err := bs.GetKey(botId)
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

}
func TestVerifyKeyData(t *testing.T) {
	bs := New(mock.BotStore())

	t.Run("verify bot key", func(t *testing.T) {
		botId := ulid.New()
		err := bs.VerifyKeyData(botId, "test")
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("verify invalid bot key", func(t *testing.T) {
		botId := ulid.New()
		err := bs.VerifyKeyData(botId, "fail")
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})
}
