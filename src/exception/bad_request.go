package exception

type BadRequestError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}

func NewBadRequestError(err string) *BadRequestError {
	return &BadRequestError{
		Message: err,
	}
}
