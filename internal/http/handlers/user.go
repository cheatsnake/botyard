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

func (u *User) CreateUser(c *fiber.Ctx) error {
	body := new(struct {
		user.User
		Id struct{} `json:"-"`
	})

	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	newUser, err := user.New(body.Nickname)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = u.store.AddUser(newUser)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cookie := &fiber.Cookie{
		Name:    "userId",
		Value:   newUser.Id,
		Expires: time.Now().Add(24 * time.Hour),
	}

	c.Cookie(cookie)
	return c.Status(fiber.StatusCreated).JSON(newUser)
}
