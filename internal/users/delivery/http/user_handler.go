package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zeross/project-demo/internal/users"
	"github.com/zeross/project-demo/pkg/response"
)

func (h handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processCreateRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.CreateUser.processCreateRequest: %v", err)
		response.Error(c, err)
		return
	}

	user, err := h.uc.CreateUser(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.CreateUser.CreateUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newUserResp(user))
}

func (h handler) DetailUser(c *gin.Context) {
	ctx := c.Request.Context()

	id, sc, err := h.processDetailUserRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.DetailUser.processDetailUserRequest: %v", err)
		response.Error(c, err)
		return
	}

	user, err := h.uc.DetailUser(ctx, sc, id)
	if err != nil {
		h.l.Errorf(ctx, "http.DetailUser.DetailUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}
	response.OK(c, h.newUserResp(user))
}

func (h handler) ListUsers(c *gin.Context) {
	ctx := c.Request.Context()

	req, sc, err := h.processListUsersRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.ListUsers.processListUsersRequest: %v", err)
		response.Error(c, err)
		return
	}

	users, err := h.uc.ListUsers(ctx, sc, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.ListUsers.ListUsers: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newListUserResp(users))
}

func (h handler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	req, pq, sc, err := h.processGetUsersRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.GetUsers.processGetUsersRequest: %v", err)
		response.Error(c, err)
		return
	}

	var input users.GetInput
	input.Filter = req.toFilter()
	pq.Adjust()
	input.PagQuery = pq

	users, err := h.uc.GetUsers(ctx, sc, input)
	if err != nil {
		h.l.Errorf(ctx, "http.GetUsers.GetUsers: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newListUserResp(users))
}

func (h handler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()

	req, sc, err := h.processUpdateRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.UpdateUser.processUpdateRequest: %v", err)
		response.Error(c, err)
		return
	}

	err = h.uc.UpdateUser(ctx, sc, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.UpdateUser.UpdateUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, nil)
}

func (h handler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	id, sc, err := h.processDeleteUserRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "http.DeleteUser.processDeleteUserRequest: %v", err)
		response.Error(c, err)
		return
	}

	err = h.uc.DeleteUser(ctx, sc, id)
	if err != nil {
		h.l.Errorf(ctx, "http.DeleteUser.DeleteUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}
	response.OK(c, nil)
}
