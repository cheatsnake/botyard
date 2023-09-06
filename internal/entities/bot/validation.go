package bot

import (
	"botyard/pkg/extlib"
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

func validateBotId(id string) error {
	if len(id) == 0 {
		return errors.New(errBotIdIsEmpty)
	}

	return nil
}

func validateKeyToken(token string) error {
	if len(token) == 0 {
		return errors.New(errKeyTokenIsEmpty)
	}

	return nil
}

func validateWebhookUrl(url string) error {
	if len(url) == 0 {
		return errors.New(errWebhookUrlIsEmpty)
	}

	if len(url) > maxWebhookUrlLen {
		return errors.New(errWebhookUrlTooLong)
	}

	err := extlib.ValidateURL(url)
	if err != nil {
		return err
	}

	return nil
}

func validateWebhookSecret(secret string) error {
	if len(secret) > maxWebhookSecretLen {
		return errors.New(errWebhookSecretTooLong)
	}

	return nil
}
