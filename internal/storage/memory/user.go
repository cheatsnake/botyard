package memory

import (
	"botyard/internal/entities/user"
	"errors"
)

func (s *Storage) AddUser(user *user.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.users = append(s.users, user)
	return nil
}

func (s *Storage) GetUser(id string) (*user.User, error) {
	for _, user := range s.users {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, errors.New("bot not found")
}
