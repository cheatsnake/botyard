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
	return nil
}
