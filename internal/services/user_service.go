package services

import (
	"botyard/internal/entities/user"
	"botyard/internal/storage"
	"botyard/pkg/extlib"
)

type UserService struct {
	store storage.UserStore
}

type UserCreateBody struct {
	user.User
	Id struct{} `json:"-"`
}

func NewUserService(s storage.UserStore) *UserService {
	return &UserService{
		store: s,
	}
}

func (s *UserService) Create(body *UserCreateBody) (*user.User, error) {
	newUser, err := user.New(body.Nickname)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	err = s.store.AddUser(newUser)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	return newUser, nil
}
