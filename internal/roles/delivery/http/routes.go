package http

import "github.com/gin-gonic/gin"

func MapRoutes(r *gin.RouterGroup, h Handler) {
	r.POST("", h.CreateRole)
	r.GET("/all", h.ListRoles)

	r.GET("/:id/permissions", h.ListRolePermissions)
	r.POST("/:id/permissions", h.AttachPermissionToRole)
	r.DELETE("/:id/permissions/:permission_id", h.DetachPermissionFromRole)

	r.GET("/:id", h.DetailRole)
	r.PATCH("", h.UpdateRole)
	r.DELETE("/:id", h.DeleteRole)
}
