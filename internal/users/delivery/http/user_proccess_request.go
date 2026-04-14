package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zeross/project-demo/internal/models"
	pkgErrors "github.com/zeross/project-demo/pkg/errors"
	"github.com/zeross/project-demo/pkg/jwt"
	"github.com/zeross/project-demo/pkg/paginator"
)

func (h handler) processCreateRequest(c *gin.Context) (createReq, error) {
	ctx := c.Request.Context()

	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processCreateRequest.ShouldBindJSON: %v", err)
		return createReq{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processCreateRequest.validate: %v", err)
		return createReq{}, errWrongBody
	}

	return req, nil
}

func (h handler) processDetailUserRequest(c *gin.Context) (string, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "users.delivery.http.processDetailUserRequest.GetPayloadFromContext: unauthorized")
		return "", models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processDetailUserRequest.Validate: %v", err)
		return "", models.Scope{}, errWrongBody
	}

	sc := jwt.NewScope(payload)

	return id, sc, nil
}

func (h handler) processListUsersRequest(c *gin.Context) (listUsersReq, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "users.delivery.http.processListUsersRequest.GetPayloadFromContext: unauthorized")
		return listUsersReq{}, models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req listUsersReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processListUsersRequest.ShouldBindQuery: %v", err)
		return listUsersReq{}, models.Scope{}, errWrongQuery
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processListUsersRequest.ShouldBindQuery: %v", err)
		return listUsersReq{}, models.Scope{}, errWrongBody
	}

	sc := jwt.NewScope(payload)

	return req, sc, nil
}

func (h handler) processGetUsersRequest(c *gin.Context) (getUsersReq, paginator.PaginatorQuery, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "users.delivery.http.processGetUsersRequest.GetPayloadFromContext: unauthorized")
		return getUsersReq{}, paginator.PaginatorQuery{}, models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req getUsersReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processGetUsersRequest.ShouldBindQuery: %v", err)
		return getUsersReq{}, paginator.PaginatorQuery{}, models.Scope{}, errWrongQuery
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processGetUsersRequest.ShouldBindQuery: %v", err)
		return getUsersReq{}, paginator.PaginatorQuery{}, models.Scope{}, errWrongBody
	}

	var pq paginator.PaginatorQuery
	if err := c.ShouldBindQuery(&pq); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processGetUsersRequest.ShouldBindQuery: %v", errWrongQuery)
		return getUsersReq{}, paginator.PaginatorQuery{}, models.Scope{}, errWrongQuery
	}

	sc := jwt.NewScope(payload)

	return req, pq, sc, nil
}

func (h handler) processUpdateRequest(c *gin.Context) (updateUserReq, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "users.delivery.http.processUpdateRequest.GetPayloadFromContext: unauthorized")
		return updateUserReq{}, models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req updateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processUpdateRequest.ShouldBindJSON: %v", err)
		return updateUserReq{}, models.Scope{}, errWrongBody
	}

	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processUpdateRequest.ShouldBindUri: %v", err)
		return updateUserReq{}, models.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processUpdateRequest.Validate: %v", err)
		return updateUserReq{}, models.Scope{}, errWrongBody
	}

	sc := jwt.NewScope(payload)

	return req, sc, nil
}

func (h handler) processDeleteUserRequest(c *gin.Context) (string, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "users.delivery.http.processDeleteUserRequest.GetPayloadFromContext: unauthorized")
		return "", models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		h.l.Errorf(ctx, "users.delivery.http.processDeleteUserRequest.Validate: %v", err)
		return "", models.Scope{}, errWrongBody
	}

	sc := jwt.NewScope(payload)

	return id, sc, nil
}
