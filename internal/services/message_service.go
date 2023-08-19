package services

import (
	"botyard/internal/entities/file"
	"botyard/internal/entities/message"
	"botyard/internal/storage"
	"time"
)

type MessageService struct {
	file  *FileService
	store storage.MessageStore
}

type CreateMessageBody struct {
	message.Message
	Id        struct{} `json:"-"`
	Timestamp struct{} `json:"-"`
}

type MessageWithFiles struct {
	Id        string       `json:"id"`
	ChatId    string       `json:"chatId"`
	SenderId  string       `json:"senderId"`
	Body      string       `json:"body"`
	Files     []*file.File `json:"files,omitempty"`
	Timestamp time.Time    `json:"timestamp"`
}

type MessagesPage struct {
	Messages []*MessageWithFiles `json:"messages"`
	Total    int                 `json:"total"`
	Page     int                 `json:"page"`
	Limit    int                 `json:"limit"`
}

func NewMessageService(s storage.MessageStore, fs *FileService) *MessageService {
	return &MessageService{
		file:  fs,
		store: s,
	}
}

func (ms *MessageService) AddMessage(body *CreateMessageBody) error {
	msg := message.New(body.ChatId, body.SenderId, body.Body, body.FileIds)
	err := ms.store.AddMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MessageService) GetMessagesByChat(chatId string, page, limit int) (*MessagesPage, error) {
	total, msgs, err := ms.store.GetMessagesByChat(chatId, page, limit)
	if err != nil {
		return nil, err
	}

	msgsWithFiles := make([]*MessageWithFiles, len(msgs))

	for i, msg := range msgs {
		var files []*file.File

		if len(msg.FileIds) > 0 {
			files, err = ms.file.GetFiles(msg.FileIds)
			if err != nil {
				return nil, err
			}
		}

		msgsWithFiles[i] = &MessageWithFiles{
			Id:        msg.Id,
			ChatId:    msg.ChatId,
			SenderId:  msg.SenderId,
			Body:      msg.Body,
			Files:     files,
			Timestamp: msg.Timestamp,
		}
	}

	return &MessagesPage{
		Messages: msgsWithFiles,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

func (ms *MessageService) DeleteMessagesByChat(chatId string) error {
	err := ms.store.DeleteMessagesByChat(chatId)
	if err != nil {
		return err
	}

	return nil
}
