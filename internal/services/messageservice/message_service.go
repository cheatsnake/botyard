package messageservice

import (
	"botyard/internal/entities/file"
	"botyard/internal/entities/message"
	"botyard/internal/services/fileservice"
	"botyard/internal/storage"
	"botyard/pkg/exterr"
	"time"
)

type Service struct {
	file  *fileservice.Service
	store storage.MessageStore
}

type CreateBody struct {
	message.Message
	Id        struct{} `json:"-"`
	Timestamp struct{} `json:"-"`
}

type PreparedMessage struct {
	Id        string       `json:"id"`
	SenderId  string       `json:"senderId"`
	Body      string       `json:"body"`
	Files     []*file.File `json:"files,omitempty"`
	Timestamp time.Time    `json:"timestamp"`
}

type MessagesPage struct {
	ChatId   string             `json:"chatId"`
	Messages []*PreparedMessage `json:"messages"`
	Total    int                `json:"total"`
	Page     int                `json:"page"`
	Limit    int                `json:"limit"`
}

func New(s storage.MessageStore, fs *fileservice.Service) *Service {
	return &Service{
		file:  fs,
		store: s,
	}
}

func (ms *Service) AddMessage(body *CreateBody) error {
	msg, err := message.New(body.ChatId, body.SenderId, body.Body, body.FileIds)
	if err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	err = ms.store.AddMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func (ms *Service) GetMessagesByChat(chatId string, page, limit int) (*MessagesPage, error) {
	total, msgs, err := ms.store.GetMessagesByChat(chatId, page, limit)
	if err != nil {
		return nil, err
	}

	msgsWithFiles := make([]*PreparedMessage, len(msgs))

	for i, msg := range msgs {
		var files []*file.File

		if len(msg.FileIds) > 0 {
			files, err = ms.file.GetFiles(msg.FileIds)
			if err != nil {
				return nil, err
			}
		}

		msgsWithFiles[i] = &PreparedMessage{
			Id:        msg.Id,
			SenderId:  msg.SenderId,
			Body:      msg.Body,
			Files:     files,
			Timestamp: msg.Timestamp,
		}
	}

	return &MessagesPage{
		ChatId:   chatId,
		Messages: msgsWithFiles,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

func (ms *Service) DeleteMessagesByChat(chatId string) error {
	err := ms.store.DeleteMessagesByChat(chatId)
	if err != nil {
		return err
	}

	return nil
}
