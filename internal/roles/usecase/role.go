package usecase

import (
	"context"

	"github.com/zeross/project-demo/internal/models"
	"github.com/zeross/project-demo/internal/roles"
	"github.com/zeross/project-demo/internal/roles/repository"
)

func (u impleUsecase) CreateRole(ctx context.Context, input roles.CreateInput) (models.Role, error) {
	role, err := u.repo.CreateRole(ctx, repository.CreateOptions{Name: input.Name})
	if err != nil {
		u.l.Errorf(ctx, "usecase.CreateRole.CreateRole: %v", err)
		return models.Role{}, err
	}
	return role, nil
}

func (u impleUsecase) DetailRole(ctx context.Context, id string) (models.Role, error) {
	role, err := u.repo.DetailRole(ctx, id)
	if err != nil {
		u.l.Errorf(ctx, "usecase.DetailRole.DetailRole: %v", err)
		return models.Role{}, roles.ErrRoleNotFound
	}
	return role, nil
}

func (u impleUsecase) ListRoles(ctx context.Context, input roles.ListInput) ([]models.Role, error) {
	list, err := u.repo.ListRoles(ctx, repository.ListOptions{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		u.l.Errorf(ctx, "usecase.ListRoles.ListRoles: %v", err)
		return nil, err
	}
	return list, nil
}

func (u impleUsecase) UpdateRole(ctx context.Context, input roles.UpdateInput) error {
	current, err := u.repo.DetailRole(ctx, input.ID)
	if err != nil {
		u.l.Errorf(ctx, "usecase.UpdateRole.DetailRole: %v", err)
		return roles.ErrRoleNotFound
	}
	name := current.Name
	if input.Name != nil {
		name = *input.Name
	}
	err = u.repo.UpdateRole(ctx, repository.UpdateOptions{
		ID:   input.ID,
		Name: name,
	})
	if err != nil {
		u.l.Errorf(ctx, "usecase.UpdateRole.UpdateRole: %v", err)
		return err
	}
	return nil
}

func (u impleUsecase) DeleteRole(ctx context.Context, id string) error {
	if _, err := u.repo.DetailRole(ctx, id); err != nil {
		u.l.Errorf(ctx, "usecase.DeleteRole.DetailRole: %v", err)
		return roles.ErrRoleNotFound
	}
	if err := u.repo.DeleteRole(ctx, id); err != nil {
		u.l.Errorf(ctx, "usecase.DeleteRole.DeleteRole: %v", err)
		return err
	}
	return nil
}
