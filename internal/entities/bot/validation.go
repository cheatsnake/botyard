package bot

import (
	"errors"
)

func validateName(name string) error {
	if len(name) == 0 {
		return errors.New(errNameIsEmpty)
	}

	if len(name) > maxNameLen {
		return errors.New(errNameTooLong)
	}

	if !regexName.MatchString(name) {
		return errors.New(errNameSymbols)
	}

	return nil
}

func validateDescription(descr string) error {
	if len(descr) > maxDescrLen {
		return errors.New(errDescrTooLong)
	}

	return nil
}

func validateCmd(alias, descr string) error {
	if len(alias) == 0 {
		return errors.New(errCmdAliasIsEmpty)
	}

	if len(alias) > maxCmdAliasLen {
		return errors.New(errCmdAliasTooLong)
	}

	if !regexCmdAlias.MatchString(alias) {
		return errors.New(errCmdAliasSymbols)
	}

	if len(descr) > maxCmdDescrLen {
		return errors.New(errCmdDescrTooLong)
	}

	return nil
}
