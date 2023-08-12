package handlers

import (
	"botyard/internal/storage"
	"botyard/internal/user"
	"time"

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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := user.New(b.Nickname)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = s.store.AddUser(user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cookie := &fiber.Cookie{
		Name:    "userId",
		Value:   user.Id,
		Expires: time.Now().Add(24 * time.Hour),
	}

	c.Cookie(cookie)
	return c.Status(fiber.StatusCreated).JSON(user)
}
