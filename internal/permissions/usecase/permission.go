package usecase

import (
	"context"

	"github.com/zeross/project-demo/internal/models"
	"github.com/zeross/project-demo/internal/permissions"
	"github.com/zeross/project-demo/internal/permissions/repository"
)

func (u impleUsecase) CreatePermission(ctx context.Context, input permissions.CreateInput) (models.Permission, error) {
	p, err := u.repo.CreatePermission(ctx, repository.CreateOptions{Name: input.Name})
	if err != nil {
		u.l.Errorf(ctx, "usecase.CreatePermission.CreatePermission: %v", err)
		return models.Permission{}, err
	}
	return p, nil
}

func (u impleUsecase) DetailPermission(ctx context.Context, id string) (models.Permission, error) {
	p, err := u.repo.DetailPermission(ctx, id)
	if err != nil {
		u.l.Errorf(ctx, "usecase.DetailPermission.DetailPermission: %v", err)
		return models.Permission{}, permissions.ErrPermissionNotFound
	}
	return p, nil
}

func (u impleUsecase) ListPermissions(ctx context.Context, input permissions.ListInput) ([]models.Permission, error) {
	list, err := u.repo.ListPermissions(ctx, repository.ListOptions{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		u.l.Errorf(ctx, "usecase.ListPermissions.ListPermissions: %v", err)
		return nil, err
	}
	return list, nil
}

func (u impleUsecase) UpdatePermission(ctx context.Context, input permissions.UpdateInput) error {
	current, err := u.repo.DetailPermission(ctx, input.ID)
	if err != nil {
		u.l.Errorf(ctx, "usecase.UpdatePermission.DetailPermission: %v", err)
		return permissions.ErrPermissionNotFound
	}
	name := current.Name
	if input.Name != nil {
		name = *input.Name
	}
	err = u.repo.UpdatePermission(ctx, repository.UpdateOptions{
		ID:   input.ID,
		Name: name,
	})
	if err != nil {
		u.l.Errorf(ctx, "usecase.UpdatePermission.UpdatePermission: %v", err)
		return err
	}
	return nil
}

func (u impleUsecase) DeletePermission(ctx context.Context, id string) error {
	if _, err := u.repo.DetailPermission(ctx, id); err != nil {
		u.l.Errorf(ctx, "usecase.DeletePermission.DetailPermission: %v", err)
		return permissions.ErrPermissionNotFound
	}
	if err := u.repo.DeletePermission(ctx, id); err != nil {
		u.l.Errorf(ctx, "usecase.DeletePermission.DeletePermission: %v", err)
		return err
	}
	return nil
}
