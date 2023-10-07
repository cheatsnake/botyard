package sqlite

import (
	"database/sql"

	"github.com/cheatsnake/botyard/internal/entities/user"
	"github.com/cheatsnake/botyard/pkg/exterr"
)

func (s *Storage) AddUser(user *user.User) error {
	q := `INSERT INTO users (id, nickname) VALUES (?, ?)`

	_, err := s.conn.Exec(q, user.Id, user.Nickname)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't save user to database.")
	}

	return nil
}

func (s *Storage) GetUser(id string) (*user.User, error) {
	q := `SELECT id, nickname FROM users WHERE id = ?`

	var u user.User

	err := s.conn.QueryRow(q, id).Scan(&u.Id, &u.Nickname)
	if err == sql.ErrNoRows {
		return nil, exterr.ErrorNotFound("User not found.")
	}
	if err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't get user from database.")
	}

	return &u, nil
}

func (s *Storage) DeleteUser(id string) error {
	q := `DELETE FROM users WHERE id = ?`

	_, err := s.conn.Exec(q, id)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't delete user from database.")
	}

	return nil
}
