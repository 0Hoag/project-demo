package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h handler) processCreateRoleRequest(c *gin.Context) (createRoleReq, error) {
	ctx := c.Request.Context()
	var req createRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processCreateRoleRequest.ShouldBindJSON: %v", err)
		return createRoleReq{}, errWrongBody
	}
	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processCreateRoleRequest.validate: %v", err)
		return createRoleReq{}, errWrongBody
	}
	return req, nil
}

func (h handler) processDetailRoleRequest(c *gin.Context) (string, error) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processDetailRoleRequest.Validate: %v", err)
		return "", errWrongBody
	}
	return id, nil
}

func (h handler) processListRolesRequest(c *gin.Context) (listRolesReq, error) {
	ctx := c.Request.Context()
	var req listRolesReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processListRolesRequest.ShouldBindQuery: %v", err)
		return listRolesReq{}, errWrongQuery
	}
	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processListRolesRequest.validate: %v", err)
		return listRolesReq{}, errWrongBody
	}
	return req, nil
}

func (h handler) processUpdateRoleRequest(c *gin.Context) (updateRoleReq, error) {
	ctx := c.Request.Context()
	var req updateRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processUpdateRoleRequest.ShouldBindJSON: %v", err)
		return updateRoleReq{}, errWrongBody
	}
	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processUpdateRoleRequest.validate: %v", err)
		return updateRoleReq{}, errWrongBody
	}
	return req, nil
}

func (h handler) processDeleteRoleRequest(c *gin.Context) (string, error) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processDeleteRoleRequest.Validate: %v", err)
		return "", errWrongBody
	}
	return id, nil
}

func (h handler) processAttachPermissionToRoleRequest(c *gin.Context) (string, attachPermToRoleReq, error) {
	ctx := c.Request.Context()
	roleID, err := h.processRoleIDFromParam(c)
	if err != nil {
		return "", attachPermToRoleReq{}, err
	}
	var req attachPermToRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processAttachPermissionToRoleRequest.ShouldBindJSON: %v", err)
		return "", attachPermToRoleReq{}, errWrongBody
	}
	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processAttachPermissionToRoleRequest.validate: %v", err)
		return "", attachPermToRoleReq{}, errWrongBody
	}
	return roleID, req, nil
}

func (h handler) processDetachPermissionFromRoleRequest(c *gin.Context) (roleID string, permissionID string, err error) {
	ctx := c.Request.Context()
	roleID, err = h.processRoleIDFromParam(c)
	if err != nil {
		return "", "", err
	}
	permissionID = c.Param("permission_id")
	if err := uuid.Validate(permissionID); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processDetachPermissionFromRoleRequest.permission_id: %v", err)
		return "", "", errWrongBody
	}
	return roleID, permissionID, nil
}

func (h handler) processRoleIDFromParam(c *gin.Context) (string, error) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		h.l.Errorf(ctx, "roles.delivery.http.processRoleIDFromParam.Validate: %v", err)
		return "", errWrongBody
	}
	return id, nil
}
