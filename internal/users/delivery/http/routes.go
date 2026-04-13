package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zeross/project-demo/internal/middleware"
)

func MapRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.POST("", h.CreateUser)

	authGroup := r.Group("")
	authGroup.Use(mw.Auth())
	{
		authGroup.GET("/all", h.ListUsers)
		authGroup.GET("/:id", h.DetailUser)
		authGroup.PATCH("", h.UpdateUser)
		authGroup.DELETE("/:id", h.DeleteUser)
	}
}
