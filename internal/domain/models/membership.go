package models

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleMember    Role = "member"
	RoleModerator Role = "moderator"
	RoleOwner     Role = "owner"
)

type Membership struct {
	WorldId    uuid.UUID `json:"world_id" db:"world_id"`
	UserId     uuid.UUID `json:"user_id" db:"user_id"`
	Role       Role      `json:"role" db:"role"`
	Reputation int       `json:"reputation" db:"reputation"`
	JoinedAt   time.Time `json:"joined_at" db:"joined_at"`
}
