package usecase

import (
	"github.com/zeross/project-demo/internal/permissions"
	"github.com/zeross/project-demo/internal/permissions/repository"
	"github.com/zeross/project-demo/pkg/log"
)

type impleUsecase struct {
	l    log.Logger
	repo repository.Repository
}

func New(l log.Logger, repo repository.Repository) permissions.Usecase {
	return &impleUsecase{l: l, repo: repo}
}
