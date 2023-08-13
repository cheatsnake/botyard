package message

import (
	"botyard/internal/http/file"
	"botyard/internal/storage"
)

type Service struct {
	file  *file.Service
	store storage.Storage
}

func NewService(s storage.Storage, fs *file.Service) *Service {
	return &Service{
		file:  fs,
		store: s,
	}
}
