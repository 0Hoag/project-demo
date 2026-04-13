-- name: AttachPermissionToRole :exec
INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2);

-- name: DetachPermissionFromRole :exec
DELETE FROM role_permissions
WHERE role_id = $1 AND permission_id = $2;

-- name: ListPermissionsForRole :many
SELECT
  p.id, p.name, p.created_at, p.updated_at, p.deleted_at
FROM permissions p
INNER JOIN role_permissions rp ON rp.permission_id = p.id
WHERE rp.role_id = $1 AND p.deleted_at IS NULL
ORDER BY p.name ASC;
