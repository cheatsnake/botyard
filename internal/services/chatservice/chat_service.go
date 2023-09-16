package chatservice

import (
	"botyard/internal/entities/chat"
	"botyard/internal/entities/file"
	"botyard/internal/services/fileservice"
	"botyard/internal/storage"
	"botyard/pkg/exterr"
	"time"
)

type Service struct {
	fileService *fileservice.Service
	store       storage.ChatStore
}

type CreateBody struct {
	chat.Message
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

func New(s storage.ChatStore, fs *fileservice.Service) *Service {
	return &Service{
		fileService: fs,
		store:       s,
	}
}

func (s *Service) Create(userId string, botId string) (*chat.Chat, error) {
	chat, err := chat.New(userId, botId)
	if err != nil {
		return nil, exterr.ErrorBadRequest(err.Error())
	}

	err = s.store.AddChat(chat)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (s *Service) GetChats(userId string, botId string) ([]*chat.Chat, error) {
	chats, err := s.store.GetChats(userId, botId)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (s *Service) Delete(id string) error {
	err := s.store.DeleteChat(id)
	if err != nil {
		return err
	}

	return nil
}

func (ms *Service) AddMessage(body *CreateBody) error {
	msg, err := chat.NewMessage(body.ChatId, body.SenderId, body.Body, body.FileIds)
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
			files, err = ms.fileService.GetFiles(msg.FileIds)
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
