-include .env
export
BINARY=project-demo

run-api:
	@echo "Running the application"
	go run cmd/api/main.go

.PHONY: db-up db-down db-logs migrate-up migrate-down sqlc-gen tidy

db-up:
	docker compose up -d postgres

db-down:
	docker compose down

db-logs:
	docker compose logs -f postgres

migrate-up:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir internal/db/migrations postgres "$$DATABASE_URL" up

# migrate-down:
# 	go run github.com/pressly/goose/v3/cmd/goose@latest -dir internal/db/migrations postgres "$$DATABASE_URL" down

sqlc-gen:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate

tidy:
	go mod tidy
