package user

import "botyard/internal/entities/user"

type createBody struct {
	user.User
	Id struct{} `json:"-"`
}
