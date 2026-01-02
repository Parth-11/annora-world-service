package models

import (
	"time"

	"github.com/google/uuid"
)

type Visibility string

const (
	VisibilityPublic  Visibility = "public"
	VisibilityPrivate Visibility = "private"
)

func (v Visibility) isValid() bool {
	return v == VisibilityPrivate || v == VisibilityPublic
}

func (v Visibility) String() string {
	return string(v)
}

type WorldStatus string

const (
	WorldStatusActive   WorldStatus = "active"
	WorldStatusArchived WorldStatus = "archived"
)

type World struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description,omitempty" db:"description"`
	CreatorID   uuid.UUID `json:"creator_id" db:"creator_user_id"`

	Visibility  Visibility `json:"visibility" db:"visibility"`
	MemberCount int        `json:"member_count" db:"member_count"`
	CoverImg    string     `json:"cover_img,omitempty" db:"cover_image_url"`

	Status WorldStatus `json:"status" db:"status"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
