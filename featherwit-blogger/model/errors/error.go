package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	Code   ErrorCode
	Msg    string
	Detail interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("[MerakError] code: %d, msg: %s", e.Code, e.Msg)
}

var ErrorMap map[ErrorCode]string

func init() {
	ErrorMap = make(map[ErrorCode]string)
	{
		ErrorMap[Unauthorized] = "unauthorized"
		ErrorMap[ResourceNotExist] = "resource not exist"
		ErrorMap[ResourceAlreadyExist] = "resource already exist"
		ErrorMap[BadRequest] = "bad request"
	}
}

func NewError(code ErrorCode, detail interface{}) error {
	if msg, ok := ErrorMap[code]; ok {
		err := &Error{
			Code:   code,
			Msg:    msg,
			Detail: detail,
		}
		return err
	} else {
		return errors.New("server fail")
	}
}
