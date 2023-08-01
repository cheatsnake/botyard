package extlib

import (
	"fmt"
	"mime"
)

func ExtensionFromContentType(contentType string) (string, error) {
	ext, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return "", err
	}

	if len(ext) == 0 {
		return "", fmt.Errorf("no file extension found for Content-Type: %s", contentType)
	}

	return ext[0], nil
}
