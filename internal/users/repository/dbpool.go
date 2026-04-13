package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	sqlc "github.com/zeross/project-demo/internal/users/repository/sqlc"
)

// DBPool is the Postgres dependency for the users repository (*pgxpool.Pool satisfies it).
type DBPool interface {
	sqlc.DBTX
	Begin(ctx context.Context) (pgx.Tx, error)
}
