package chat

import "botyard/pkg/ulid"

type File struct {
	Id       string `json:"id"`
	Path     string `json:"path"`
	MimeType string `json:"mimeType"`
}

type FileStore interface {
	AddFile(file *File) error
	GetFile(id string) (*File, error)
	GetFiles(ids []string) ([]*File, error)
	DeleteFile(id string) error
}

func NewFile(path, mime string) *File {
	return &File{
		Id:       ulid.New(),
		Path:     path,
		MimeType: mime,
	}
}
