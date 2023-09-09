package services

import (
	"botyard/internal/entities/bot"
	mock "botyard/internal/storage/_mock"
	"botyard/internal/tools/ulid"
	"strings"
	"testing"
)

func TestBotServiceCreateBot(t *testing.T) {
	botService := NewBotService(mock.BotStore())

	t.Run("add a new bot", func(t *testing.T) {
		res, err := botService.CreateBot(&BotCreateBody{
			Bot: bot.Bot{
				Name: "test",
			},
		})
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}

		if res.Key.Value == "" {
			t.Errorf("got: %v,\nexpect: %v\n", res.Key, "key")
		}
	})

	t.Run("add a new bot without name", func(t *testing.T) {
		_, err := botService.CreateBot(&BotCreateBody{
			Bot: bot.Bot{},
		})
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})
}

func TestBotServiceGetBotById(t *testing.T) {
	botService := NewBotService(mock.BotStore())

	t.Run("get bot by id", func(t *testing.T) {
		_, err := botService.GetBotById(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestBotServiceGetAllBots(t *testing.T) {
	botService := NewBotService(mock.BotStore())

	t.Run("get all bots", func(t *testing.T) {
		_, err := botService.GetAllBots()
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestBotServiceEditBot(t *testing.T) {
	botService := NewBotService(mock.BotStore())

	t.Run("edit bot", func(t *testing.T) {
		_, err := botService.EditBot(ulid.New(), &BotEditBody{})
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}
func TestBotServiceAddCommands(t *testing.T) {
	botService := NewBotService(mock.BotStore())

	t.Run("add bot commands", func(t *testing.T) {
		err := botService.AddCommands(ulid.New(), &BotCommandsBody{})
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}

func TestBotServiceRemoveCommand(t *testing.T) {
	botService := NewBotService(mock.BotStore())

	t.Run("remove bot command", func(t *testing.T) {
		err := botService.RemoveCommand(ulid.New(), "")
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})
}

func TestBotServiceDeleteBot(t *testing.T) {}

func TestBotServiceCreateWebhook(t *testing.T) {

}
func TestBotServiceGetWebhook(t *testing.T) {}

func TestBotServiceDeleteWebhook(t *testing.T) {}

func TestBotServiceGenerateKey(t *testing.T) {
	botService := NewBotService(mock.BotStore())

	t.Run("generate bot key", func(t *testing.T) {
		botId := ulid.New()
		botKey, err := botService.GenerateKey(botId)
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}

		if !strings.Contains(botKey.Value, botId) {
			t.Errorf("got: %v,\nexpect: %v\n", botKey, botId)
		}
	})
}
func TestBotServiceGetKey(t *testing.T) {
	botService := NewBotService(mock.BotStore())

	t.Run("get bot key", func(t *testing.T) {
		botId := ulid.New()
		_, err := botService.GetKey(botId)
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

}
func TestBotServiceVerifyKeyData(t *testing.T) {
	botService := NewBotService(mock.BotStore())

	t.Run("verify bot key", func(t *testing.T) {
		botId := ulid.New()
		err := botService.VerifyKeyData(botId, "test")
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("verify invalid bot key", func(t *testing.T) {
		botId := ulid.New()
		err := botService.VerifyKeyData(botId, "fail")
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})
}
