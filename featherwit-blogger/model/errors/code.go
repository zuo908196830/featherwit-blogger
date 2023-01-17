package errors

type ErrorCode int

const (
	Unauthorized         ErrorCode = 1006
	ResourceNotExist     ErrorCode = 1001
	ResourceAlreadyExist ErrorCode = 1002
	BadRequest           ErrorCode = 1003
	PasswordWrong        ErrorCode = 1004
	TokenWrong           ErrorCode = 1005
)
