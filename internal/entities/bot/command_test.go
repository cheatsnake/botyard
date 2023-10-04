package bot

import (
	"strings"
	"testing"

	"github.com/cheatsnake/botyard/internal/tools/ulid"
)

func TestNewCommand(t *testing.T) {
	botId := ulid.New()
	alias := "test"
	descr := "hello world"

	t.Run("valid input", func(t *testing.T) {
		got, err := NewCommand(botId, alias, descr)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", got, err.Error(), nil)
		}

		if got.BotId != botId {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", got, got.BotId, botId)
		}

		if got.Alias != alias {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", got, got.Alias, alias)
		}

		if got.Description != descr {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", got, got.Description, descr)
		}
	})

	t.Run("empty botId", func(t *testing.T) {
		emptyBotId := ""
		got, err := NewCommand(emptyBotId, alias, descr)
		if err.Error() != errBotIdIsEmpty {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", got, err.Error(), errBotIdIsEmpty)
		}
	})

	t.Run("empty alias", func(t *testing.T) {
		emptyAlias := ""
		got, err := NewCommand(botId, emptyAlias, descr)
		if err.Error() != errCmdAliasIsEmpty {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", got, err.Error(), errCmdAliasIsEmpty)
		}
	})

	t.Run("too long alias", func(t *testing.T) {
		longAliasses := []string{strings.Repeat("x", maxCmdAliasLen+1), strings.Repeat("x", maxCmdAliasLen*2)}

		for _, longAlias := range longAliasses {
			got, err := NewCommand(botId, longAlias, descr)
			if err.Error() != errCmdAliasTooLong {
				t.Errorf("%#v\ngot: %s,\nexpect: %v\n", got, err.Error(), errCmdAliasTooLong)
			}
		}
	})

	t.Run("alias with not allowed symbols", func(t *testing.T) {
		invalidAliasses := []string{"$udo", "ST@RT", "!@#$%^&*()_"}

		for _, invalidAlias := range invalidAliasses {
			got, err := NewCommand(botId, invalidAlias, descr)
			if err.Error() != errCmdAliasSymbols {
				t.Errorf("%#v\ngot: %s,\nexpect: %v\n", got, err.Error(), errCmdAliasSymbols)
			}
		}
	})

	t.Run("too long description", func(t *testing.T) {
		tooLongDescr := strings.Repeat("x", maxCmdDescrLen+1)
		got, err := NewCommand(botId, alias, tooLongDescr)
		if err.Error() != errCmdDescrTooLong {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", got, err.Error(), errCmdDescrTooLong)
		}
	})
}
