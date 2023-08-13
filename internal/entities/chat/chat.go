package chat

import (
	"botyard/internal/tools/ulid"
)

type Chat struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	BotId  string `json:"botId"`
}

func New(userId, botId string) *Chat {
	return &Chat{
		Id:     ulid.New(),
		UserId: userId,
		BotId:  botId,
	}
}
