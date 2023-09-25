package user

import (
	"botyard/internal/config"
	"fmt"
	"regexp"
)

var (
	errNicknameTooShort = func(minLen int) string {
		return fmt.Sprintf(
			"nickname must be at least %d characters long",
			config.Global.Limits.User.MinNicknameLength,
		)
	}
	errNicknameTooLong = func(minLen int) string {
		return fmt.Sprintf(
			"nickname should not be longer than %d characters",
			config.Global.Limits.User.MaxNicknameLength,
		)
	}
)

var regexNickname = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)

const errNicknameSymbols = "nickname should contain only upper/lower case English letters, underscores, dashes and numbers"
