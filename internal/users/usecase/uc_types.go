package usecase

import (
	"time"

	"github.com/google/uuid"
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
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Phone     string     `json:"phone"`
	Password  *string    `json:"password"`
	AvatarUrl *string    `json:"avatar_url"`
	Bio       *string    `json:"bio"`
	Birthday  *time.Time `json:"birthday"`
}

type ListInput struct {
	Pagination paginator.PaginatorQuery `json:"pagination"`
}
