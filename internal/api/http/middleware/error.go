package middleware

import "errors"

var (
	ErrWorldNotFound          = errors.New("world not found")
	ErrWorldNameAlreadyExists = errors.New("world name alreay in use")
)
