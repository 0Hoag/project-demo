-- +goose Up
-- RBAC baseline + seed admin. New signups get role "user" in app code (CreateUser).
-- Login admin: phone __seed_admin__, password admin123 (bcrypt below).

INSERT INTO roles (name) VALUES ('admin') ON CONFLICT (name) DO NOTHING;
INSERT INTO roles (name) VALUES ('user') ON CONFLICT (name) DO NOTHING;

INSERT INTO permissions (name) VALUES ('create') ON CONFLICT (name) DO NOTHING;
INSERT INTO permissions (name) VALUES ('get') ON CONFLICT (name) DO NOTHING;
INSERT INTO permissions (name) VALUES ('update') ON CONFLICT (name) DO NOTHING;
INSERT INTO permissions (name) VALUES ('delete') ON CONFLICT (name) DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
CROSS JOIN permissions p
WHERE r.name = 'admin' AND r.deleted_at IS NULL AND p.deleted_at IS NULL
  AND p.name IN ('create', 'get', 'update', 'delete')
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
CROSS JOIN permissions p
WHERE r.name = 'user' AND r.deleted_at IS NULL AND p.deleted_at IS NULL AND p.name = 'get'
ON CONFLICT DO NOTHING;

INSERT INTO users (username, phone, password, avatar_url, bio)
SELECT
  'admin',
  '__seed_admin__',
  '$2a$10$DnUpv19YSGVup.PEZTKWeupywyHydwInr9Ix1kDg0vtMWNuvyaytm',
  '',
  ''
WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = 'admin' AND deleted_at IS NULL);

INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id
FROM users u
CROSS JOIN roles r
WHERE u.username = 'admin'
  AND u.deleted_at IS NULL
  AND r.name = 'admin'
  AND r.deleted_at IS NULL
ON CONFLICT DO NOTHING;

-- +goose Down
DELETE FROM user_roles
WHERE user_id IN (SELECT id FROM users WHERE username = 'admin');

DELETE FROM users WHERE username = 'admin';
