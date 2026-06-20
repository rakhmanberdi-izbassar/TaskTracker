package apperrors

import "net/http"

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, msg string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func NotFound(msg string, err error) *AppError {
	return New(http.StatusNotFound, msg, err)
}

func BadRequest(msg string, err error) *AppError {
	return New(http.StatusBadRequest, msg, err)
}

func InternalServerError(msg string, err error) *AppError {
	return New(http.StatusInternalServerError, msg, err)
}

func Unauthorized(msg string, err error) *AppError {
	return New(http.StatusUnauthorized, msg, err)
}

func Forbidden(msg string, err error) *AppError {
	return New(http.StatusForbidden, msg, err)
}

func Conflict(msg string, err error) *AppError {
	return New(http.StatusConflict, msg, err)
}
