package services

import (
	"botyard/internal/entities/file"
	"botyard/internal/storage"
)

type FileService struct {
	store storage.FileStore
}

func NewFileService(s storage.FileStore) *FileService {
	return &FileService{
		store: s,
	}
}

func (fs *FileService) AddFile(f *file.File) error {
	err := fs.store.AddFile(f)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileService) GetFiles(ids []string) ([]*file.File, error) {
	files, err := fs.store.GetFiles(ids)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (fs *FileService) DeleteFile(id string) error {
	err := fs.store.DeleteFile(id)
	if err != nil {
		return err
	}

	return nil
}
