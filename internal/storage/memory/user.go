package memory

import (
	"botyard/internal/entities/user"
	"botyard/pkg/exterr"
	"botyard/pkg/extlib"

	"golang.org/x/exp/slices"
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

	return nil, exterr.ErrorNotFound("user not found")
}

func (s *Storage) DeleteUser(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delIndex := slices.IndexFunc(s.users, func(u *user.User) bool {
		return u.Id == id
	})

	if delIndex == -1 {
		return exterr.ErrorNotFound("user not found")
	}

	s.users = extlib.SliceRemoveElement(s.users, delIndex)
	return nil
}
