package bot

import (
	"botyard/internal/tools/ulid"
	"errors"
	"sync"
)

type Bot struct {
	Id          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Avatar      string     `json:"avatar,omitempty"`
	Commands    []Command  `json:"commands,omitempty"`
	mu          sync.Mutex `json:"-"`
}

// Command represent a Bot's action that contain short alias name and description to it
type Command struct {
	Alias       string `json:"alias"`
	Description string `json:"description"`
}

func New(name string) *Bot {
	return &Bot{
		Id:   ulid.New(),
		Name: name,
	}
}

func (b *Bot) SetName(name string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.Name = name
}

func (b *Bot) SetDescription(descr string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.Description = descr
}

func (b *Bot) SetAvatar(avatar string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.Avatar = avatar
}

func (b *Bot) AddCommand(alias, descr string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, cmd := range b.Commands {
		if cmd.Alias == alias {
			cmd.Description = descr
			return
		}
	}

	newCmd := Command{Alias: alias, Description: descr}
	b.Commands = append(b.Commands, newCmd)
}

func (b *Bot) Command(alias string) (Command, error) {
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
