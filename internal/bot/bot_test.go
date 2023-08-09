package bot

import "testing"

func TestBot(t *testing.T) {
	testName := "test"
	testCommands := []Command{
		{Alias: "start", Description: "test start"},
		{Alias: "help", Description: "test help"},
		{Alias: "ping", Description: "test pong"},
	}

	bot := New(testName)

	t.Run("check name", func(t *testing.T) {
		expect := testName
		if bot.Name != expect {
			t.Errorf("%#v got: %s, expect: %s", bot, bot.Name, expect)
		}
	})

	t.Run("check id", func(t *testing.T) {
		expect := 26
		if len(bot.Id) != expect {
			t.Errorf("%#v got: %d, expect: %d", bot, len(bot.Id), expect)
		}
	})

	t.Run("change name", func(t *testing.T) {
		expect := "Rob Pike"
		bot.SetName(expect)

		if bot.Name != expect {
			t.Errorf("%#v got: %s, expect: %s", bot, bot.Name, expect)
		}
	})

	t.Run("set description", func(t *testing.T) {
		expect := "this is test bot"
		bot.SetDescription(expect)

		if bot.Description != expect {
			t.Errorf("%#v got: %s, expect: %s", bot, bot.Description, expect)
		}
	})

	t.Run("set avatar", func(t *testing.T) {
		expect := "https://google.com/avatar.jpg"
		bot.SetAvatar(expect)

		if bot.Avatar != expect {
			t.Errorf("%#v got: %s, expect: %s", bot, bot.Avatar, expect)
		}
	})

	t.Run("set commands", func(t *testing.T) {
		for _, cmd := range testCommands {
			bot.AddCommand(cmd.Alias, cmd.Description)
		}

		expect := len(testCommands)

		if len(bot.Commands) != expect {
			t.Errorf("%#v got: %d, expect: %d", bot, len(bot.Commands), expect)
		}
	})

	t.Run("find commands", func(t *testing.T) {
		for _, expect := range testCommands {
			cmd, err := bot.Command(expect.Alias)
			if err != nil {
				t.Errorf("%#v got: %s, expect: %v", bot, err.Error(), nil)
			}

			if cmd.Alias != expect.Alias {
				t.Errorf("%#v got: %s, expect: %s", bot, cmd.Alias, expect.Alias)
			}

			if cmd.Description != expect.Description {
				t.Errorf("%#v got: %s, expect: %s", bot, cmd.Description, expect.Description)
			}
		}
	})

	t.Run("set existent command", func(t *testing.T) {
		expect := testCommands[0]
		bot.AddCommand(expect.Alias, expect.Description)

		cmd, err := bot.Command(expect.Alias)
		if err != nil {
			t.Errorf("%#v got: %s, expect: %v", bot, err.Error(), nil)
		}

		if cmd.Alias != expect.Alias {
			t.Errorf("%#v got: %s, expect: %s", bot, cmd.Alias, expect.Alias)
		}

		if cmd.Description != expect.Description {
			t.Errorf("%#v got: %s, expect: %s", bot, cmd.Description, expect.Description)
		}
	})

	t.Run("remove commands", func(t *testing.T) {
		for _, cmd := range testCommands {
			err := bot.RemoveCommand(cmd.Alias)
			if err != nil {
				t.Errorf("%#v got: %s, expect: %v", bot, err.Error(), nil)
			}
		}

		expect := 0

		if len(bot.Commands) != expect {
			t.Errorf("%#v got: %d, expect: %d", bot, len(bot.Commands), expect)
		}
	})

	t.Run("find non-existent command", func(t *testing.T) {
		testCommand := testCommands[0]
		expect := Command{}

		cmd, err := bot.Command(testCommand.Alias)
		if err.Error() != errCmdNotFound {
			t.Errorf("%#v got: %s, expect: %s", bot, err.Error(), errCmdNotFound)
		}

		if cmd.Alias != expect.Alias {
			t.Errorf("%#v got: %s, expect: %s", bot, cmd.Alias, expect.Alias)
		}

		if cmd.Description != expect.Description {
			t.Errorf("%#v got: %s, expect: %s", bot, cmd.Description, expect.Description)
		}
	})

	t.Run("remove non-existent command", func(t *testing.T) {
		err := bot.RemoveCommand(testCommands[0].Alias)

		expect := errCmdNotFound

		if err.Error() != expect {
			t.Errorf("%#v got: %s, expect: %s", bot, err.Error(), expect)
		}
	})

	t.Run("get all commands", func(t *testing.T) {
		for _, cmd := range testCommands {
			bot.AddCommand(cmd.Alias, cmd.Description)
		}

		expect := len(testCommands)
		result := len(bot.GetCommands())

		if result != expect {
			t.Errorf("%#v got: %d, expect: %d", bot, result, expect)
		}
	})
}
