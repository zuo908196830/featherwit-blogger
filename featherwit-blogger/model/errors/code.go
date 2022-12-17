package errors

type ErrorCode int

const (
	Unauthorized         ErrorCode = 403
	ResourceNotExist     ErrorCode = 1001
	ResourceAlreadyExist ErrorCode = 1002
)
