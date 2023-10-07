package sqlite

import (
	"fmt"

	"github.com/cheatsnake/botyard/internal/entities/file"
	"github.com/cheatsnake/botyard/pkg/exterr"
	"github.com/cheatsnake/botyard/pkg/extlib"
)

func (s *Storage) AddFile(file *file.File) error {
	q := `INSERT INTO files (id, path, name, size, mime_type) VALUES (?, ?, ?, ?, ?)`

	_, err := s.conn.Exec(q, file.Id, file.Path, file.Name, file.Size, file.MimeType)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't save file to database.")
	}

	return nil
}

func (s *Storage) GetFiles(ids []string) ([]*file.File, error) {
	ph := extlib.SQLQueryPlaceholders(len(ids))
	q := fmt.Sprintf("SELECT id, path, name, size, mime_type FROM files WHERE id IN (%s)", ph)
	argIds := make([]interface{}, 0, len(ids))

	for _, id := range ids {
		argIds = append(argIds, id)
	}

	rows, err := s.conn.Query(q, argIds...)
	if err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't get files from database.")
	}
	defer rows.Close()

	files := make([]*file.File, 0, len(ids))

	for rows.Next() {
		var f file.File
		if err := rows.Scan(&f.Id, &f.Path, &f.Name, &f.Size, &f.MimeType); err != nil {
			// TODO logger
			return nil, exterr.ErrorInternalServer("Can't retrive file from database.")
		}
		files = append(files, &f)
	}

	if err := rows.Err(); err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't retrive files from database.")
	}

	return files, nil
}

func (s *Storage) DeleteFiles(ids []string) error {
	ph := extlib.SQLQueryPlaceholders(len(ids))
	q := fmt.Sprintf("DELETE FROM files WHERE id IN (%s)", ph)
	argIds := make([]interface{}, 0, len(ids))

	for _, id := range ids {
		argIds = append(argIds, id)
	}

	_, err := s.conn.Exec(q, argIds...)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't delete files from database.")
	}

	return nil
}
