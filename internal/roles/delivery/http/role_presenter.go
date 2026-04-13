package http

import (
	"time"

	"github.com/google/uuid"
	domain "github.com/zeross/project-demo/internal/models"
	"github.com/zeross/project-demo/internal/roles"
)

type createRoleReq struct {
	Name string `json:"name" binding:"required"`
}

func (r createRoleReq) toInput() roles.CreateInput {
	return roles.CreateInput{Name: r.Name}
}

func (r createRoleReq) validate() error {
	if r.Name == "" {
		return errWrongBody
	}
	return nil
}

type updateRoleReq struct {
	ID   string  `json:"id" binding:"required"`
	Name *string `json:"name"`
}

func (r updateRoleReq) toInput() roles.UpdateInput {
	return roles.UpdateInput{ID: r.ID, Name: r.Name}
}

func (r updateRoleReq) validate() error {
	if uuid.Validate(r.ID) != nil {
		return errWrongQuery
	}
	return nil
}

type listRolesReq struct {
	Limit  int32 `form:"limit"`
	Offset int32 `form:"offset"`
}

func (r listRolesReq) toInput() roles.ListInput {
	return roles.ListInput{Limit: r.Limit, Offset: r.Offset}
}

func (r listRolesReq) validate() error {
	if r.Limit <= 0 {
		return errWrongBody
	}
	return nil
}

type roleResp struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h handler) newRoleResp(r domain.Role) roleResp {
	return roleResp{
		ID:        r.ID.String(),
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func (h handler) newListRoleResp(rs []domain.Role) []roleResp {
	out := make([]roleResp, 0, len(rs))
	for _, r := range rs {
		out = append(out, h.newRoleResp(r))
	}
	return out
}

// attachPermToRoleReq body: chỉ cần permission_id — role_id lấy từ URL /roles/:id/permissions
type attachPermToRoleReq struct {
	PermissionID string `json:"permission_id" binding:"required"`
}

func (r attachPermToRoleReq) toAttachInput(roleID string) roles.AttachPermissionInput {
	return roles.AttachPermissionInput{
		RoleID:       roleID,
		PermissionID: r.PermissionID,
	}
}

func (r attachPermToRoleReq) validate() error {
	if uuid.Validate(r.PermissionID) != nil {
		return errWrongBody
	}
	return nil
}

type permissionEmbedResp struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h handler) newPermissionEmbedResp(p domain.Permission) permissionEmbedResp {
	return permissionEmbedResp{
		ID:        p.ID.String(),
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (h handler) newListPermissionEmbedResp(ps []domain.Permission) []permissionEmbedResp {
	out := make([]permissionEmbedResp, 0, len(ps))
	for _, p := range ps {
		out = append(out, h.newPermissionEmbedResp(p))
	}
	return out
}
