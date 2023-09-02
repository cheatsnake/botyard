package mock

import "botyard/internal/entities/user"

type userStore struct{}

func (mus *userStore) AddUser(user *user.User) error {
	return nil
}

func (mus *userStore) GetUser(id string) (*user.User, error) {
	u := &user.User{
		Id:       id,
		Nickname: "test",
	}
	return u, nil
}

func UserStore() *userStore {
	return &userStore{}
}
