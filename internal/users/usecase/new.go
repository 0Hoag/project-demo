package usecase

import (
	rolesrepo "github.com/zeross/project-demo/internal/roles/repository"
	"github.com/zeross/project-demo/internal/users"
	"github.com/zeross/project-demo/internal/users/repository"
	"github.com/zeross/project-demo/pkg/log"
)

type impleUsecase struct {
	l        log.Logger
	repo     repository.Repository
	roleRepo rolesrepo.Repository
	db       repository.DBPool
}

func New(
	l log.Logger,
	repo repository.Repository,
	roleRepo rolesrepo.Repository,
	db repository.DBPool,
) users.Usecase {
	return &impleUsecase{
		l:        l,
		repo:     repo,
		roleRepo: roleRepo,
		db:       db,
	}
}
