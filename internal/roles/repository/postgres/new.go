package postgres

import (
	"github.com/zeross/project-demo/internal/roles/repository"
	sqlc "github.com/zeross/project-demo/internal/roles/repository/sqlc"
	"github.com/zeross/project-demo/pkg/log"
)

type impleRepository struct {
	l  log.Logger
	db sqlc.DBTX
}

func New(l log.Logger, db sqlc.DBTX) repository.Repository {
	return &impleRepository{l: l, db: db}
}
