package message

import (
	"botyard/internal/tools/ulid"
	"errors"
	"time"
)

type Message struct {
	Id        string    `json:"id"`
	ChatId    string    `json:"chatId"`
	SenderId  string    `json:"senderId"`
	Body      string    `json:"body"`
	FileIds   []string  `json:"fileIds,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

func New(chatId, senderId, body string, fileIds []string) (*Message, error) {
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

	err = validateFileIds(fileIds)
	if err != nil {
		return nil, err
	}

	return &Message{
		Id:        ulid.New(),
		ChatId:    chatId,
		SenderId:  senderId,
		Body:      body,
		FileIds:   fileIds,
		Timestamp: time.Now(),
	}, nil
}
