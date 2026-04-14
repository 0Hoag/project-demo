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
			uc.l.Errorf(ctx, "auth.usecase.Login.GetUserByPhon.ErrNoRowse: %v", err)
			return auth.LoginResponse{}, auth.ErrInvalidCreds
		}
		uc.l.Errorf(ctx, "auth.usecase.Login.GetUserByPhone: %v", err)
		return auth.LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		uc.l.Errorf(ctx, "auth.usecase.Login.CompareHashAndPassword: %v", err)
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
		select {
		case <-ctx.Done():
			return
		case rolesCh <- namesResult{names: r, err: e}:
		}
	}()
	go func() {
		p, e := uc.userRepo.ListPermissionNamesByUserID(ctx, uid)
		select {
		case <-ctx.Done():
			return
		case permsCh <- namesResult{names: p, err: e}:
		}
		permsCh <- namesResult{names: p, err: e}
	}()

	var rr, rp namesResult
	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			uc.l.Errorf(ctx, "auth.usecase.Login canceled early: %v", ctx.Err())
			return auth.LoginResponse{}, ctx.Err()

		case res := <-rolesCh:
			if res.err != nil {
				uc.l.Errorf(ctx, "auth.usecase.Login.ListRoleNamesByUserID: %v", res.err)
				return auth.LoginResponse{}, res.err
			}
			rr = res

		case res := <-permsCh:
			if res.err != nil {
				uc.l.Errorf(ctx, "auth.usecase.Login.ListPermissionNamesByUserID: %v", res.err)
				return auth.LoginResponse{}, res.err
			}
			rp = res
		}
	}

	token, err := uc.jwt.Generate(uid, rr.names, rp.names)
	if err != nil {
		uc.l.Errorf(ctx, "auth.usecase.Login.Generate: %v", err)
		return auth.LoginResponse{}, err
	}

	return auth.LoginResponse{Token: token}, nil
}
