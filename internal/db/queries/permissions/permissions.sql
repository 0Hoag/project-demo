-- name: CreatePermission :one
INSERT INTO permissions (name) VALUES ($1)
RETURNING id, name, created_at, updated_at, deleted_at;

-- name: GetPermissionByID :one
SELECT id, name, created_at, updated_at, deleted_at
FROM permissions
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListPermissions :many
SELECT id, name, created_at, updated_at, deleted_at
FROM permissions
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdatePermission :one
UPDATE permissions SET
  name = $2,
  updated_at = now()
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, name, created_at, updated_at, deleted_at;

-- name: DeletePermission :exec
UPDATE permissions SET
  deleted_at = now(),
  updated_at = now()
WHERE id = $1 AND deleted_at IS NULL;
