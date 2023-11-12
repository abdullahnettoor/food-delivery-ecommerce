package e

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrConflict        = errors.New("already exist")
	ErrDb              = errors.New("db error")
	ErrInvalidPassword = errors.New("invalid password")
	ErrNotAvailable = errors.New("not available")
)
