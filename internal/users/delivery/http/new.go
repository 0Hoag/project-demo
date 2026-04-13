package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zeross/project-demo/internal/users"
	"github.com/zeross/project-demo/pkg/log"
)

type Handler interface {
	CreateUser(c *gin.Context)
	DetailUser(c *gin.Context)
	ListUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type handler struct {
	l  log.Logger
	uc users.Usecase
}

func New(l log.Logger, uc users.Usecase) Handler {
	return handler{
		l:  l,
		uc: uc,
	}
}
