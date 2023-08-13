package user

import (
	"botyard/internal/entities/user"
	"botyard/internal/storage"
	"time"

	"github.com/gofiber/fiber/v2"
)

type handlers struct {
	service *service
}

func Handlers(s storage.Storage) *handlers {
	return &handlers{
		service: &service{
			store: s,
		},
	}
}

func (h *handlers) CreateUser(c *fiber.Ctx) error {
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

	err = h.service.store.AddUser(newUser)
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
