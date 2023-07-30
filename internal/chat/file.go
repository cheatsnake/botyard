package chat

import "botyard/pkg/ulid"

type File struct {
	Id        string
	Content   []byte
	Extension string
}

type FileStore interface {
	AddFile(file *File) error
	GetFile(id string) (*File, error)
	GetFiles(ids []string) ([]*File, error)
	DeleteFile(id string) error
}

func newFile(content []byte, ext string) *File {
	return &File{
		Id:        ulid.New(),
		Content:   content,
		Extension: ext,
	}
}
