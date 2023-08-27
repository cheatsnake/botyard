package chat

import (
	"botyard/internal/tools/ulid"
	"errors"
)

type Chat struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	BotId  string `json:"botId"`
}

func New(userId, botId string) (*Chat, error) {
	if len(userId) == 0 {
		return nil, errors.New(errUserIdIsEmpty)
	}

	if len(botId) == 0 {
		return nil, errors.New(errBotIdIsEmpty)
	}

	return &Chat{
		Id:     ulid.New(),
		UserId: userId,
		BotId:  botId,
	}, nil
}
