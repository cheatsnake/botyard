package bot

import "sync"

type Webhook struct {
	BotId  string `json:"botId"`
	Url    string `json:"url"`
	Secret string `json:"secret,omitempty"`
	mu     sync.Mutex
}

func NewWebhook(botId, url, secret string) (*Webhook, error) {
	err := validateBotId(botId)
	if err != nil {
		return nil, err
	}

	err = validateWebhookUrl(url)
	if err != nil {
		return nil, err
	}

	err = validateWebhookSecret(secret)
	if err != nil {
		return nil, err
	}

	return &Webhook{
		BotId:  botId,
		Url:    url,
		Secret: secret,
	}, nil
}

func (w *Webhook) SetUrl(url string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	err := validateWebhookUrl(url)
	if err != nil {
		return err
	}

	w.Url = url
	return nil
}

func (w *Webhook) SetSecret(secret string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	err := validateWebhookSecret(secret)
	if err != nil {
		return err
	}

	w.Secret = secret
	return nil
}
