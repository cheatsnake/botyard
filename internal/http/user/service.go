package user

import (
	"botyard/internal/entities/user"
	"botyard/internal/http/helpers"
	"botyard/internal/storage"
	"net/http"
)

type Service struct {
	store storage.Storage
}

func NewService(s storage.Storage) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) Create(body *createBody) (*user.User, error) {
	newUser, err := user.New(body.Nickname)
	if err != nil {
		return nil, helpers.NewHttpError(http.StatusBadRequest, err.Error())
	}

	err = s.store.AddUser(newUser)
	if err != nil {
		return nil, helpers.NewHttpError(http.StatusBadRequest, err.Error())
	}

	return newUser, nil
}
