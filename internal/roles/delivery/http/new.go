package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zeross/project-demo/internal/roles"
	"github.com/zeross/project-demo/pkg/log"
)

type Handler interface {
	CreateRole(c *gin.Context)
	DetailRole(c *gin.Context)
	ListRoles(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)

	AttachPermissionToRole(c *gin.Context)
	DetachPermissionFromRole(c *gin.Context)
	ListRolePermissions(c *gin.Context)
}

type handler struct {
	l  log.Logger
	uc roles.Usecase
}

func New(l log.Logger, uc roles.Usecase) Handler {
	return handler{l: l, uc: uc}
}
