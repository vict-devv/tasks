.PHONY: build
build:
	@echo "Building project..."
	go build -o bin/tasks cmd/tasks/main.go

.PHONY: run
run: build
	@echo "Building and Running project..."
	./bin/tasks

.PHONY: clean
clean:
	@echo "Cleaning project..."
	rm -rf bin/

.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

.PHONY: lint
lint:
	@echo "Running linter..."
	golangci-lint-v2 run

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	golangci-lint-v2 run --fix

.PHONY: migrate-local-up
migrate-local-up:
	@echo "Running local database migrations..."
	migrate -path "./internal/db/migrations" -database "postgres://admin:123456@localhost:5432/tasks_db?sslmode=disable" up

.PHONY: migrate-local-down
migrate-local-down:
	@echo "Reverting local database migrations..."
	migrate -path "./internal/db/migrations" -database "postgres://admin:123456@localhost:5432/tasks_db?sslmode=disable" down

.PHONY: migrate-local-up-steps
migrate-local-up-steps:
	@echo "Running local database migrations by steps: ${steps}..."
	migrate -path "./internal/db/migrations" -database "postgres://admin:123456@localhost:5432/tasks_db?sslmode=disable" up ${steps}

.PHONY: migrate-local-down-steps
migrate-local-down-steps:
	@echo "Reverting local database migrations by steps: ${steps}..."
	migrate -path "./internal/db/migrations" -database "postgres://admin:123456@localhost:5432/tasks_db?sslmode=disable" down ${steps}

.PHONY: migrate-local-force
migrate-local-force:
	@echo "Forcing migration to specific version: ${version}..."
	migrate -path "./internal/db/migrations" -database "postgres://admin:123456@localhost:5432/tasks_db?sslmode=disable" force ${version}
