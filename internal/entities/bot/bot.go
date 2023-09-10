package bot

import (
	"botyard/internal/tools/ulid"
	"botyard/pkg/extlib"
	"sync"
)

type Bot struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	mu          sync.Mutex
}

func New(name string) (*Bot, error) {
	err := validateName(name)
	if err != nil {
		return nil, err
	}

	return &Bot{
		Id:   ulid.New(),
		Name: name,
	}, nil
}

func (b *Bot) SetName(name string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	err := validateName(name)
	if err != nil {
		return err
	}

	b.Name = name
	return nil
}

func (b *Bot) SetDescription(descr string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	err := validateDescription(descr)
	if err != nil {
		return err
	}

	b.Description = descr
	return nil
}

func (b *Bot) SetAvatar(avatar string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	err := extlib.ValidateURL(avatar)
	if err != nil {
		return err
	}

	b.Avatar = avatar
	return nil
}
