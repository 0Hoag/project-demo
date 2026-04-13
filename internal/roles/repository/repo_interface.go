package repository

import (
	"context"

	"github.com/zeross/project-demo/internal/models"
)

type Repository interface {
	CreateRole(ctx context.Context, opts CreateOptions) (models.Role, error)
	DetailRole(ctx context.Context, id string) (models.Role, error)
	GetRoleByName(ctx context.Context, name string) (models.Role, error)
	ListRoles(ctx context.Context, opts ListOptions) ([]models.Role, error)
	UpdateRole(ctx context.Context, opts UpdateOptions) error
	DeleteRole(ctx context.Context, id string) error

	AttachPermissionToRole(ctx context.Context, roleID, permissionID string) error
	DetachPermissionFromRole(ctx context.Context, roleID, permissionID string) error
	ListPermissionsForRole(ctx context.Context, roleID string) ([]models.Permission, error)
}
