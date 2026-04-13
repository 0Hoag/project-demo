package http

import (
	"time"

	"github.com/google/uuid"
	domain "github.com/zeross/project-demo/internal/models"
	"github.com/zeross/project-demo/internal/permissions"
)

type createPermissionReq struct {
	Name string `json:"name" binding:"required"`
}

func (r createPermissionReq) toInput() permissions.CreateInput {
	return permissions.CreateInput{Name: r.Name}
}

func (r createPermissionReq) validate() error {
	if r.Name == "" {
		return errWrongBody
	}
	return nil
}

type updatePermissionReq struct {
	ID   string  `json:"id" binding:"required"`
	Name *string `json:"name"`
}

func (r updatePermissionReq) toInput() permissions.UpdateInput {
	return permissions.UpdateInput{ID: r.ID, Name: r.Name}
}

func (r updatePermissionReq) validate() error {
	if uuid.Validate(r.ID) != nil {
		return errWrongQuery
	}
	return nil
}

type listPermissionsReq struct {
	Limit  int32 `form:"limit"`
	Offset int32 `form:"offset"`
}

func (r listPermissionsReq) toInput() permissions.ListInput {
	return permissions.ListInput{Limit: r.Limit, Offset: r.Offset}
}

func (r listPermissionsReq) validate() error {
	if r.Limit <= 0 {
		return errWrongBody
	}
	return nil
}

type permissionResp struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h handler) newPermissionResp(p domain.Permission) permissionResp {
	return permissionResp{
		ID:        p.ID.String(),
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (h handler) newListPermissionResp(ps []domain.Permission) []permissionResp {
	out := make([]permissionResp, 0, len(ps))
	for _, p := range ps {
		out = append(out, h.newPermissionResp(p))
	}
	return out
}
