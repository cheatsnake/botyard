package handlers

import (
	"botyard/internal/storage"
	"botyard/internal/user"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	store storage.Storage
}

func NewUser(store storage.Storage) *User {
	return &User{
		store: store,
	}
}

type userBody struct {
	user.User
	Id struct{} `json:"-"`
}

func (s *User) CreateUser(c *fiber.Ctx) error {
	b := new(userBody)

	if err := c.BodyParser(b); err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	user, err := user.New(b.Nickname)
	if err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	err = s.store.AddUser(user)
	if err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
