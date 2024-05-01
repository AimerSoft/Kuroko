package errno

import "fmt"

type Errno struct {
	Code    int
	Message string
}

func (e Errno) Error() string {
	return e.Message
}

type Err struct {
	Code    int
	Message string
	Err     string
}

func (e Err) Error() string {
	return fmt.Sprintf("Err-code: %d, message:%s, error:%s", e.Code, e.Message, e.Err)
}

func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: err.Error()}
}

func DecodeErr(err error) (int, string) {
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Err
	case *Errno:
		return typed.Code, typed.Message
	default:

	}
	return 500, err.Error()
}
