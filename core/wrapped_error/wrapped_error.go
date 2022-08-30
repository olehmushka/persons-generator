package wrapped_error

type WrappedError struct {
	StatusCode int
	Err        error
	Msg        string
}

func New(statusCode int, err error, msg string) error {
	return &WrappedError{
		StatusCode: statusCode,
		Err:        err,
		Msg:        msg,
	}
}

func (we *WrappedError) Error() string {
	if we == nil {
		return ""
	}
	if we.Err == nil {
		return we.Msg
	}

	return we.Err.Error()
}
