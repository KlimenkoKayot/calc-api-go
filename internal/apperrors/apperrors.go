package apperrors

import (
	"errors"
)

var (
	ErrBadRequest       = errors.New("invalid body decode")
	ErrMethodNotAllowed = errors.New("only post requests")
)
