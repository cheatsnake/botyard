package file

import (
	"errors"
	"mime"
)

func validateMimeType(mt string) error {
	ext, err := mime.ExtensionsByType(mt)
	if len(ext) == 0 || err != nil {
		return errors.New(errInvalidMimeType)
	}

	return nil
}
