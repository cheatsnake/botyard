package user

import "botyard/pkg/ulid"

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func New(name string) *User {
	return &User{
		Id:   ulid.New(),
		Name: name,
	}
}
