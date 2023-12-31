package chatservice

import (
	"github.com/cheatsnake/botyard/internal/entities/chat"
	"github.com/cheatsnake/botyard/internal/entities/file"
	"github.com/cheatsnake/botyard/internal/services/fileservice"
	"github.com/cheatsnake/botyard/internal/storage"
	"github.com/cheatsnake/botyard/pkg/exterr"
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
	ChatId      string       `json:"chatId,omitempty"`
	SenderId    string       `json:"senderId"`
	Body        string       `json:"body"`
	Attachments []*file.File `json:"attachments,omitempty"`
	Timestamp   int64        `json:"timestamp"`
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
	chats, err := s.store.GetChats(userId, botId)
	if err != nil {
		return nil, err
	}

	if len(chats) >= 10 {
		return nil, exterr.ErrorBadRequest("Cannot create more than 10 chat rooms with the same bot.")
	}

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

	if len(userId) == 0 && len(botId) == 0 {
		return nil, exterr.ErrorForbidden("User id or bot id is required.")
	}

	if len(userId) > 0 && chat.UserId != userId {
		return nil, exterr.ErrorForbidden("Chat is not related to current user.")
	}

	if len(botId) > 0 && chat.BotId != botId {
		return nil, exterr.ErrorForbidden("Chat is not related to current bot.")
	}

	return chat, nil
}

func (ms *Service) AddMessage(body *CreateBody) (*PreparedMessage, error) {
	msg, err := chat.NewMessage(body.ChatId, body.SenderId, body.Body, body.AttachmentIds)
	if err != nil {
		return nil, exterr.ErrorBadRequest(err.Error())
	}

	files, err := ms.fileService.GetFiles(body.AttachmentIds)
	if err != nil {
		return nil, err
	}

	err = ms.store.AddMessage(msg)
	if err != nil {
		return nil, err
	}

	return &PreparedMessage{
		Id:          msg.Id,
		ChatId:      msg.ChatId,
		SenderId:    msg.SenderId,
		Body:        msg.Body,
		Attachments: files,
		Timestamp:   msg.Timestamp,
	}, nil
}

func (s *Service) GetMessage(id string) (*chat.Message, error) {
	msg, err := s.store.GetMessage(id)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ms *Service) GetMessagesByChat(chatId, senderId string, page, limit int, since int64) (*MessagesPage, error) {
	total, msgs, err := ms.store.GetMessagesByChat(chatId, senderId, page, limit, since)
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
