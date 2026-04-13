package users

import (
	"time"
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

type ListInput struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}
