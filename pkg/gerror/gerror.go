package gerror

import (
	"github.com/pkg/errors"
)

const (
	CodeNil = -1
)

type MyError struct {
	err  error
	code int
}

func (e MyError) Error() string {
	return e.err.Error()
}

func Code(err error) int {

	if e := new(MyError); errors.As(err, &e) {
		return e.code
	}

	return CodeNil
}
func Stack(err error) error {
	if se := new(MyError); errors.As(err, &se) {
		return se.err
	}

	return err
}

func New(err error) error {
	return &MyError{
		code: CodeNil,
		err:  err,
	}
}

func NewWithCode(code int, err error) error {
	return &MyError{
		code: code,
		err:  err,
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

func Failure(msg ...string) error {
	m := ResponseMsg.Failure
	if len(msg) > 0 {
		m = msg[0]
	}
	return NewWithCode(ResponseCode.Failure, errors.New(m))
}
