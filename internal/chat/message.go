package chat

import (
	"botyard/pkg/ulid"
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

type MessagesPage struct {
	ChatId   string                 `json:"chatId"`
	Total    int                    `json:"total"`
	Page     int                    `json:"page"`
	Limit    int                    `json:"limit"`
	Messages []messageWithoutChatId `json:"messages"`
}

type messageWithoutChatId struct {
	Id        string    `json:"id"`
	SenderId  string    `json:"senderId"`
	Body      string    `json:"body"`
	FileIds   []string  `json:"fileIds,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

func newMessage(chatId, senderId, body string, fileIds []string) *Message {
	return &Message{
		Id:        ulid.New(),
		ChatId:    chatId,
		SenderId:  senderId,
		Body:      body,
		FileIds:   fileIds,
		Timestamp: time.Now(),
	}
}
