package services

import (
	"botyard/internal/entities/bot"
	mock "botyard/internal/storage/_mock"
	"botyard/internal/tools/ulid"
	"strings"
	"testing"
)

func TestBotService(t *testing.T) {
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

		if res.Key == "" {
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

	t.Run("get bot by id", func(t *testing.T) {
		_, err := botService.GetBotById(ulid.New())
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("get all bots", func(t *testing.T) {
		_, err := botService.GetAllBots()
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("edit bot", func(t *testing.T) {
		_, err := botService.EditBot(ulid.New(), &BotEditBody{})
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("add bot commands", func(t *testing.T) {
		err := botService.AddCommands(ulid.New(), &BotCommandsBody{})
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("remove bot command", func(t *testing.T) {
		err := botService.RemoveCommand(ulid.New(), "")
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})

	t.Run("generate bot key", func(t *testing.T) {
		botId := ulid.New()
		bkr, err := botService.GenerateAuthKey(botId)
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}

		if !strings.Contains(bkr.Key, botId) {
			t.Errorf("got: %v,\nexpect: %v\n", bkr, botId)
		}
	})

	t.Run("verify bot key", func(t *testing.T) {
		botId := ulid.New()
		err := botService.VerifyAuthKey(botId, "test")
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})

	t.Run("verify invalid bot key", func(t *testing.T) {
		botId := ulid.New()
		err := botService.VerifyAuthKey(botId, "fail")
		if err == nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, "error")
		}
	})
}
