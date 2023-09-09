package services

import (
	"botyard/internal/entities/user"
	mock "botyard/internal/storage/_mock"
	"botyard/pkg/extlib"
	"errors"
	"testing"
)

func TestUserServiceCreate(t *testing.T) {
	testService := NewUserService(mock.UserStore())
	t.Run("create new user", func(t *testing.T) {
		testBody := &UserCreateBody{
			User: user.User{
				Nickname: "test",
			},
		}

		user, err := testService.Create(testBody)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", user, err, nil)
		}

		if user.Nickname != testBody.Nickname {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", user, user.Nickname, testBody.Nickname)
		}
	})

	t.Run("create new user with incorrect name", func(t *testing.T) {
		testBody := &UserCreateBody{
			User: user.User{
				Nickname: "-",
			},
		}

		var extErr *extlib.ExtendedError
		user, err := testService.Create(testBody)

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
