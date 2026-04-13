package http

import "github.com/gin-gonic/gin"

func MapRoutes(r *gin.RouterGroup, h Handler) {
	r.POST("", h.CreatePermission)
	r.GET("/all", h.ListPermissions)
	r.GET("/:id", h.DetailPermission)
	r.PATCH("", h.UpdatePermission)
	r.DELETE("/:id", h.DeletePermission)
}
