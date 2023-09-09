package services

import (
	"botyard/internal/entities/file"
	"botyard/internal/storage"
	"botyard/pkg/extlib"
)

type FileService struct {
	store storage.FileStore
}

func NewFileService(s storage.FileStore) *FileService {
	return &FileService{
		store: s,
	}
}

func (fs *FileService) AddFile(path, mime string) (*file.File, error) {
	newFile, err := file.New(path, mime)
	if err != nil {
		return nil, extlib.ErrorBadRequest(err.Error())
	}

	err = fs.store.AddFile(newFile)
	if err != nil {
		return nil, err
	}

	return newFile, nil
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
