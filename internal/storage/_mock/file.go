package mock

import "github.com/cheatsnake/botyard/internal/entities/file"

type fileStore struct{}

func (mfs *fileStore) AddFile(file *file.File) error {
	return nil
}

func (mfs *fileStore) GetFiles(ids []string) ([]*file.File, error) {
	return []*file.File{}, nil
}

func (mfs *fileStore) DeleteFiles(ids []string) error {
	return nil
}

func FileStore() *fileStore {
	return &fileStore{}
}
