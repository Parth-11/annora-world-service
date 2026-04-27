package middleware

import "errors"

var (
	ErrWorldNotFound          = errors.New("world not found")
	ErrWorldNameAlreadyExists = errors.New("world name alreay in use")
	ErrAuthMissingSUB         = errors.New("missing sub")
	ErrAuthTokenExp           = errors.New("token expired")
	ErrAuthInvalidIss         = errors.New("invalid issuer")
	ErrAuthInvalidAud         = errors.New("invalid audience")
)
