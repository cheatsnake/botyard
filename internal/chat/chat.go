package chat

import (
	"botyard/internal/tools/ulid"
	"errors"

	"golang.org/x/exp/slices"
)

type Chat struct {
	Id        string   `json:"id"`
	MemberIds []string `json:"memberIds"`
	store     store
}

type store interface {
	messageStore
	fileStore
}

type messageStore interface {
	AddMessage(msg *Message) error
	GetMessagesByChat(chatId string, page, limit int) (int, []*Message, error)
	DeleteMessagesByChat(chatId string) error
}

type fileStore interface {
	AddFile(file *File) error
	GetFile(id string) (*File, error)
	GetFiles(ids []string) ([]*File, error)
	DeleteFile(id string) error
}

func New(memberIds []string, s store) *Chat {
	return &Chat{
		Id:        ulid.New(),
		MemberIds: memberIds,
		store:     s,
	}
}

func (c *Chat) SendMessage(senderId, body string, fileIds []string) error {
	isMemeber := slices.Contains(c.MemberIds, senderId)
	if !isMemeber {
		return errors.New(errSenderNotMember)
	}

	msg := newMessage(c.Id, senderId, body, fileIds)

	err := (c.store).AddMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func (c *Chat) GetMessages(page, limit int) (MessagesPage, error) {
	total, msgs, err := (c.store).GetMessagesByChat(c.Id, page, limit)
	if err != nil {
		return MessagesPage{}, err
	}

	result := MessagesPage{
		ChatId:   c.Id,
		Total:    total,
		Page:     page,
		Limit:    limit,
		Messages: make([]messageWithoutChatId, len(msgs)),
	}

	for i, msg := range msgs {
		result.Messages[i] = messageWithoutChatId{
			Id:        msg.Id,
			SenderId:  msg.SenderId,
			Body:      msg.Body,
			FileIds:   msg.FileIds,
			Timestamp: msg.Timestamp,
		}
	}

	return result, nil
}

func (c *Chat) Clear() error {
	err := (c.store).DeleteMessagesByChat(c.Id)
	if err != nil {
		return err
	}

	return nil
}
