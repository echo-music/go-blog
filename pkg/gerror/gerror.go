package gerror

import (
	"errors"
)

const (
	CodeNil = -1
)

type MyError struct {
	error error
	code  int
}

func (e MyError) Error() string {
	if e.error != nil {
		return e.error.Error()
	}
	return ""
}

func Code(err error) int {
	if e, ok := err.(MyError); ok {
		return e.code
	}
	return CodeNil
}

func New(err error) error {
	return MyError{
		code:  CodeNil,
		error: err,
	}
}

func NewWithCode(code int, e error) error {
	return MyError{
		code:  code,
		error: e,
	}
}

func Exception(msg ...string) error {
	m := ResponseMsg.Exception
	if len(msg) > 0 {
		m = msg[0]
	}
	return NewWithCode(ResponseCode.Exception, errors.New(m))
}

func Unauthorized(msg ...string) error {
	m := ResponseMsg.Unauthorized
	if len(msg) > 0 {
		m = msg[0]
	}
	return NewWithCode(ResponseCode.Unauthorized, errors.New(m))
}

func Forbidden(msg ...string) error {
	m := ResponseMsg.Forbidden
	if len(msg) > 0 {
		m = msg[0]
	}
	return NewWithCode(ResponseCode.Forbidden, errors.New(m))
}

func TooManyRequests(msg ...string) error {
	m := ResponseMsg.TooManyRequests
	if len(msg) > 0 {
		m = msg[0]
	}
	return NewWithCode(ResponseCode.TooManyRequests, errors.New(m))
}
