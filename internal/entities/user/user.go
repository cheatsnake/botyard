package user

import (
	"botyard/internal/tools/ulid"
)

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}

func New(nick string) (*User, error) {
	err := validateNickname(nick)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:       ulid.New(),
		Nickname: nick,
	}, nil
}
