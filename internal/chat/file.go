package chat

import "botyard/pkg/ulid"

type File struct {
	Id       string `json:"id"`
	Path     string `json:"path"`
	MimeType string `json:"mimeType"`
}

func NewFile(path, mime string) *File {
	return &File{
		Id:       ulid.New(),
		Path:     path,
		MimeType: mime,
	}
}
