package apperror

type ErrorCode string

const (
	Unknown          ErrorCode = "unknown_error"
	Validation       ErrorCode = "validation_error"
	Database         ErrorCode = "database_error"
	Redis            ErrorCode = "redis_error"
	PermissionDenied ErrorCode = "permission_denied_error"
	Internal         ErrorCode = "internal_error"
	BadParams        ErrorCode = "bad_params_error"
)
