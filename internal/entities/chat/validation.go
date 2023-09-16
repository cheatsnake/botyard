package chat

import (
	"errors"
)

func validateBody(body string) error {
	if len(body) < 1 {
		return errors.New(errBodyIsEmpty)
	}

	if len(body) > maxBodyLen {
		return errors.New(errBodyTooLong)
	}

	return nil
}

func validateFileIds(fileIds []string) error {
	if len(fileIds) > maxFiles {
		return errors.New(errTooManyFiles)
	}

	return nil
}