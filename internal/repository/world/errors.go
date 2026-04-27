package worldrepo

import "errors"

var (
	ErrWorldNotFound          = errors.New("world not found")
	ErrWorldNameAlreadyExists = errors.New("world name already in use")
)