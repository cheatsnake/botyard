package userservice

import (
	"errors"
	"testing"

	"github.com/cheatsnake/botyard/internal/entities/user"
	mock "github.com/cheatsnake/botyard/internal/storage/_mock"
	"github.com/cheatsnake/botyard/pkg/exterr"
)

func TestCreate(t *testing.T) {
	us := New(mock.UserStore())
	t.Run("create new user", func(t *testing.T) {
		testBody := &CreateBody{
			User: user.User{
				Nickname: "test",
			},
		}

		user, err := us.CreateUser(testBody)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", user, err, nil)
		}

		if user.Nickname != testBody.Nickname {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", user, user.Nickname, testBody.Nickname)
		}
	})

	t.Run("create new user with incorrect name", func(t *testing.T) {
		testBody := &CreateBody{
			User: user.User{
				Nickname: "-",
			},
		}

		var extErr *exterr.ExtendedError
		user, err := us.CreateUser(testBody)

		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", user, err, extErr)
		}

		if user != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", user, user, nil)
		}

		if !errors.As(err, &extErr) {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", user, err, extErr)
		}
	})
}
