package exception

type ConflictError struct {
	Message string
}

func (e *ConflictError) Error() string {
	return e.Message
}

func NewConflictError(err string) *ConflictError {
	return &ConflictError{
		Message: err,
	}
}
