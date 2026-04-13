package usecase

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/zeross/project-demo/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

func (uc impleUsecase) Login(ctx context.Context, input auth.LoginInput) (auth.LoginResponse, error) {
	u, err := uc.userRepo.GetUserByPhone(ctx, input.Phone)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return auth.LoginResponse{}, auth.ErrInvalidCreds
		}
		uc.l.Errorf(ctx, "auth.usecase.Login.GetUserByPhone: %v", err)
		return auth.LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		return auth.LoginResponse{}, auth.ErrInvalidCreds
	}

	uid := u.ID.String()
	type namesResult struct {
		names []string
		err   error
	}
	rolesCh := make(chan namesResult, 1)
	permsCh := make(chan namesResult, 1)

	go func() {
		r, e := uc.userRepo.ListRoleNamesByUserID(ctx, uid)
		rolesCh <- namesResult{names: r, err: e}
	}()
	go func() {
		p, e := uc.userRepo.ListPermissionNamesByUserID(ctx, uid)
		permsCh <- namesResult{names: p, err: e}
	}()

	rr := <-rolesCh
	if rr.err != nil {
		uc.l.Errorf(ctx, "auth.usecase.Login.ListRoleNamesByUserID: %v", rr.err)
		return auth.LoginResponse{}, rr.err
	}
	rp := <-permsCh
	if rp.err != nil {
		uc.l.Errorf(ctx, "auth.usecase.Login.ListPermissionNamesByUserID: %v", rp.err)
		return auth.LoginResponse{}, rp.err
	}

	token, err := uc.jwt.Generate(uid, rr.names, rp.names)
	if err != nil {
		uc.l.Errorf(ctx, "auth.usecase.Login.Generate: %v", err)
		return auth.LoginResponse{}, err
	}

	return auth.LoginResponse{Token: token}, nil
}
