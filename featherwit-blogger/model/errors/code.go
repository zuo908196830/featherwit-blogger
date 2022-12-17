package errors

type ErrorCode int

const (
	ResourceNotExist     ErrorCode = 1001
	ResourceAlreadyExist ErrorCode = 1002
)
