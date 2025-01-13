package errors

import "fmt"

type Error struct {
	Code int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func New(code int, message string) *Error {
	return &Error{
		Code: code,
		Message: message,
	}
}