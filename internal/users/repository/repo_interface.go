package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/zeross/project-demo/internal/models"
)

//go:generate mockery --name=Repository
type Repository interface {
	CreateUser(ctx context.Context, opts CreateOptions) (models.User, error)
	CreateUserInTx(ctx context.Context, tx pgx.Tx, opts CreateOptions) (models.User, error)
	InsertUserRoleInTx(ctx context.Context, tx pgx.Tx, userID, roleID string) error
	DetailUser(ctx context.Context, sc models.Scope, id string) (models.User, error)
	GetUserByPhone(ctx context.Context, phone string) (models.User, error)
	ListRoleNamesByUserID(ctx context.Context, userID string) ([]string, error)
	ListPermissionNamesByUserID(ctx context.Context, userID string) ([]string, error)
	ListUsers(ctx context.Context, sc models.Scope, opts ListOptions) ([]models.User, error)
	UpdateUser(ctx context.Context, sc models.Scope, opts UpdateOptions) error
	DeleteUser(ctx context.Context, sc models.Scope, id string) error
}
