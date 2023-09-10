package bot

import (
	"fmt"
	"regexp"
)

const (
	maxNameLen          = 32
	maxDescrLen         = 512
	maxAvatarLen        = 256
	maxCmdAliasLen      = 32
	maxCmdDescrLen      = 128
	maxWebhookUrlLen    = 128
	maxWebhookSecretLen = 64
)

var (
	regexName     = regexp.MustCompile(`^[a-zA-Z0-9 \-]+$`)
	regexCmdAlias = regexp.MustCompile(`^[a-z0-9_\-]+$`)
)

var (
	errNameIsEmpty          = "bot name must not be empty"
	errBotIdIsEmpty         = "bot id must not be empty"
	errKeyTokenIsEmpty      = "key token must not be empty"
	errWebhookUrlIsEmpty    = "webhook url must not be empty"
	errCmdAliasIsEmpty      = "command alias must not be empty"
	errNameTooLong          = fmt.Sprintf("bot name should not be longer than %d characters", maxNameLen)
	errDescrTooLong         = fmt.Sprintf("bot description should not be longer than %d characters", maxDescrLen)
	errCmdAliasTooLong      = fmt.Sprintf("command alias should not be longer than %d characters", maxCmdAliasLen)
	errCmdDescrTooLong      = fmt.Sprintf("command description should not be longer than %d characters", maxCmdDescrLen)
	errWebhookUrlTooLong    = fmt.Sprintf("webhook url should not be longer than %d characters", maxWebhookUrlLen)
	errWebhookSecretTooLong = fmt.Sprintf("webhook secret should not be longer than %d characters", maxWebhookSecretLen)
	errCmdAliasSymbols      = "command alias should contain lower case English letters, underscores, dashes or numbers"
	errNameSymbols          = "bot name should contain upper/lower case English letters, spaces, dashes or numbers"
)
