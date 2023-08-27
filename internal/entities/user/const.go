package user

import (
	"fmt"
	"regexp"
)

const (
	minNicknameLen = 3
	maxNicknameLen = 32
)

var (
	regexNickname = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)
)

var (
	errNicknameTooShort = fmt.Sprintf("nickname must be at least %d characters long", minNicknameLen)
	errNicknameTooLong  = fmt.Sprintf("nickname should not be longer than %d characters", maxNicknameLen)
	errNicknameSymbols  = "nickname should contain only upper/lower case English letters, underscores, dashes and numbers"
)
