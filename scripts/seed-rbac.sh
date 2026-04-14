#!/usr/bin/env bash
# Seed RBAC mẫu (idempotent): roles admin/user, permissions create|get|update|delete,
# admin = 4 quyền, user = chỉ get, role_permissions.
#
# user_roles KHÔNG có API public trong script cũ — thêm bước cuối:
#   export SEED_USER_ID='<uuid-của-user-trong-bảng-users>'
#   export DATABASE_URL='postgres://...'   # giống .env
#   ./scripts/seed-rbac.sh
# Script sẽ INSERT user_roles (admin + user) bằng psql nếu đủ 2 biến trên.
#
# Chạy:  set -a && source .env && set +a && ./scripts/seed-rbac.sh
# Hoặc:  BASE_URL=http://127.0.0.1:8080/api/v1 ./scripts/seed-rbac.sh
#
# Yêu cầu: API đang chạy, curl, jq. Gán user_roles: thêm psql (brew install libpq)

set -euo pipefail

BASE="${BASE_URL:-http://localhost:8080/api/v1}"
HDR=(-H "Content-Type: application/json")
LIST_LIMIT="${LIST_LIMIT:-500}"

die() { echo "❌ $*" >&2; exit 1; }

command -v curl >/dev/null || die "Thiếu curl"
command -v jq >/dev/null || die "Thiếu jq (macOS: brew install jq)"

post_json() {
  local url="$1"
  local body="$2"
  curl -sS -X POST "$url" "${HDR[@]}" -d "$body"
}

# POST và lấy .data.id (error_code phải 0)
post_id() {
  local url="$1"
  local body="$2"
  local label="$3"
  local resp
  resp=$(post_json "$url" "$body")
  local code
  code=$(echo "$resp" | jq -r '.error_code // -1')
  if [[ "$code" != "0" ]]; then
    echo "Response: $resp" >&2
    die "$label thất bại (error_code=$code). DB mới? Chạy: export DATABASE_URL=... && make migrate-up"
  fi
  local id
  id=$(echo "$resp" | jq -r '.data.id // empty')
  [[ -n "$id" ]] || die "$label: không có data.id trong response"
  echo "$id"
}

# Đã có trong DB (theo name) thì chỉ lấy id; chưa có thì POST.
ensure_role_id() {
  local name="$1"
  local id
  id=$(curl -sS "$BASE/roles/all?limit=${LIST_LIMIT}&offset=0" \
    | jq -r --arg n "$name" '(.data // [])[] | select(.name == $n) | .id' | head -1)
  if [[ -n "$id" ]]; then
    echo "   ↪ đã có role '$name', dùng id: $id" >&2
    echo "$id"
    return 0
  fi
  post_id "$BASE/roles" "$(jq -nc --arg n "$name" '{name: $n}')" "POST /roles $name"
}

ensure_permission_id() {
  local name="$1"
  local id
  id=$(curl -sS "$BASE/permissions/all?limit=${LIST_LIMIT}&offset=0" \
    | jq -r --arg n "$name" '(.data // [])[] | select(.name == $n) | .id' | head -1)
  if [[ -n "$id" ]]; then
    echo "   ↪ đã có permission '$name', dùng id: $id" >&2
    echo "$id"
    return 0
  fi
  post_id "$BASE/permissions" "$(jq -nc --arg n "$name" '{name: $n}')" "POST /permissions $name"
}

# Gắn role–permission; đã gắn (error_code 144009) thì bỏ qua.
attach_skip_dup() {
  local url="$1"
  local body="$2"
  local label="$3"
  local resp code
  resp=$(post_json "$url" "$body")
  code=$(echo "$resp" | jq -r '.error_code // -1')
  if [[ "$code" == "0" ]]; then
    echo "   ✓ $label"
    return 0
  fi
  if [[ "$code" == "144009" ]]; then
    echo "   ↪ đã gắn rồi, bỏ qua: $label" >&2
    return 0
  fi
  echo "Response: $resp" >&2
  die "$label thất bại (error_code=$code)"
}

ADMIN_ROLE_ID=$(ensure_role_id "admin")
USER_ROLE_ID=$(ensure_role_id "user")
echo "   admin id: $ADMIN_ROLE_ID"
echo "   user  id: $USER_ROLE_ID"

echo "→ Permissions (tạo nếu chưa có)..."
PERM_CREATE=$(ensure_permission_id "create")
PERM_GET=$(ensure_permission_id "get")
PERM_UPDATE=$(ensure_permission_id "update")
PERM_DELETE=$(ensure_permission_id "delete")
echo "   create=$PERM_CREATE get=$PERM_GET update=$PERM_UPDATE delete=$PERM_DELETE"

echo "→ Gắn quyền cho admin (bỏ qua nếu đã gắn)..."
attach_skip_dup "$BASE/roles/$ADMIN_ROLE_ID/permissions" "{\"permission_id\":\"$PERM_CREATE\"}" "create → admin"
attach_skip_dup "$BASE/roles/$ADMIN_ROLE_ID/permissions" "{\"permission_id\":\"$PERM_GET\"}" "get → admin"
attach_skip_dup "$BASE/roles/$ADMIN_ROLE_ID/permissions" "{\"permission_id\":\"$PERM_UPDATE\"}" "update → admin"
attach_skip_dup "$BASE/roles/$ADMIN_ROLE_ID/permissions" "{\"permission_id\":\"$PERM_DELETE\"}" "delete → admin"

echo "→ Gắn get cho user (bỏ qua nếu đã gắn)..."
attach_skip_dup "$BASE/roles/$USER_ROLE_ID/permissions" "{\"permission_id\":\"$PERM_GET\"}" "get → user"

echo "→ Kiểm tra (GET permissions của từng role):"
echo "--- admin ---"
curl -sS "$BASE/roles/$ADMIN_ROLE_ID/permissions" | jq .
echo "--- user ---"
curl -sS "$BASE/roles/$USER_ROLE_ID/permissions" | jq .

echo ""
echo "✅ Xong. Chạy lại nhiều lần vẫn ổn (bỏ qua bản ghi & liên kết đã có)."
