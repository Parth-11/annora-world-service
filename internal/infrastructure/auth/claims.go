package auth

import (
	"time"
)

type Claims struct {
	Sub string
	Exp int64
	Iss string
	Aud string
}

func (c *Claims) Validate(expectedIss, expectedAud string) error {
	if c.Sub == "" {
		return ErrMissingSub
	}

	if time.Now().Unix() > c.Exp {
		return ErrTokenExpired
	}

	if c.Iss != expectedIss {
		return ErrInvalidIssuer
	}

	return nil
}
