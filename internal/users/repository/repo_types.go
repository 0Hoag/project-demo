package repository

import (
	"time"

	"github.com/zeross/project-demo/internal/models"
	"github.com/zeross/project-demo/pkg/paginator"
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

type Filter struct {
	ID       string
	Username string
	Phone    string
}

type ListOptions struct {
	Filter
}

type GetUsersOptions struct {
	Filter
	PagQuery paginator.PaginatorQuery
}
