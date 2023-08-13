package memory

import (
	"botyard/internal/entities/file"
	"botyard/pkg/extlib"
	"errors"

	"golang.org/x/exp/slices"
)

func (s *Storage) AddFile(file *file.File) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.files = append(s.files, file)
	return nil
}

func (s *Storage) GetFile(id string) (*file.File, error) {
	for _, file := range s.files {
		if file.Id == id {
			return file, nil
		}
	}

	return nil, errors.New("file not found")
}

func (s *Storage) GetFiles(ids []string) ([]*file.File, error) {
	files := make([]*file.File, 0, len(ids))

	for _, file := range s.files {
		if slices.Contains(ids, file.Id) {
			files = append(files, file)
		}
	}

	return files, nil
}

func (s *Storage) DeleteFile(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	idx := slices.IndexFunc(s.files, func(f *file.File) bool {
		return f.Id == id
	})

	if idx == -1 {
		return errors.New("file not found")
	}

	s.messages = extlib.SliceRemoveElement(s.messages, idx)
	return nil
}
