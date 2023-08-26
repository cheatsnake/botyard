package user

import (
	"errors"
	"regexp"
)

func validateNickname(nick string) error {
	if len(nick) < minNicknameLen {
		return errors.New(errNicknameTooShort)
	}

	if len(nick) > maxNicknameLen {
		return errors.New(errNicknameTooLong)
	}

	re, _ := regexp.Compile(regexNickname)
	if !re.MatchString(nick) {
		return errors.New(errNicknameSymbols)
	}

	return nil
}
