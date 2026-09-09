package errors

import (
	"fmt"
	"net/http"
)

// AppError represents an application error
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new application error
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Predefined error constructors
func NewBadRequestError(message string, err error) *AppError {
	return NewAppError(http.StatusBadRequest, message, err)
}

func NewUnauthorizedError(message string, err error) *AppError {
	return NewAppError(http.StatusUnauthorized, message, err)
}

func NewForbiddenError(message string, err error) *AppError {
	return NewAppError(http.StatusForbidden, message, err)
}

func NewNotFoundError(message string, err error) *AppError {
	return NewAppError(http.StatusNotFound, message, err)
}

func NewInternalServerError(message string, err error) *AppError {
	return NewAppError(http.StatusInternalServerError, message, err)
}

func NewValidationError(message string, err error) *AppError {
	return NewAppError(http.StatusUnprocessableEntity, message, err)
}

