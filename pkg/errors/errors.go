package errors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
)

type Error struct {
	Err error
	Msg string
}

func (e *Error) Error() string {
	return e.Msg
}

func (e *Error) Unwrap() error {
	return e.Err
}

func Wrap(err error, msg string) error {
	return &Error{
		Err: err,
		Msg: fmt.Sprintf("%s: %v", msg, err),
	}
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func NotFound(msg string) error {
	return &Error{
		Err: ErrNotFound,
		Msg: msg,
	}
}
