package errors

import "net/http"

// APIError ...
type APIError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AError   string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.AStatus
}

func (e *apiError) Message() string {
	return e.AMessage
}

func (e *apiError) Error() string {
	return e.AError
}

// NewApiError
func NewApiError(statusCode int, message string) APIError {
	return &apiError{AStatus: statusCode, AMessage: message}
}

// NewNotFoundError ...
func NewNotFoundError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}

// NewInternalServerError ...
func NewInternalServerError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: message,
	}
}

// NewBadRequestError ...
func NewBadRequestError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: message,
	}
}
