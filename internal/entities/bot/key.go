package bot

import "sync"

type Key struct {
	BotId string
	Token string
	mu    sync.Mutex
}

func NewKey(botId, token string) (*Key, error) {
	err := validateBotId(botId)
	if err != nil {
		return nil, err
	}

	err = validateKeyToken(token)
	if err != nil {
		return nil, err
	}

	return &Key{
		BotId: botId,
		Token: token,
	}, nil
}

func (k *Key) SetToken(token string) error {
	k.mu.Lock()
	defer k.mu.Unlock()

	err := validateKeyToken(token)
	if err != nil {
		return err
	}

	k.Token = token
	return nil
}

func (k *Key) Assemble() string {
	return k.BotId + ":" + k.Token
}
