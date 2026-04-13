package permissions

import (
	"context"

	"github.com/zeross/project-demo/internal/models"
)

type Usecase interface {
	CreatePermission(ctx context.Context, input CreateInput) (models.Permission, error)
	DetailPermission(ctx context.Context, id string) (models.Permission, error)
	ListPermissions(ctx context.Context, input ListInput) ([]models.Permission, error)
	UpdatePermission(ctx context.Context, input UpdateInput) error
	DeletePermission(ctx context.Context, id string) error
}
