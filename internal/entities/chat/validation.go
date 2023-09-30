package chat

import (
	"botyard/internal/config"
	"errors"
)

func validateBody(body string) error {
	if len(body) < 1 {
		return errors.New(errBodyIsEmpty)
	}

	if len(body) > config.Global.Limits.Message.MaxBodyLength {
		return errors.New(errBodyTooLong(config.Global.Limits.Message.MaxBodyLength))
	}

	return nil
}

func validateAttachmentIds(ids []string) error {
	if len(ids) > config.Global.Limits.Message.MaxAttachedFiles {
		return errors.New(errTooManyFiles(config.Global.Limits.Message.MaxAttachedFiles))
	}

	return nil
}
