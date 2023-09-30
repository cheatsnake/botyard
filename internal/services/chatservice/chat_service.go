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
	Id          string       `json:"id"`
	SenderId    string       `json:"senderId"`
	Body        string       `json:"body"`
	Attachments []*file.File `json:"attachments,omitempty"`
	Timestamp   time.Time    `json:"timestamp"`
}

type MessagesPage struct {
	ChatId   string             `json:"chatId"`
	Total    int                `json:"total"`
	Page     int                `json:"page"`
	Limit    int                `json:"limit"`
	Messages []*PreparedMessage `json:"messages"`
}

func New(s storage.ChatStore, fs *fileservice.Service) *Service {
	return &Service{
		fileService: fs,
		store:       s,
	}
}

func (s *Service) CreateChat(userId string, botId string) (*chat.Chat, error) {
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

func (s *Service) GetChat(id string) (*chat.Chat, error) {
	ch, err := s.store.GetChat(id)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (s *Service) GetChats(userId string, botId string) ([]*chat.Chat, error) {
	chats, err := s.store.GetChats(userId, botId)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (s *Service) DeleteChat(id string) error {
	err := s.store.DeleteChat(id)
	if err != nil {
		return err
	}

	s.DeleteMessagesByChat(id)
	return nil
}

func (s *Service) CheckChatAccess(chatId, botId, userId string) (*chat.Chat, error) {
	chat, err := s.GetChat(chatId)
	if err != nil {
		return nil, err
	}

	if chat.UserId != userId {
		return nil, exterr.ErrorForbidden("chat is not related to current user")
	}

	if chat.BotId != botId {
		return nil, exterr.ErrorForbidden("chat is not related to current bot")
	}

	return chat, nil
}

func (ms *Service) AddMessage(body *CreateBody) (*chat.Message, error) {
	msg, err := chat.NewMessage(body.ChatId, body.SenderId, body.Body, body.AttachmentIds)
	if err != nil {
		return nil, exterr.ErrorBadRequest(err.Error())
	}

	err = ms.store.AddMessage(msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ms *Service) GetMessagesByChat(chatId string, page, limit int) (*MessagesPage, error) {
	total, msgs, err := ms.store.GetMessagesByChat(chatId, page, limit)
	if err != nil {
		return nil, err
	}

	msgsWithAttachments := make([]*PreparedMessage, len(msgs))

	for i, msg := range msgs {
		var attachments []*file.File

		if len(msg.AttachmentIds) > 0 {
			attachments, err = ms.fileService.GetFiles(msg.AttachmentIds)
			if err != nil {
				return nil, err
			}
		}

		msgsWithAttachments[i] = &PreparedMessage{
			Id:          msg.Id,
			SenderId:    msg.SenderId,
			Body:        msg.Body,
			Attachments: attachments,
			Timestamp:   msg.Timestamp,
		}
	}

	return &MessagesPage{
		ChatId:   chatId,
		Messages: msgsWithAttachments,
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
