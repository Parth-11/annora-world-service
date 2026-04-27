package auth

import "errors"

var (
	ErrMissingSub     = errors.New("missing sub")
	ErrTokenExpired    = errors.New("token expired")
	ErrInvalidIssuer   = errors.New("invalid issuer")
	ErrInvalidAudience = errors.New("invalid audience")
)