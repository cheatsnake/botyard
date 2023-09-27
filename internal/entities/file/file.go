package file

import (
	"botyard/internal/tools/ulid"
	"errors"
)

type File struct {
	Id       string `json:"id"`
	Path     string `json:"path"`
	Name     string `json:"name"`
	Size     int    `json:"size"`
	MimeType string `json:"mimeType"`
}

func New(path, name, mime string, size int) (*File, error) {
	if len(path) == 0 {
		return nil, errors.New(errPathIsEmpty)
	}

	err := validateMimeType(mime)
	if err != nil {
		return nil, err
	}

	return &File{
		Id:       ulid.New(),
		Path:     path,
		MimeType: mime,
	}, nil
}
