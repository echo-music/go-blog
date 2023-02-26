package gerror

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
