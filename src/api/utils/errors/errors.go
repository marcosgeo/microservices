package errors

import "net/http"

// APIError ...
type APIError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	status  int    `json:"status"`
	message string `json:"message"`
	error   string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.status
}

func (e *apiError) Message() string {
	return e.message
}

func (e *apiError) Error() string {
	return e.error
}

// NewApiError
func NewApiError(statusCode int, message string) APIError {
	return &apiError{status: statusCode, message: message}
}

// NewNotFoundError ...
func NewNotFoundError(message string) APIError {
	return &apiError{
		status:  http.StatusNotFound,
		message: message,
	}
}

// NewInternalServerError ...
func NewInternalServerError(message string) APIError {
	return &apiError{
		status:  http.StatusInternalServerError,
		message: message,
	}
}

// NewBadRequestError ...
func NewBadRequestError(message string) APIError {
	return &apiError{
		status:  http.StatusBadRequest,
		message: message,
	}
}
