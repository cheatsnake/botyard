package extlib

import "net/http"

type ExtendedError {
	Code    int
	Message string
}

func NewExtendedError(code int, message string) *ExtendedError {
	return &HttpError{
		Code:    code,
		Message: message,
	}
}

func (e *ExtendedError) Error() string {
	return e.Message
}

func ErrorBadRequest(message string) *ExtendedError {
	return &HttpError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func ErrorUnauthorized(message string) *ExtendedError {
	return &HttpError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func ErrorForbidden(message string) *ExtendedError {
	return &HttpError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func ErrorNotFound(message string) *ExtendedError {
	return &HttpError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}