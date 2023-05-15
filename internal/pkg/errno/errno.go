package errno

import "fmt"

type Errno struct {
	HTTPCode int
	Code     string
	Message  string
}

func (e *Errno) Error() string {
	return e.Message
}

func (e *Errno) SetMessage(format string, opts ...interface{}) *Errno {
	e.Message = fmt.Sprintf(format, opts...)
	return e
}

func Decode(e error) (int, string, string) {
	if e == nil {
		return HTTPOK.HTTPCode, HTTPOK.Code, HTTPOK.Message
	}

	switch typed := e.(type) {
	case *Errno:
		return typed.HTTPCode, typed.Code, typed.Message
	}

	return InternalServerError.HTTPCode, InternalServerError.Code, InternalServerError.Message
}
