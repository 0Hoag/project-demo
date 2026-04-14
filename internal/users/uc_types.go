package users

import (
	"time"

	"github.com/zeross/project-demo/pkg/paginator"
)

type CreateInput struct {
	Username  string     `json:"username"`
	Phone     string     `json:"phone"`
	Password  string     `json:"password"`
	AvatarUrl string     `json:"avatar_url"`
	Bio       string     `json:"bio"`
	Birthday  *time.Time `json:"birthday"`
}

type UpdateInput struct {
	ID        string     `json:"id"`
	Password  *string    `json:"password"`
	AvatarUrl *string    `json:"avatar_url"`
	Bio       *string    `json:"bio"`
	Birthday  *time.Time `json:"birthday"`
}

type Filter struct {
	ID       string
	Username string
	Phone    string
}

type ListInput struct {
	Filter
}

type GetInput struct {
	Filter
	PagQuery paginator.PaginatorQuery
}
