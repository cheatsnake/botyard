package services

import "botyard/internal/storage"

type MessageService struct {
	file  *FileService
	store storage.MessageStore
}

func NewMessageService(s storage.MessageStore, fs *FileService) *MessageService {
	return &MessageService{
		file:  fs,
		store: s,
	}
}
