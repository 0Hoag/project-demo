package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h handler) processCreatePermissionRequest(c *gin.Context) (createPermissionReq, error) {
	ctx := c.Request.Context()
	var req createPermissionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "permissions.delivery.http.processCreatePermissionRequest.ShouldBindJSON: %v", err)
		return createPermissionReq{}, errWrongBody
	}
	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "permissions.delivery.http.processCreatePermissionRequest.validate: %v", err)
		return createPermissionReq{}, errWrongBody
	}
	return req, nil
}

func (h handler) processDetailPermissionRequest(c *gin.Context) (string, error) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		h.l.Errorf(ctx, "permissions.delivery.http.processDetailPermissionRequest.Validate: %v", err)
		return "", errWrongBody
	}
	return id, nil
}

func (h handler) processListPermissionsRequest(c *gin.Context) (listPermissionsReq, error) {
	ctx := c.Request.Context()
	var req listPermissionsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Errorf(ctx, "permissions.delivery.http.processListPermissionsRequest.ShouldBindQuery: %v", err)
		return listPermissionsReq{}, errWrongQuery
	}
	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "permissions.delivery.http.processListPermissionsRequest.validate: %v", err)
		return listPermissionsReq{}, errWrongBody
	}
	return req, nil
}

func (h handler) processUpdatePermissionRequest(c *gin.Context) (updatePermissionReq, error) {
	ctx := c.Request.Context()
	var req updatePermissionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "permissions.delivery.http.processUpdatePermissionRequest.ShouldBindJSON: %v", err)
		return updatePermissionReq{}, errWrongBody
	}
	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "permissions.delivery.http.processUpdatePermissionRequest.validate: %v", err)
		return updatePermissionReq{}, errWrongBody
	}
	return req, nil
}

func (h handler) processDeletePermissionRequest(c *gin.Context) (string, error) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		h.l.Errorf(ctx, "permissions.delivery.http.processDeletePermissionRequest.Validate: %v", err)
		return "", errWrongBody
	}
	return id, nil
}
