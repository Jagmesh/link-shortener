package apperror

import "net/http"

// type AppError interface {

// }

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func New(message string, code int) *AppError {
	return &AppError{Message: message, Code: code}
}

func (e *AppError) Error() string {
	return e.Message
}

/** 4xx */
func BadRequest(msg string) *AppError {
	return New(msg, http.StatusBadRequest)
}

func Unauthorized(msg string) *AppError {
	return New(msg, http.StatusUnauthorized)
}

func Forbidden(msg string) *AppError {
	return New(msg, http.StatusForbidden)
}

func NotFound(msg string) *AppError {
	return New(msg, http.StatusNotFound)
}

func Conflict(msg string) *AppError {
	return New(msg, http.StatusConflict)
}

/** 5xx */
func Internal(msg string) *AppError {
	return New(msg, http.StatusInternalServerError)
}
