package postgres

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zeross/project-demo/internal/models"
	"github.com/zeross/project-demo/internal/roles"
	"github.com/zeross/project-demo/internal/roles/repository"
	sqlc "github.com/zeross/project-demo/internal/roles/repository/sqlc"
)

func (r impleRepository) CreateRole(ctx context.Context, opts repository.CreateOptions) (models.Role, error) {
	q := sqlc.New(r.db)
	row, err := q.CreateRole(ctx, opts.Name)
	if err != nil {
		r.l.Errorf(ctx, "postgres.roles.CreateRole: %v", err)
		return models.Role{}, err
	}
	return mapRole(row), nil
}

func (r impleRepository) DetailRole(ctx context.Context, id string) (models.Role, error) {
	q := sqlc.New(r.db)
	uid := pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true}
	row, err := q.GetRoleByID(ctx, uid)
	if err != nil {
		r.l.Errorf(ctx, "postgres.roles.GetRoleByID: %v", err)
		return models.Role{}, err
	}
	return mapRole(row), nil
}

func (r impleRepository) GetRoleByName(ctx context.Context, name string) (models.Role, error) {
	q := sqlc.New(r.db)
	row, err := q.GetRoleByName(ctx, name)
	if err != nil {
		if err == pgx.ErrNoRows {
			r.l.Errorf(ctx, "postgres.roles.GetRoleByID.ErrNoRows: %v", err)
			return models.Role{}, roles.ErrRoleNotFound
		}
		r.l.Errorf(ctx, "postgres.roles.GetRoleByName: %v", err)
		return models.Role{}, err
	}
	return mapRole(row), nil
}

func (r impleRepository) ListRoles(ctx context.Context, opts repository.ListOptions) ([]models.Role, error) {
	q := sqlc.New(r.db)
	rows, err := q.ListRoles(ctx, sqlc.ListRolesParams{
		Limit:  opts.Limit,
		Offset: opts.Offset,
	})
	if err != nil {
		r.l.Errorf(ctx, "postgres.roles.ListRoles: %v", err)
		return nil, err
	}
	out := make([]models.Role, 0, len(rows))
	for _, row := range rows {
		out = append(out, mapRole(row))
	}
	return out, nil
}

func (r impleRepository) UpdateRole(ctx context.Context, opts repository.UpdateOptions) error {
	q := sqlc.New(r.db)
	uid := pgtype.UUID{Bytes: uuid.MustParse(opts.ID), Valid: true}
	_, err := q.UpdateRole(ctx, sqlc.UpdateRoleParams{
		ID:   uid,
		Name: opts.Name,
	})
	if err != nil {
		r.l.Errorf(ctx, "postgres.roles.UpdateRole: %v", err)
		return err
	}
	return nil
}

func (r impleRepository) DeleteRole(ctx context.Context, id string) error {
	q := sqlc.New(r.db)
	uid := pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true}
	if err := q.DeleteRole(ctx, uid); err != nil {
		r.l.Errorf(ctx, "postgres.roles.DeleteRole: %v", err)
		return err
	}
	return nil
}

func (r impleRepository) AttachPermissionToRole(ctx context.Context, roleID, permissionID string) error {
	q := sqlc.New(r.db)
	rid := pgtype.UUID{Bytes: uuid.MustParse(roleID), Valid: true}
	pid := pgtype.UUID{Bytes: uuid.MustParse(permissionID), Valid: true}
	err := q.AttachPermissionToRole(ctx, sqlc.AttachPermissionToRoleParams{
		RoleID:       rid,
		PermissionID: pid,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
			return roles.ErrPermissionAlreadyLinked
		}
		r.l.Errorf(ctx, "postgres.roles.AttachPermissionToRole: %v", err)
		return err
	}
	return nil
}

func (r impleRepository) DetachPermissionFromRole(ctx context.Context, roleID, permissionID string) error {
	q := sqlc.New(r.db)
	rid := pgtype.UUID{Bytes: uuid.MustParse(roleID), Valid: true}
	pid := pgtype.UUID{Bytes: uuid.MustParse(permissionID), Valid: true}
	if err := q.DetachPermissionFromRole(ctx, sqlc.DetachPermissionFromRoleParams{
		RoleID:       rid,
		PermissionID: pid,
	}); err != nil {
		r.l.Errorf(ctx, "postgres.roles.DetachPermissionFromRole: %v", err)
		return err
	}
	return nil
}

func (r impleRepository) ListPermissionsForRole(ctx context.Context, roleID string) ([]models.Permission, error) {
	q := sqlc.New(r.db)
	rid := pgtype.UUID{Bytes: uuid.MustParse(roleID), Valid: true}
	rows, err := q.ListPermissionsForRole(ctx, rid)
	if err != nil {
		r.l.Errorf(ctx, "postgres.roles.ListPermissionsForRole: %v", err)
		return nil, err
	}
	out := make([]models.Permission, 0, len(rows))
	for _, row := range rows {
		out = append(out, mapPermission(row))
	}
	return out, nil
}

func mapPermission(row sqlc.Permission) models.Permission {
	return models.Permission{
		ID:        uuid.UUID(row.ID.Bytes),
		Name:      row.Name,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
}

func mapRole(row sqlc.Role) models.Role {
	return models.Role{
		ID:        uuid.UUID(row.ID.Bytes),
		Name:      row.Name,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
}
