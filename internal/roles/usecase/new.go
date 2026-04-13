package usecase

import (
	permrepo "github.com/zeross/project-demo/internal/permissions/repository"
	"github.com/zeross/project-demo/internal/roles"
	"github.com/zeross/project-demo/internal/roles/repository"
	"github.com/zeross/project-demo/pkg/log"
)

type impleUsecase struct {
	l         log.Logger
	repo      repository.Repository
	permRepo  permrepo.Repository
}

func New(l log.Logger, repo repository.Repository, permRepo permrepo.Repository) roles.Usecase {
	return &impleUsecase{l: l, repo: repo, permRepo: permRepo}
}
