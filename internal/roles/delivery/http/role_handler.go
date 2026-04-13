package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zeross/project-demo/pkg/response"
)

func (h handler) CreateRole(c *gin.Context) {
	ctx := c.Request.Context()
	req, err := h.processCreateRoleRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.CreateRole.processCreateRoleRequest: %v", err)
		response.Error(c, err)
		return
	}
	role, err := h.uc.CreateRole(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.CreateRole.CreateRole: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, h.newRoleResp(role))
}

func (h handler) DetailRole(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := h.processDetailRoleRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.DetailRole.processDetailRoleRequest: %v", err)
		response.Error(c, err)
		return
	}
	role, err := h.uc.DetailRole(ctx, id)
	if err != nil {
		h.l.Errorf(ctx, "http.DetailRole.DetailRole: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, h.newRoleResp(role))
}

func (h handler) ListRoles(c *gin.Context) {
	ctx := c.Request.Context()
	req, err := h.processListRolesRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.ListRoles.processListRolesRequest: %v", err)
		response.Error(c, err)
		return
	}
	list, err := h.uc.ListRoles(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.ListRoles.ListRoles: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, h.newListRoleResp(list))
}

func (h handler) UpdateRole(c *gin.Context) {
	ctx := c.Request.Context()
	req, err := h.processUpdateRoleRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.UpdateRole.processUpdateRoleRequest: %v", err)
		response.Error(c, err)
		return
	}
	if err := h.uc.UpdateRole(ctx, req.toInput()); err != nil {
		h.l.Errorf(ctx, "http.UpdateRole.UpdateRole: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, nil)
}

func (h handler) DeleteRole(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := h.processDeleteRoleRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.DeleteRole.processDeleteRoleRequest: %v", err)
		response.Error(c, err)
		return
	}
	if err := h.uc.DeleteRole(ctx, id); err != nil {
		h.l.Errorf(ctx, "http.DeleteRole.DeleteRole: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, nil)
}

func (h handler) AttachPermissionToRole(c *gin.Context) {
	ctx := c.Request.Context()
	roleID, req, err := h.processAttachPermissionToRoleRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.AttachPermissionToRole.processAttachPermissionToRoleRequest: %v", err)
		response.Error(c, err)
		return
	}
	in := req.toAttachInput(roleID)
	if err := h.uc.AttachPermissionToRole(ctx, in); err != nil {
		h.l.Errorf(ctx, "http.AttachPermissionToRole.AttachPermissionToRole: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, nil)
}

func (h handler) DetachPermissionFromRole(c *gin.Context) {
	ctx := c.Request.Context()
	roleID, permissionID, err := h.processDetachPermissionFromRoleRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.DetachPermissionFromRole.processDetachPermissionFromRoleRequest: %v", err)
		response.Error(c, err)
		return
	}
	if err := h.uc.DetachPermissionFromRole(ctx, roleID, permissionID); err != nil {
		h.l.Errorf(ctx, "http.DetachPermissionFromRole.DetachPermissionFromRole: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, nil)
}

func (h handler) ListRolePermissions(c *gin.Context) {
	ctx := c.Request.Context()
	roleID, err := h.processRoleIDFromParam(c)
	if err != nil {
		h.l.Errorf(ctx, "http.ListRolePermissions.processRoleIDFromParam: %v", err)
		response.Error(c, err)
		return
	}
	list, err := h.uc.ListRolePermissions(ctx, roleID)
	if err != nil {
		h.l.Errorf(ctx, "http.ListRolePermissions.ListRolePermissions: %v", err)
		response.Error(c, h.mapError(err))
		return
	}
	response.OK(c, h.newListPermissionEmbedResp(list))
}
