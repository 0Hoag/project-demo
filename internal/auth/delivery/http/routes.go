package http

import "github.com/gin-gonic/gin"

func MapRoutes(r *gin.RouterGroup, h Handler) {
	r.POST("/login", h.Login)
}
