package user

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	correctNicknames := []string{"Tom", "Rob_Pike", "-_-", strings.Repeat("x", maxNicknameLen)}
	incorrectNicknames := []string{"Rob Pike", "No%name", "ma$ter"}
	tooShortNicknames := []string{"", "x", "Xx"}
	tooLongNicknames := []string{strings.Repeat("x", maxNicknameLen+1), strings.Repeat("X", maxNicknameLen*2)}

	t.Run("check correct nicknames", func(t *testing.T) {
		for _, n := range correctNicknames {
			user, err := New(n)
			if err != nil {
				t.Errorf("%#v\ngot: %s,\nexpect: %v\n", user, err.Error(), nil)
			}
		}
	})

	t.Run("check incorrect nicknames", func(t *testing.T) {
		expect := errNicknameSymbols
		for _, n := range incorrectNicknames {
			user, err := New(n)
			if err.Error() != expect {
				t.Errorf("%#v\ngot: %s,\nexpect: %v\n", user, err.Error(), expect)
			}
		}
	})

	t.Run("check too short nicknames", func(t *testing.T) {
		expect := errNicknameTooShort
		for _, n := range tooShortNicknames {
			user, err := New(n)
			if err.Error() != expect {
				t.Errorf("%#v\ngot: %s,\nexpect: %s\n", user, err.Error(), expect)
			}
		}
	})

	t.Run("check too long nicknames", func(t *testing.T) {
		expect := errNicknameTooLong
		for _, n := range tooLongNicknames {
			user, err := New(n)
			if err.Error() != expect {
				t.Errorf("%#v\ngot: %s,\nexpect: %s\n", user, err.Error(), expect)
			}
		}
	})
}
