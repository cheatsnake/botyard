package services

import (
	"botyard/internal/entities/user"
	"testing"
)

func TestUserService(t *testing.T) {
	testStore := &mockUserStore{}
	testService := NewUserService(testStore)
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
}

type mockUserStore struct{}

func (mus *mockUserStore) AddUser(user *user.User) error {
	return nil
}

func (mus *mockUserStore) GetUser(id string) (*user.User, error) {
	u := &user.User{
		Id:       id,
		Nickname: "test",
	}
	return u, nil
}
