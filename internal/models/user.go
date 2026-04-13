package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Phone     string
	Password  string
	AvatarUrl string
	Bio       string
	Birthday  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
