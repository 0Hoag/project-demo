package usecase

import (
	"context"

	"github.com/zeross/project-demo/internal/models"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, input CreateInput) (models.User, error)
	DetailUser(ctx context.Context, id string) (models.User, error)
	ListUsers(ctx context.Context, input ListInput) ([]models.User, error)
	UpdateUser(ctx context.Context, input UpdateInput) (models.User, error)
	DeleteUser(ctx context.Context, id string) error
}
