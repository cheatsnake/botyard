package extlib

import "net/http"

type ExtendedError struct {
	Code    int
	Message string
}

func NewExtendedError(code int, message string) *ExtendedError {
	return &ExtendedError{
		Code:    code,
		Message: message,
	}
}

func (e *ExtendedError) Error() string {
	return e.Message
}

func ErrorBadRequest(message string) *ExtendedError {
	return &ExtendedError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func ErrorUnauthorized(message string) *ExtendedError {
	return &ExtendedError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func ErrorForbidden(message string) *ExtendedError {
	return &ExtendedError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func ErrorNotFound(message string) *ExtendedError {
	return &ExtendedError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
