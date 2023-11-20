package e

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrConflict        = errors.New("already exist")
	ErrIsEmpty        = errors.New("is empty")
	ErrDb              = errors.New("db error")
	ErrInvalidPassword = errors.New("invalid password")
	ErrNotAvailable    = errors.New("not available")
	ErrQuantityExceeds = errors.New("selected quantity not available")
)
