package auth

import "context"

//go:generate mockery --name=Usecase
type Usecase interface {
	Login(ctx context.Context, input LoginInput) (LoginResponse, error)
}
