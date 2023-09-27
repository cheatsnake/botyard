package fileservice

import (
	"botyard/internal/entities/file"
	"botyard/internal/storage"
	"botyard/pkg/exterr"
)

type Service struct {
	store storage.FileStore
}

func New(s storage.FileStore) *Service {
	return &Service{
		store: s,
	}
}

func (fs *Service) AddFile(path, name, mime string, size int) (*file.File, error) {
	newFile, err := file.New(path, name, mime, size)
	if err != nil {
		return nil, exterr.ErrorBadRequest(err.Error())
	}

	err = fs.store.AddFile(newFile)
	if err != nil {
		return nil, err
	}

	return newFile, nil
}

func (fs *Service) GetFiles(ids []string) ([]*file.File, error) {
	files, err := fs.store.GetFiles(ids)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (fs *Service) DeleteFiles(ids []string) error {
	err := fs.store.DeleteFiles(ids)
	if err != nil {
		return err
	}

	return nil
}
