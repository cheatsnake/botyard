package user

import (
	"fmt"
	"regexp"

	"github.com/cheatsnake/botyard/internal/config"
)

var (
	errNicknameTooShort = func(minLen int) string {
		return fmt.Sprintf(
			"Nickname must be at least %d characters long.",
			config.Global.Limits.User.MinNicknameLength,
		)
	}
	errNicknameTooLong = func(minLen int) string {
		return fmt.Sprintf(
			"Nickname should not be longer than %d characters.",
			config.Global.Limits.User.MaxNicknameLength,
		)
	}
)

var regexNickname = regexp.MustCompile(`^[A-Za-z0-9 -]+$`)

const errNicknameSymbols = "Nickname should contain only upper/lower case English letters, spaces, dashes and numbers."
