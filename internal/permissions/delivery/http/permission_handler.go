package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zeross/project-demo/pkg/response"
)

func (h handler) CreatePermission(c *gin.Context) {
	ctx := c.Request.Context()
	req, err := h.processCreatePermissionRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.CreatePermission.processCreatePermissionRequest: %v", err)
		response.Error(c, err)
		return
	}
	p, err := h.uc.CreatePermission(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.CreatePermission.CreatePermission: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, h.newPermissionResp(p))
}

func (h handler) DetailPermission(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := h.processDetailPermissionRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.DetailPermission.processDetailPermissionRequest: %v", err)
		response.Error(c, err)
		return
	}
	p, err := h.uc.DetailPermission(ctx, id)
	if err != nil {
		h.l.Errorf(ctx, "http.DetailPermission.DetailPermission: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, h.newPermissionResp(p))
}

func (h handler) ListPermissions(c *gin.Context) {
	ctx := c.Request.Context()
	req, err := h.processListPermissionsRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.ListPermissions.processListPermissionsRequest: %v", err)
		response.Error(c, err)
		return
	}
	list, err := h.uc.ListPermissions(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.ListPermissions.ListPermissions: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, h.newListPermissionResp(list))
}

func (h handler) UpdatePermission(c *gin.Context) {
	ctx := c.Request.Context()
	req, err := h.processUpdatePermissionRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.UpdatePermission.processUpdatePermissionRequest: %v", err)
		response.Error(c, err)
		return
	}
	if err := h.uc.UpdatePermission(ctx, req.toInput()); err != nil {
		h.l.Errorf(ctx, "http.UpdatePermission.UpdatePermission: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, nil)
}

func (h handler) DeletePermission(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := h.processDeletePermissionRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.DeletePermission.processDeletePermissionRequest: %v", err)
		response.Error(c, err)
		return
	}
	if err := h.uc.DeletePermission(ctx, id); err != nil {
		h.l.Errorf(ctx, "http.DeletePermission.DeletePermission: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, nil)
}
