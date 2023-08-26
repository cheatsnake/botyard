package message

import "fmt"

const (
	maxBodyLen = 4096
	maxFiles   = 10
)

var (
	errBodyTooLong  = fmt.Sprintf("maximum allowed body length is %d characters", maxBodyLen)
	errTooManyFiles = fmt.Sprintf("maximum number of files allowed is %d", maxFiles)
)
