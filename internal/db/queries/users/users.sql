-- name: CreateUser :one
INSERT INTO users (
  username, phone, password, avatar_url, bio, birthday
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING
  id, username, phone, password, avatar_url, bio, birthday,
  created_at, updated_at, deleted_at;

-- name: DetailUser :one
SELECT
  id, username, phone, password, avatar_url, bio, birthday,
  created_at, updated_at, deleted_at
FROM users
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetUserByPhone :one
SELECT
  id, username, phone, password, avatar_url, bio, birthday,
  created_at, updated_at, deleted_at
FROM users
WHERE phone = $1 AND deleted_at IS NULL;

-- name: InsertUserRole :exec
INSERT INTO user_roles (user_id, role_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: ListRoleNamesForUser :many
SELECT r.name
FROM roles r
INNER JOIN user_roles ur ON ur.role_id = r.id
WHERE ur.user_id = $1 AND r.deleted_at IS NULL
ORDER BY r.name;

-- name: ListPermissionNamesForUser :many
SELECT DISTINCT p.name
FROM permissions p
INNER JOIN role_permissions rp ON rp.permission_id = p.id
INNER JOIN user_roles ur ON ur.role_id = rp.role_id
WHERE ur.user_id = $1 AND p.deleted_at IS NULL
ORDER BY p.name;

-- name: ListUsers :many
SELECT
  id, username, phone, password, avatar_url, bio, birthday,
  created_at, updated_at, deleted_at
FROM users
WHERE deleted_at IS NULL
  AND (sqlc.narg('id')::uuid IS NULL OR id = sqlc.narg('id'))
  AND (sqlc.narg('username')::text IS NULL OR username ILIKE '%' || sqlc.narg('username') || '%')
  AND (sqlc.narg('phone')::text IS NULL OR phone = sqlc.narg('phone'))
ORDER BY created_at DESC;

-- name: GetUsers :many
SELECT
  id, username, phone, password, avatar_url, bio, birthday,
  created_at, updated_at, deleted_at
FROM users
WHERE deleted_at IS NULL
  AND (sqlc.narg('id')::uuid IS NULL OR id = sqlc.narg('id'))
  AND (sqlc.narg('username')::text IS NULL OR username ILIKE '%' || sqlc.narg('username') || '%')
  AND (sqlc.narg('phone')::text IS NULL OR phone = sqlc.narg('phone'))
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountUsers :one
SELECT count(*) FROM users
WHERE deleted_at IS NULL
  AND (sqlc.narg('id')::uuid IS NULL OR id = sqlc.narg('id'))
  AND (sqlc.narg('username')::text IS NULL OR username ILIKE '%' || sqlc.narg('username') || '%')
  AND (sqlc.narg('phone')::text IS NULL OR phone = sqlc.narg('phone'));

-- name: UpdateUser :one
UPDATE users SET
  username = $2, phone = $3, password = $4, avatar_url = $5, bio = $6, birthday = $7
WHERE id = $1 AND deleted_at IS NULL RETURNING id, username, phone, password, avatar_url, bio, birthday, created_at, updated_at, deleted_at;


-- name: DeleteUser :exec
UPDATE users SET
  deleted_at = now(),
  updated_at = now()
WHERE id = $1 AND deleted_at IS NULL;
