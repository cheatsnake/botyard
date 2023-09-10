package bot

import (
	"botyard/internal/tools/ulid"
	"botyard/pkg/extlib"
	"strings"
	"testing"
)

func TestNewBot(t *testing.T) {
	testName := "test"

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
}
