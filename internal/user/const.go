package user

import "fmt"

const (
	minNicknameLen = 3
	maxNicknameLen = 32
)

var (
	errNicknameTooShort = fmt.Sprintf("nickname must be at least %d characters long", minNicknameLen)
	errNicknameTooLong  = fmt.Sprintf("nickname should not be longer than %d characters", maxNicknameLen)
)
