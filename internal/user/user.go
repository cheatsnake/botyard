package user

import (
	"botyard/pkg/ulid"
	"errors"
)

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}

func New(name string) (*User, error) {
	if len(name) < minNicknameLen {
		return nil, errors.New(errNicknameTooShort)
	}

	if len(name) > maxNicknameLen {
		return nil, errors.New(errNicknameTooLong)
	}

	return &User{
		Id:       ulid.New(),
		Nickname: name,
	}, nil
}
