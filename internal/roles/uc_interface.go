package roles

import (
	"context"

	"github.com/zeross/project-demo/internal/models"
)

type Usecase interface {
	CreateRole(ctx context.Context, input CreateInput) (models.Role, error)
	DetailRole(ctx context.Context, id string) (models.Role, error)
	ListRoles(ctx context.Context, input ListInput) ([]models.Role, error)
	UpdateRole(ctx context.Context, input UpdateInput) error
	DeleteRole(ctx context.Context, id string) error

	AttachPermissionToRole(ctx context.Context, input AttachPermissionInput) error
	DetachPermissionFromRole(ctx context.Context, roleID, permissionID string) error
	ListRolePermissions(ctx context.Context, roleID string) ([]models.Permission, error)
}
