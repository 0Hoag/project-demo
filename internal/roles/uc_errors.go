package roles

import "errors"

var (
	ErrRoleNotFound = errors.New("role not found")

	// ErrPermissionAlreadyLinked khi INSERT role_permissions bị trùng (role_id, permission_id).
	ErrPermissionAlreadyLinked = errors.New("permission already linked to this role")
)
