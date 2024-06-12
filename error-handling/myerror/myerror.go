package myerror

import "fmt"

// Error Invalid argument
type ErrorInvalidArgument interface {
	Error() string
}

type errorInvalidArgument struct {
	msg string
}

func (e errorInvalidArgument) Error() string {
	return fmt.Sprintf("[Invalid Argument]: %s", e.msg)
}

func NewErrorInvalidArgument(msg string) ErrorInvalidArgument {
	return errorInvalidArgument{msg: msg}
}

// Error Not Found
type ErrorNotFound interface {
	Error() string
}

type errorNotFound struct {
	msg string
}

func (e errorNotFound) Error() string {
	return fmt.Sprintf("[Not Found]: %s", e.msg)
}

func NewErrorNotFound(msg string) ErrorNotFound {
	return errorNotFound{msg: msg}
}
