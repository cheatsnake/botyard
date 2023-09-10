package bot

type Command struct {
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
		BotId:       botId,
		Alias:       alias,
		Description: descr,
	}, nil
}
