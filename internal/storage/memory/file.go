package memory

import (
	"github.com/cheatsnake/botyard/internal/entities/file"
	"github.com/cheatsnake/botyard/pkg/exterr"
	"github.com/cheatsnake/botyard/pkg/extlib"

	"golang.org/x/exp/slices"
)

func (s *Storage) AddFile(file *file.File) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.files = append(s.files, file)
	return nil
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

func (s *Storage) DeleteFiles(ids []string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, id := range ids {
		delIndex := slices.IndexFunc(s.files, func(f *file.File) bool {
			return f.Id == id
		})

		if delIndex == -1 {
			return exterr.ErrorNotFound("File not found.")
		}

		s.messages = extlib.SliceRemoveElement(s.messages, delIndex)
	}

	return nil
}
