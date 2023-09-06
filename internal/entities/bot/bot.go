package bot

import (
	"botyard/internal/tools/ulid"
	"botyard/pkg/extlib"
	"errors"
	"sync"
)

type Bot struct {
	Id          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Avatar      string    `json:"avatar,omitempty"`
	Commands    []Command `json:"commands,omitempty"`
	mu          sync.Mutex
}

// Command represent a Bot's action that contain short alias name and description to it
type Command struct {
	Alias       string `json:"alias"`
	Description string `json:"description"`
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

func (b *Bot) AddCommand(alias, descr string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	err := validateCmd(alias, descr)
	if err != nil {
		return err
	}

	for _, cmd := range b.Commands {
		if cmd.Alias == alias {
			cmd.Description = descr
			return nil
		}
	}

	newCmd := Command{Alias: alias, Description: descr}
	b.Commands = append(b.Commands, newCmd)
	return nil
}

func (b *Bot) GetCommand(alias string) (Command, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, cmd := range b.Commands {
		if cmd.Alias == alias {
			return cmd, nil
		}
	}

	return Command{}, errors.New(errCmdNotFound)
}

func (b *Bot) GetCommands() []Command {
	return b.Commands
}

func (b *Bot) RemoveCommand(alias string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	for i, cmd := range b.Commands {
		if cmd.Alias == alias {
			b.Commands = append(b.Commands[:i], b.Commands[i+1:]...)
			return nil
		}
	}

	return errors.New(errCmdNotFound)
}
