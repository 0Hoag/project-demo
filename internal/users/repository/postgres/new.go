package postgres

import (
	"github.com/zeross/project-demo/internal/users/repository"
	"github.com/zeross/project-demo/pkg/log"
)

type impleRepository struct {
	l  log.Logger
	db repository.DBPool
}

func New(
	l log.Logger,
	db repository.DBPool,
) repository.Repository {
	return &impleRepository{
		l:  l,
		db: db,
	}
}
