package bot

import "github.com/cheatsnake/botyard/internal/tools/ulid"

type Command struct {
	Id          string `json:"id"`
	BotId       string `json:"botId"`
	Alias       string `json:"alias"`
	Description string `json:"description,omitempty"`
}

func NewCommand(botId, alias, descr string) (*Command, error) {
	err := validateBotId(botId)
	if err != nil {
		return nil, err
	}

	err = validateCmd(alias, descr)
	if err != nil {
		return nil, err
	}

	return &Command{
		Id:          ulid.New(),
		BotId:       botId,
		Alias:       alias,
		Description: descr,
	}, nil
}
