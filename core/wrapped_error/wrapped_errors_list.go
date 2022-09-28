package wrapped_error

import "net/http"

func NewInternalServerError(err error, msg string) error {
	return New(http.StatusInternalServerError, err, msg)
}

func NewBadRequestError(err error, msg string) error {
	return New(http.StatusBadRequest, err, msg)
}

func NewNotFoundError(err error, msg string) error {
	return New(http.StatusNotFound, err, msg)
}
