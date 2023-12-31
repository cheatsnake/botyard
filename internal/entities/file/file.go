package file

import (
	"errors"

	"github.com/cheatsnake/botyard/internal/tools/ulid"
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

	if len(name) == 0 {
		return nil, errors.New(errNameIsEmpty)
	}

	err := validateMimeType(mime)
	if err != nil {
		return nil, err
	}

	return &File{
		Id:       ulid.New(),
		Path:     path,
		Name:     name,
		Size:     size,
		MimeType: mime,
	}, nil
}
