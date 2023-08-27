package bot

import (
	"fmt"
	"regexp"
)

const (
	maxNameLen     = 32
	maxDescrLen    = 512
	maxAvatarLen   = 256
	maxCmdAliasLen = 32
	maxCmdDescrLen = 128
)

var (
	regexName     = regexp.MustCompile(`^[a-zA-Z0-9 \-]+$`)
	regexCmdAlias = regexp.MustCompile(`^[a-z0-9_\-]+$`)
)

var (
	errNameIsEmpty     = "bot name must not be empty"
	errNameTooLong     = fmt.Sprintf("bot name should not be longer than %d characters", maxNameLen)
	errNameSymbols     = "bot name should contain upper/lower case English letters, spaces, dashes or numbers"
	errDescrTooLong    = fmt.Sprintf("bot description should not be longer than %d characters", maxDescrLen)
	errCmdAliasIsEmpty = "command alias must not be empty"
	errCmdAliasTooLong = fmt.Sprintf("command alias should not be longer than %d characters", maxCmdAliasLen)
	errCmdAliasSymbols = "command alias should contain lower case English letters, underscores, dashes or numbers"
	errCmdDescrTooLong = fmt.Sprintf("command description should not be longer than %d characters", maxCmdDescrLen)
	errCmdNotFound     = "command not found"
)
