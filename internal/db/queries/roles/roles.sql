-- name: CreateRole :one
INSERT INTO roles (name) VALUES ($1)
RETURNING id, name, created_at, updated_at, deleted_at;

-- name: GetRoleByID :one
SELECT id, name, created_at, updated_at, deleted_at
FROM roles
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetRoleByName :one
SELECT id, name, created_at, updated_at, deleted_at
FROM roles
WHERE name = $1 AND deleted_at IS NULL;

-- name: ListRoles :many
SELECT id, name, created_at, updated_at, deleted_at
FROM roles
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateRole :one
UPDATE roles SET
  name = $2,
  updated_at = now()
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, name, created_at, updated_at, deleted_at;

-- name: DeleteRole :exec
UPDATE roles SET
  deleted_at = now(),
  updated_at = now()
WHERE id = $1 AND deleted_at IS NULL;
