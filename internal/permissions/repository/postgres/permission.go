package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zeross/project-demo/internal/models"
	"github.com/zeross/project-demo/internal/permissions/repository"
	sqlc "github.com/zeross/project-demo/internal/permissions/repository/sqlc"
)

func (r impleRepository) CreatePermission(ctx context.Context, opts repository.CreateOptions) (models.Permission, error) {
	q := sqlc.New(r.db)
	row, err := q.CreatePermission(ctx, opts.Name)
	if err != nil {
		r.l.Errorf(ctx, "postgres.permissions.CreatePermission: %v", err)
		return models.Permission{}, err
	}
	return mapPermission(row), nil
}

func (r impleRepository) DetailPermission(ctx context.Context, id string) (models.Permission, error) {
	q := sqlc.New(r.db)
	uid := pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true}
	row, err := q.GetPermissionByID(ctx, uid)
	if err != nil {
		r.l.Errorf(ctx, "postgres.permissions.GetPermissionByID: %v", err)
		return models.Permission{}, err
	}
	return mapPermission(row), nil
}

func (r impleRepository) ListPermissions(ctx context.Context, opts repository.ListOptions) ([]models.Permission, error) {
	q := sqlc.New(r.db)
	rows, err := q.ListPermissions(ctx, sqlc.ListPermissionsParams{
		Limit:  opts.Limit,
		Offset: opts.Offset,
	})
	if err != nil {
		r.l.Errorf(ctx, "postgres.permissions.ListPermissions: %v", err)
		return nil, err
	}
	out := make([]models.Permission, 0, len(rows))
	for _, row := range rows {
		out = append(out, mapPermission(row))
	}
	return out, nil
}

func (r impleRepository) UpdatePermission(ctx context.Context, opts repository.UpdateOptions) error {
	q := sqlc.New(r.db)
	uid := pgtype.UUID{Bytes: uuid.MustParse(opts.ID), Valid: true}
	_, err := q.UpdatePermission(ctx, sqlc.UpdatePermissionParams{
		ID:   uid,
		Name: opts.Name,
	})
	if err != nil {
		r.l.Errorf(ctx, "postgres.permissions.UpdatePermission: %v", err)
		return err
	}
	return nil
}

func (r impleRepository) DeletePermission(ctx context.Context, id string) error {
	q := sqlc.New(r.db)
	uid := pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true}
	if err := q.DeletePermission(ctx, uid); err != nil {
		r.l.Errorf(ctx, "postgres.permissions.DeletePermission: %v", err)
		return err
	}
	return nil
}

func mapPermission(row sqlc.Permission) models.Permission {
	return models.Permission{
		ID:        uuid.UUID(row.ID.Bytes),
		Name:      row.Name,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
}
