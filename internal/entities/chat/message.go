package chat

import (
	"errors"
	"time"

	"github.com/cheatsnake/botyard/internal/tools/ulid"
)

type Message struct {
	Id            string   `json:"id"`
	ChatId        string   `json:"chatId"`
	SenderId      string   `json:"senderId"`
	Body          string   `json:"body"`
	AttachmentIds []string `json:"attachmentIds,omitempty"`
	Timestamp     int64    `json:"timestamp"`
}

func NewMessage(chatId, senderId, body string, attachmentIds []string) (*Message, error) {
	if len(chatId) == 0 {
		return nil, errors.New(errChatIdIsEmpty)
	}

	if len(senderId) == 0 {
		return nil, errors.New(errSenderIdIsEmpty)
	}

	err := validateBody(body)
	if err != nil {
		return nil, err
	}

	err = validateAttachmentIds(attachmentIds)
	if err != nil {
		return nil, err
	}

	return &Message{
		Id:            ulid.New(),
		ChatId:        chatId,
		SenderId:      senderId,
		Body:          body,
		AttachmentIds: attachmentIds,
		Timestamp:     time.Now().UnixMilli(),
	}, nil
}
