package users

import (
	"context"

	"github.com/zeross/project-demo/internal/models"
)

//go:generate mockery --name=Usecase
type Usecase interface {
	CreateUser(ctx context.Context, input CreateInput) (models.User, error)
	GetSessionUser(ctx context.Context, sc models.Scope) (models.User, error)
	DetailUser(ctx context.Context, sc models.Scope, id string) (models.User, error)
	ListUsers(ctx context.Context, sc models.Scope, input ListInput) ([]models.User, error)
	UpdateUser(ctx context.Context, sc models.Scope, input UpdateInput) error
	DeleteUser(ctx context.Context, sc models.Scope, id string) error
}
