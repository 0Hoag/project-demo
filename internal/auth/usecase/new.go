package usecase

import (
	"github.com/zeross/project-demo/internal/auth"
	userrepo "github.com/zeross/project-demo/internal/users/repository"
	"github.com/zeross/project-demo/pkg/jwt"
	"github.com/zeross/project-demo/pkg/log"
)

type impleUsecase struct {
	l        log.Logger
	jwt      jwt.Manager
	userRepo userrepo.Repository
}

func New(l log.Logger, j jwt.Manager, userRepo userrepo.Repository) auth.Usecase {
	return &impleUsecase{
		l:        l,
		jwt:      j,
		userRepo: userRepo,
	}
}
