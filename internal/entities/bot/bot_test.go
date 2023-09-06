package bot

import (
	"botyard/internal/tools/ulid"
	"botyard/pkg/extlib"
	"strings"
	"testing"
)

func TestNewBot(t *testing.T) {
	testName := "test"
	testCommands := []Command{
		{Alias: "start", Description: "test start"},
		{Alias: "help", Description: "test help"},
		{Alias: "ping", Description: "test pong"},
	}

	t.Run("check name", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		expect := testName
		if testBot.Name != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, testBot.Name, expect)
		}
	})

	t.Run("check empty name", func(t *testing.T) {
		expect := errNameIsEmpty
		testBot, err := New("")
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, expect)
		}

		got := err.Error()
		if got != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, got, expect)
		}
	})

	t.Run("check too long names", func(t *testing.T) {
		expect := errNameTooLong
		testNames := []string{strings.Repeat("a", maxNameLen+1), strings.Repeat("a", maxNameLen*2)}
		for _, tn := range testNames {
			testBot, err := New(tn)
			if err == nil {
				t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, expect)
			}

			got := err.Error()
			if got != expect {
				t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, got, expect)
			}
		}
	})

	t.Run("check invalid names", func(t *testing.T) {
		expect := errNameSymbols
		testNames := []string{"rob_pike", "ma$ter", ":D"}
		for _, tn := range testNames {
			testBot, err := New(tn)
			if err == nil {
				t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, expect)
			}

			got := err.Error()
			if got != expect {
				t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, got, expect)
			}
		}
	})

	t.Run("check id", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		got := len(testBot.Id)
		if got == 0 {
			t.Errorf("%#v\ngot: %d,\nexpect: %d\n", testBot, got, ulid.Length)
		}
	})

	t.Run("change name", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		expect := "Rob Pike"
		err = testBot.SetName(expect)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		if testBot.Name != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, testBot.Name, expect)
		}

		// Change to empty name
		err = testBot.SetName("")
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, errNameIsEmpty)
		}
		if err.Error() != errNameIsEmpty {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, err.Error(), errNameIsEmpty)
		}

		// Change to too long name
		err = testBot.SetName(strings.Repeat("a", maxAvatarLen+1))
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, errNameTooLong)
		}
		if err.Error() != errNameTooLong {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, err.Error(), errNameTooLong)
		}

		// Change to invalid name
		err = testBot.SetName("el0n_mu$k")
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, errNameSymbols)
		}
		if err.Error() != errNameSymbols {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, err.Error(), errNameSymbols)
		}
	})

	t.Run("set description", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		expect := "this is test bot"
		err = testBot.SetDescription(expect)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		if testBot.Description != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, testBot.Description, expect)
		}
	})

	t.Run("set too long description", func(t *testing.T) {
		expect := errDescrTooLong
		testBot, err := New("test")
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		testDescrs := []string{strings.Repeat("a", maxDescrLen+1), strings.Repeat("a", maxDescrLen*2)}
		for _, td := range testDescrs {
			err := testBot.SetDescription(td)
			if err == nil {
				t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, expect)
			}

			got := err.Error()
			if got != expect {
				t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, got, expect)
			}
		}
	})

	t.Run("set avatar", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		expect := "https://google.com/avatar.jpg"
		err = testBot.SetAvatar(expect)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		if testBot.Avatar != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, testBot.Avatar, expect)
		}
	})

	t.Run("set invalid avatar", func(t *testing.T) {
		expect := extlib.ErrInvalidUrl
		testBot, err := New("test")
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		testAvatars := []string{"http//goo.gl/img.png", "https//goo.gl/img.png", "", "helloworld"}
		for _, ta := range testAvatars {
			err := testBot.SetAvatar(ta)
			if err == nil {
				t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, expect)
			}

			got := err.Error()
			if got != expect {
				t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, got, expect)
			}
		}
	})

	t.Run("set commands", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		for _, cmd := range testCommands {
			err = testBot.AddCommand(cmd.Alias, cmd.Description)
			if err != nil {
				t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
			}
		}

		expect := len(testCommands)

		if len(testBot.Commands) != expect {
			t.Errorf("%#v\ngot: %d,\nexpect: %d\n", testBot, len(testBot.Commands), expect)
		}
	})

	t.Run("set invalid commands", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		// set empty alias
		err = testBot.AddCommand("", "")
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, errCmdAliasIsEmpty)
		}
		if err.Error() != errCmdAliasIsEmpty {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), errCmdAliasIsEmpty)
		}

		// set too long alias
		err = testBot.AddCommand(strings.Repeat("a", maxCmdAliasLen+1), "")
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, errCmdAliasTooLong)
		}
		if err.Error() != errCmdAliasTooLong {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), errCmdAliasTooLong)
		}

		// set invalid alias
		err = testBot.AddCommand("Test Cmd Alia$", "")
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, errCmdAliasSymbols)
		}
		if err.Error() != errCmdAliasSymbols {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), errCmdAliasSymbols)
		}

		// set too long description
		err = testBot.AddCommand("test", strings.Repeat("a", maxCmdDescrLen+1))
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testBot, nil, errCmdDescrTooLong)
		}
		if err.Error() != errCmdDescrTooLong {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), errCmdDescrTooLong)
		}
	})

	t.Run("find commands", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		testBot.Commands = make([]Command, len(testCommands))
		copy(testBot.Commands, testCommands)

		for _, expect := range testCommands {
			cmd, err := testBot.GetCommand(expect.Alias)
			if err != nil {
				t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
			}

			if cmd.Alias != expect.Alias {
				t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, cmd.Alias, expect.Alias)
			}

			if cmd.Description != expect.Description {
				t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, cmd.Description, expect.Description)
			}
		}
	})

	t.Run("set existent command", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		testBot.Commands = make([]Command, len(testCommands))
		copy(testBot.Commands, testCommands)

		expect := testCommands[0]
		err = testBot.AddCommand(expect.Alias, expect.Description)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		cmd, err := testBot.GetCommand(expect.Alias)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		if cmd.Alias != expect.Alias {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, cmd.Alias, expect.Alias)
		}

		if cmd.Description != expect.Description {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, cmd.Description, expect.Description)
		}
	})

	t.Run("remove commands", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		testBot.Commands = make([]Command, len(testCommands))
		copy(testBot.Commands, testCommands)

		for _, cmd := range testCommands {
			err := testBot.RemoveCommand(cmd.Alias)
			if err != nil {
				t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
			}
		}

		expect := 0

		if len(testBot.Commands) != expect {
			t.Errorf("%#v\ngot: %d,\nexpect: %d\n", testBot, len(testBot.Commands), expect)
		}
	})

	t.Run("find non-existent command", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		testCommand := testCommands[0]
		expect := Command{}

		cmd, err := testBot.GetCommand(testCommand.Alias)
		if err.Error() != errCmdNotFound {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, err.Error(), errCmdNotFound)
		}

		if cmd.Alias != expect.Alias {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, cmd.Alias, expect.Alias)
		}

		if cmd.Description != expect.Description {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, cmd.Description, expect.Description)
		}
	})

	t.Run("remove non-existent command", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		err = testBot.RemoveCommand(testCommands[0].Alias)
		expect := errCmdNotFound

		if err.Error() != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testBot, err.Error(), expect)
		}
	})

	t.Run("get all commands", func(t *testing.T) {
		testBot, err := New(testName)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
		}

		for _, cmd := range testCommands {
			err = testBot.AddCommand(cmd.Alias, cmd.Description)
			if err != nil {
				t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testBot, err.Error(), nil)
			}
		}

		expect := len(testCommands)
		result := len(testBot.GetCommands())

		if result != expect {
			t.Errorf("%#v\ngot: %d,\nexpect: %d\n", testBot, result, expect)
		}
	})
}
