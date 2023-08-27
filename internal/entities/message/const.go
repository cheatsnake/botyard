package message

import "fmt"

const (
	maxBodyLen = 4096
	maxFiles   = 10
)

var (
	errBodyTooLong     = fmt.Sprintf("maximum allowed body length is %d characters", maxBodyLen)
	errChatIdIsEmpty   = "message chat id must not be empty"
	errSenderIdIsEmpty = "message sender id must not be empty"
	errBodyIsEmpty     = "message body must not be empty"
	errTooManyFiles    = fmt.Sprintf("maximum number of files allowed is %d", maxFiles)
)
