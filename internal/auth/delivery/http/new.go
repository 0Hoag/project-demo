package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zeross/project-demo/internal/auth"
	"github.com/zeross/project-demo/pkg/log"
)

type Handler interface {
	Login(c *gin.Context)
}

type handler struct {
	l  log.Logger
	uc auth.Usecase
}

func New(l log.Logger, uc auth.Usecase) Handler {
	return handler{l: l, uc: uc}
}
