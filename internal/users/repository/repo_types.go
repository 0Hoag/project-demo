package repository

import (
	"time"

	"github.com/zeross/project-demo/internal/models"
)

type CreateOptions struct {
	Username     string
	Phone        string
	PasswordHash string
	AvatarUrl    string
	Bio          string
	Birthday     *time.Time
}

type UpdateOptions struct {
	User         models.User
	PasswordHash *string
	AvatarUrl    *string
	Bio          *string
	Birthday     *time.Time
}

type ListOptions struct {
	Limit  int32
	Offset int32
}
