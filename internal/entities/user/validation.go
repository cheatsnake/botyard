package user

import (
	"errors"
	"fmt"

	"github.com/cheatsnake/botyard/internal/config"
)

func validateNickname(nick string) error {
	if len(nick) < config.Global.Limits.User.MinNicknameLength {
		return errors.New(errNicknameTooShort(config.Global.Limits.User.MinNicknameLength))
	}

	if len(nick) > config.Global.Limits.User.MaxNicknameLength {
		return errors.New(errNicknameTooLong(config.Global.Limits.User.MaxNicknameLength))
	}

	if !regexNickname.MatchString(nick) {
		return fmt.Errorf("%s", errNicknameSymbols)
	}

	return nil
}
