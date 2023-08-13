package helpers

type HttpError struct {
	Code    int
	Message string
}

func (e *HttpError) Error() string {
	return e.Message
}

func NewHttpError(code int, message string) *HttpError {
	return &HttpError{
		Code:    code,
		Message: message,
	}
}

func JsonMessage(body string) interface{} {
	return struct {
		Message string `json:"message"`
	}{
		Message: body,
	}
}
