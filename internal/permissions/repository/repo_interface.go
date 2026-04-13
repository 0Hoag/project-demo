package repository

import (
	"context"

	"github.com/zeross/project-demo/internal/models"
)

type Repository interface {
	CreatePermission(ctx context.Context, opts CreateOptions) (models.Permission, error)
	DetailPermission(ctx context.Context, id string) (models.Permission, error)
	ListPermissions(ctx context.Context, opts ListOptions) ([]models.Permission, error)
	UpdatePermission(ctx context.Context, opts UpdateOptions) error
	DeletePermission(ctx context.Context, id string) error
}
