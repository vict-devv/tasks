build:
	@echo "Building project..."
	go build -o bin/tasks cmd/tasks/main.go

run: build
	@echo "Building and Running project..."
	./bin/tasks

clean:
	@echo "Cleaning project..."
	rm -rf bin/

test:
	@echo "Running tests..."
	go test ./...

migrate-local-up:
	@echo "Running local database migrations..."
	migrate -path "./internal/db/migrations" -database "postgres://admin:123456@localhost:5432/tasks_db?sslmode=disable" up

migrate-local-down:
	@echo "Reverting local database migrations..."
	migrate -path "./internal/db/migrations" -database "postgres://admin:123456@localhost:5432/tasks_db?sslmode=disable" down

migrate-local-force:
	@echo "Forcing migration to specific version: ${version}..."
	migrate -path "./internal/db/migrations" -database "postgres://admin:123456@localhost:5432/tasks_db?sslmode=disable" force ${version}

.PHONY: build run clean test migrate-local-up migrate-local-down migrate-local-force