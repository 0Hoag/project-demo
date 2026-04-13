package roles

type CreateInput struct {
	Name string `json:"name"`
}

type UpdateInput struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
}

type ListInput struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

// AttachPermissionInput gắn permission vào role (insert vào bảng role_permissions).
type AttachPermissionInput struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}
