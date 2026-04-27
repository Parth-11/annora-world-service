package auth

import (
	"time"

	"github.com/Parth-11/annora-world-service/internal/api/http/middleware"
)

type Claims struct {
	Sub string
	Exp int64
	Iss string
	Aud string
}

func (c *Claims) Validate(expectedIss, expectedAud string) error {
	if c.Sub == "" {
		return middleware.ErrAuthMissingSUB
	}

	if time.Now().Unix() > c.Exp {
		return middleware.ErrAuthTokenExp
	}

	if c.Iss != expectedIss {
		return middleware.ErrAuthInvalidIss
	}

	return nil
}
