package chat

import "fmt"

var (
	errBodyTooLong = func(maxBodyLen int) string {
		return fmt.Sprintf("maximum allowed body length is %d characters", maxBodyLen)
	}
	errTooManyFiles = func(maxFiles int) string {
		return fmt.Sprintf("maximum number of files allowed is %d", maxFiles)
	}

	errChatIdIsEmpty   = "message chat id must not be empty"
	errUserIdIsEmpty   = "user id must not be empty"
	errBotIdIsEmpty    = "bot id must not be empty"
	errSenderIdIsEmpty = "message sender id must not be empty"
)
