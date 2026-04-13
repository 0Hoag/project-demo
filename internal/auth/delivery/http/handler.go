package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zeross/project-demo/pkg/response"
)

func (h handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processLoginRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "auth.delivery.http.Login.processLoginRequest: %v", err)
		response.Error(c, err)
		return
	}

	out, err := h.uc.Login(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "auth.delivery.http.Login.Login: %v", err)
		response.Error(c, h.mapError(err))
		return
	}

	response.OK(c, h.newAuthResp(out))
}
