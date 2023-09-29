package userservice

import (
	"botyard/internal/entities/user"
	"botyard/internal/storage"
	"botyard/pkg/exterr"
)

type Service struct {
	store storage.UserStore
}

type CreateBody struct {
	user.User
	Id struct{} `json:"-"`
}

func New(s storage.UserStore) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) CreateUser(body *CreateBody) (*user.User, error) {
	newUser, err := user.New(body.Nickname)
	if err != nil {
		return nil, exterr.ErrorBadRequest(err.Error())
	}

	err = s.store.AddUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *Service) GetUserById(id string) (*user.User, error) {
	user, err := s.store.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
