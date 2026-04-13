package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zeross/project-demo/internal/permissions"
	"github.com/zeross/project-demo/pkg/log"
)

type Handler interface {
	CreatePermission(c *gin.Context)
	DetailPermission(c *gin.Context)
	ListPermissions(c *gin.Context)
	UpdatePermission(c *gin.Context)
	DeletePermission(c *gin.Context)
}

type handler struct {
	l  log.Logger
	uc permissions.Usecase
}

func New(l log.Logger, uc permissions.Usecase) Handler {
	return handler{l: l, uc: uc}
}
