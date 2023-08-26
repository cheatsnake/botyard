package user

import "fmt"

const (
	minNicknameLen = 3
	maxNicknameLen = 32
)
const (
	regexNickname = `^[A-Za-z0-9_-]+$`
)

var (
	errNicknameTooShort = fmt.Sprintf("nickname must be at least %d characters long", minNicknameLen)
	errNicknameTooLong  = fmt.Sprintf("nickname should not be longer than %d characters", maxNicknameLen)
	errNicknameSymbols  = "nickname should contain only upper/lower case English letters, underscores, dashes and numbers"
)
