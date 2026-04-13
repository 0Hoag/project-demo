package usecase

import (
	"context"

	"github.com/zeross/project-demo/internal/models"
	"github.com/zeross/project-demo/internal/permissions"
	"github.com/zeross/project-demo/internal/roles"
)

func (u impleUsecase) AttachPermissionToRole(ctx context.Context, input roles.AttachPermissionInput) error {
	if _, err := u.repo.DetailRole(ctx, input.RoleID); err != nil {
		u.l.Errorf(ctx, "usecase.AttachPermissionToRole.DetailRole: %v", err)
		return roles.ErrRoleNotFound
	}
	if _, err := u.permRepo.DetailPermission(ctx, input.PermissionID); err != nil {
		u.l.Errorf(ctx, "usecase.AttachPermissionToRole.DetailPermission: %v", err)
		return permissions.ErrPermissionNotFound
	}
	if err := u.repo.AttachPermissionToRole(ctx, input.RoleID, input.PermissionID); err != nil {
		u.l.Errorf(ctx, "usecase.AttachPermissionToRole.AttachPermissionToRole: %v", err)
		return err
	}
	return nil
}

func (u impleUsecase) DetachPermissionFromRole(ctx context.Context, roleID, permissionID string) error {
	if _, err := u.repo.DetailRole(ctx, roleID); err != nil {
		u.l.Errorf(ctx, "usecase.DetachPermissionFromRole.DetailRole: %v", err)
		return roles.ErrRoleNotFound
	}
	if _, err := u.permRepo.DetailPermission(ctx, permissionID); err != nil {
		u.l.Errorf(ctx, "usecase.DetachPermissionFromRole.DetailPermission: %v", err)
		return permissions.ErrPermissionNotFound
	}
	if err := u.repo.DetachPermissionFromRole(ctx, roleID, permissionID); err != nil {
		u.l.Errorf(ctx, "usecase.DetachPermissionFromRole.DetachPermissionFromRole: %v", err)
		return err
	}
	return nil
}

func (u impleUsecase) ListRolePermissions(ctx context.Context, roleID string) ([]models.Permission, error) {
	if _, err := u.repo.DetailRole(ctx, roleID); err != nil {
		u.l.Errorf(ctx, "usecase.ListRolePermissions.DetailRole: %v", err)
		return nil, roles.ErrRoleNotFound
	}
	list, err := u.repo.ListPermissionsForRole(ctx, roleID)
	if err != nil {
		u.l.Errorf(ctx, "usecase.ListRolePermissions.ListPermissionsForRole: %v", err)
		return nil, err
	}
	return list, nil
}
