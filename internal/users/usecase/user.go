package usecase

import (
	"context"
	"errors"

	"github.com/zeross/project-demo/internal/models"
	"github.com/zeross/project-demo/internal/roles"
	"github.com/zeross/project-demo/internal/users"
	"github.com/zeross/project-demo/internal/users/repository"
	"golang.org/x/crypto/bcrypt"
)

const defaultSignupRoleName = "user"

func (u impleUsecase) CreateUser(ctx context.Context, input users.CreateInput) (models.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		u.l.Errorf(ctx, "usecase.CreateUser.GenerateFromPassword: %v", err)
		return models.User{}, err
	}

	role, err := u.roleRepo.GetRoleByName(ctx, defaultSignupRoleName)
	if err != nil {
		if errors.Is(err, roles.ErrRoleNotFound) {
			u.l.Errorf(ctx, "usecase.CreateUser.GetRoleByName.ErrRoleNotFound: %v", err)
			return models.User{}, users.ErrDefaultRoleUnavailable
		}
		u.l.Errorf(ctx, "usecase.CreateUser.GetRoleByName: %v", err)
		return models.User{}, err
	}

	tx, err := u.db.Begin(ctx)
	if err != nil {
		u.l.Errorf(ctx, "usecase.CreateUser.Begin: %v", err)
		return models.User{}, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	user, err := u.repo.CreateUserInTx(ctx, tx, repository.CreateOptions{
		Username:     input.Username,
		Phone:        input.Phone,
		PasswordHash: string(passwordHash),
		AvatarUrl:    input.AvatarUrl,
		Bio:          input.Bio,
		Birthday:     input.Birthday,
	})
	if err != nil {
		u.l.Errorf(ctx, "usecase.CreateUser.CreateUserInTx: %v", err)
		return models.User{}, err
	}

	if err := u.repo.InsertUserRoleInTx(ctx, tx, user.ID.String(), role.ID.String()); err != nil {
		u.l.Errorf(ctx, "usecase.CreateUser.InsertUserRoleInTx: %v", err)
		return models.User{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		u.l.Errorf(ctx, "usecase.CreateUser.Commit: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (uc impleUsecase) GetSessionUser(ctx context.Context, sc models.Scope) (models.User, error) {
	u, err := uc.repo.DetailUser(ctx, sc, sc.UserID)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.GetSessionUser.Detail: %v", err)
		return models.User{}, users.ErrUserNotFound
	}

	return u, nil
}

func (u impleUsecase) DetailUser(ctx context.Context, sc models.Scope, id string) (models.User, error) {
	user, err := u.repo.DetailUser(ctx, sc, id)
	if err != nil {
		u.l.Errorf(ctx, "usecase.DetailUser.DetailUser: %v", err)
		return models.User{}, users.ErrUserNotFound
	}
	return user, nil
}

func (u impleUsecase) ListUsers(ctx context.Context, sc models.Scope, input users.ListInput) ([]models.User, error) {
	users, err := u.repo.ListUsers(ctx, sc, repository.ListOptions{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		u.l.Errorf(ctx, "usecase.ListUsers.ListUsers: %v", err)
		return nil, err
	}
	return users, nil
}

func (u impleUsecase) UpdateUser(ctx context.Context, sc models.Scope, input users.UpdateInput) error {
	current, err := u.repo.DetailUser(ctx, sc, input.ID)
	if err != nil {
		u.l.Errorf(ctx, "usecase.UpdateUser.DetailUser: %v", err)
		return users.ErrUserNotFound
	}

	var passwordHash *string
	if input.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		if err != nil {
			u.l.Errorf(ctx, "usecase.UpdateUser.GenerateFromPassword: %v", err)
			return err
		}
		s := string(hash)
		passwordHash = &s
	}

	err = u.repo.UpdateUser(ctx, sc, repository.UpdateOptions{
		User:         current,
		PasswordHash: passwordHash,
		AvatarUrl:    input.AvatarUrl,
		Bio:          input.Bio,
		Birthday:     input.Birthday,
	})
	if err != nil {
		u.l.Errorf(ctx, "usecase.UpdateUser.UpdateUser: %v", err)
		return err
	}
	return nil
}

func (u impleUsecase) DeleteUser(ctx context.Context, sc models.Scope, id string) error {
	if _, err := u.repo.DetailUser(ctx, sc, id); err != nil {
		u.l.Errorf(ctx, "usecase.DeleteUser.DetailUser: %v", err)
		return users.ErrUserNotFound
	}

	if err := u.repo.DeleteUser(ctx, sc, id); err != nil {
		u.l.Errorf(ctx, "usecase.DeleteUser.DeleteUser: %v", err)
		return err
	}

	return nil
}
