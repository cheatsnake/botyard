package message

import "botyard/internal/tools/ulid"

func validateBody(body string) error {
	return nil
}

func validateFileIds(fileIds []string) error {
	for _, fi := range fileIds {
		if err := ulid.Verify(fi); err != nil {
			return err
		}
	}

	return nil
}
