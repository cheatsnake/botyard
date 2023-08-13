package message

import (
	"botyard/internal/tools/ulid"
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

func New(chatId, senderId, body string, fileIds []string) *Message {
	return &Message{
		Id:        ulid.New(),
		ChatId:    chatId,
		SenderId:  senderId,
		Body:      body,
		FileIds:   fileIds,
		Timestamp: time.Now(),
	}
}
