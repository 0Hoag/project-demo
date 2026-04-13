package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zeross/project-demo/internal/models"
	"github.com/zeross/project-demo/internal/users/repository"
	sqlc "github.com/zeross/project-demo/internal/users/repository/sqlc"
)

func (r impleRepository) CreateUser(ctx context.Context, opts repository.CreateOptions) (models.User, error) {
	return r.createUser(ctx, r.db, opts)
}

func (r impleRepository) CreateUserInTx(ctx context.Context, tx pgx.Tx, opts repository.CreateOptions) (models.User, error) {
	return r.createUser(ctx, tx, opts)
}

func (r impleRepository) createUser(ctx context.Context, db sqlc.DBTX, opts repository.CreateOptions) (models.User, error) {
	q := sqlc.New(db)

	var bday pgtype.Date
	if opts.Birthday != nil {
		bday = pgtype.Date{Time: *opts.Birthday, Valid: true}
	}

	u, err := q.CreateUser(ctx, sqlc.CreateUserParams{
		Username:  opts.Username,
		Phone:     opts.Phone,
		Password:  opts.PasswordHash,
		AvatarUrl: opts.AvatarUrl,
		Bio:       opts.Bio,
		Birthday:  bday,
	})
	if err != nil {
		r.l.Errorf(ctx, "postgres.users.createUser: %v", err)
		return models.User{}, err
	}

	return mapUser(u), nil
}

func (r impleRepository) InsertUserRoleInTx(ctx context.Context, tx pgx.Tx, userID, roleID string) error {
	q := sqlc.New(tx)
	uid := pgtype.UUID{Bytes: uuid.MustParse(userID), Valid: true}
	rid := pgtype.UUID{Bytes: uuid.MustParse(roleID), Valid: true}
	if err := q.InsertUserRole(ctx, sqlc.InsertUserRoleParams{UserID: uid, RoleID: rid}); err != nil {
		r.l.Errorf(ctx, "postgres.users.InsertUserRoleInTx: %v", err)
		return err
	}
	return nil
}

func (r impleRepository) DetailUser(ctx context.Context, sc models.Scope, id string) (models.User, error) {
	q := sqlc.New(r.db)

	uid := pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true}
	u, err := q.DetailUser(ctx, uid)
	if err != nil {
		r.l.Errorf(ctx, "postgres.users.DetailUser: %v", err)
		return models.User{}, err
	}

	return mapUser(u), nil
}

func (r impleRepository) GetUserByPhone(ctx context.Context, phone string) (models.User, error) {
	q := sqlc.New(r.db)
	u, err := q.GetUserByPhone(ctx, phone)
	if err != nil {
		if err == pgx.ErrNoRows {
			r.l.Errorf(ctx, "postgres.users.DetailUse.ErrNoRows: %v", err)
			return models.User{}, err
		}
		r.l.Errorf(ctx, "postgres.users.GetUserByPhone: %v", err)
		return models.User{}, err
	}
	return mapUser(u), nil
}

func (r impleRepository) ListRoleNamesByUserID(ctx context.Context, userID string) ([]string, error) {
	q := sqlc.New(r.db)
	uid := pgtype.UUID{Bytes: uuid.MustParse(userID), Valid: true}
	names, err := q.ListRoleNamesForUser(ctx, uid)
	if err != nil {
		r.l.Errorf(ctx, "postgres.users.ListRoleNamesForUser: %v", err)
		return nil, err
	}
	return names, nil
}

func (r impleRepository) ListPermissionNamesByUserID(ctx context.Context, userID string) ([]string, error) {
	q := sqlc.New(r.db)
	uid := pgtype.UUID{Bytes: uuid.MustParse(userID), Valid: true}
	names, err := q.ListPermissionNamesForUser(ctx, uid)
	if err != nil {
		r.l.Errorf(ctx, "postgres.users.ListPermissionNamesForUser: %v", err)
		return nil, err
	}
	return names, nil
}

func (r impleRepository) ListUsers(ctx context.Context, sc models.Scope, opts repository.ListOptions) ([]models.User, error) {
	q := sqlc.New(r.db)

	users, err := q.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  opts.Limit,
		Offset: opts.Offset,
	})
	if err != nil {
		r.l.Errorf(ctx, "postgres.userrepo.ListUsers: %v", err)
		return nil, err
	}

	out := make([]models.User, 0, len(users))
	for _, u := range users {
		out = append(out, mapUser(u))
	}
	return out, nil
}

func (r impleRepository) UpdateUser(ctx context.Context, sc models.Scope, opts repository.UpdateOptions) error {
	q := sqlc.New(r.db)

	uid := pgtype.UUID{Bytes: uuid.MustParse(opts.User.ID.String()), Valid: true}

	password := opts.User.Password
	if opts.PasswordHash != nil {
		password = *opts.PasswordHash
	}
	avatar := opts.User.AvatarUrl
	if opts.AvatarUrl != nil {
		avatar = *opts.AvatarUrl
	}
	bio := opts.User.Bio
	if opts.Bio != nil {
		bio = *opts.Bio
	}
	var bday pgtype.Date
	if opts.Birthday != nil {
		bday = pgtype.Date{Time: *opts.Birthday, Valid: true}
	} else if opts.User.Birthday != nil {
		bday = pgtype.Date{Time: *opts.User.Birthday, Valid: true}
	}

	_, err := q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:        uid,
		Username:  opts.User.Username,
		Phone:     opts.User.Phone,
		Password:  password,
		AvatarUrl: avatar,
		Bio:       bio,
		Birthday:  bday,
	})
	if err != nil {
		r.l.Errorf(ctx, "postgres.users.UpdateUser: %v", err)
		return err
	}

	return nil
}

func (r impleRepository) DeleteUser(ctx context.Context, sc models.Scope, id string) error {
	q := sqlc.New(r.db)

	uid := pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true}

	if err := q.DeleteUser(ctx, uid); err != nil {
		r.l.Errorf(ctx, "postgres.users.DeleteUser: %v", err)
		return err
	}

	return nil
}

func mapUser(u sqlc.User) models.User {
	var birthday *time.Time
	if u.Birthday.Valid {
		t := u.Birthday.Time
		birthday = &t
	}

	return models.User{
		ID:        uuid.UUID(u.ID.Bytes),
		Username:  u.Username,
		Phone:     u.Phone,
		Password:  u.Password,
		AvatarUrl: u.AvatarUrl,
		Bio:       u.Bio,
		Birthday:  birthday,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}
}
