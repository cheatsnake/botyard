package user

import (
	"errors"
)

func validateNickname(nick string) error {
	if len(nick) < minNicknameLen {
		return errors.New(errNicknameTooShort)
	}

	if len(nick) > maxNicknameLen {
		return errors.New(errNicknameTooLong)
	}

	if !regexNickname.MatchString(nick) {
		return errors.New(errNicknameSymbols)
	}

	return nil
}
