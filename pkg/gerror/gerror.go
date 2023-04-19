package gerror

import "github.com/pkg/errors"

const (
	CodeNil = -1
)

type MyError struct {
	msg  string
	code int
}

func (e MyError) Error() string {
	return e.msg
}

func Code(err error) int {

	if se := new(MyError); errors.As(err, &se) {
		return se.code
	}

	return CodeNil
}

func New(err string) error {
	return MyError{
		code: CodeNil,
		msg:  err,
	}
}

func NewWithCode(code int, message string) error {
	return MyError{
		code: code,
		msg:  message,
	}
}

func Exception(msg ...string) error {
	m := ResponseMsg.Exception
	if len(msg) > 0 {
		m = msg[0]
	}
	return NewWithCode(ResponseCode.Exception, m)
}

func Unauthorized(msg ...string) error {
	m := ResponseMsg.Unauthorized
	if len(msg) > 0 {
		m = msg[0]
	}
	return NewWithCode(ResponseCode.Unauthorized, m)
}

func Forbidden(msg ...string) error {
	m := ResponseMsg.Forbidden
	if len(msg) > 0 {
		m = msg[0]
	}
	return NewWithCode(ResponseCode.Forbidden, m)
}

func TooManyRequests(msg ...string) error {
	m := ResponseMsg.TooManyRequests
	if len(msg) > 0 {
		m = msg[0]
	}
	return NewWithCode(ResponseCode.TooManyRequests, m)
}

func Failure(msg ...string) error {
	m := ResponseMsg.Failure
	if len(msg) > 0 {
		m = msg[0]
	}
	return NewWithCode(ResponseCode.Failure, m)
}
