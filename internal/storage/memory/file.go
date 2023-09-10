package memory

import (
	"botyard/internal/entities/file"
	"botyard/pkg/extlib"

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

	return nil, extlib.ErrorNotFound("file not found")
}

func (s *Storage) GetFiles(ids []string) ([]*file.File, error) {
	files := make([]*file.File, 0, len(ids))

	for _, f := range s.files {
		if slices.Contains(ids, f.Id) {
			files = append(files, f)
		}
	}

	return files, nil
}

func (s *Storage) DeleteFile(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delIndex := slices.IndexFunc(s.files, func(f *file.File) bool {
		return f.Id == id
	})

	if delIndex == -1 {
		return extlib.ErrorNotFound("file not found")
	}

	s.messages = extlib.SliceRemoveElement(s.messages, delIndex)
	return nil
}
