# Development Workflow

---

## 1. Local Setup

```bash
# Clone the repo
git clone <repo-url>
cd kochappi-api

# Copy environment variables and fill in values
cp config/.env.example config/.env.local

# Install Go dependencies
go mod download

# Start PostgreSQL via Docker
docker-compose up -d

# Run database migrations
make migrate

# (Optional) Seed test data
make seed

# Start the development server with hot reload
make dev
```

---

## 2. Makefile Commands

| Command | What it does |
|---------|-------------|
| `make dev` | Run server with hot reload (`air`) |
| `make test` | Run all tests with coverage |
| `make test-unit` | Run unit tests only (`./internal/...`) |
| `make test-int` | Run integration tests only |
| `make migrate` | Apply pending database migrations |
| `make seed` | Seed development data |
| `make lint` | Run `golangci-lint` |
| `make build` | Compile binary to `bin/kochappi-api` |
| `make run` | Run the compiled binary |
| `make docker-up` | Start Docker services |
| `make docker-down` | Stop Docker services |

Full Makefile:

```makefile
.PHONY: help dev test test-unit test-int migrate seed lint build run docker-up docker-down

dev:
	air

test:
	go test ./... -v -cover

test-unit:
	go test ./internal/... -v

test-int:
	go test ./test/integration/... -v

migrate:
	migrate -path internal/adapter/persistence/postgres/migrations -database "$$DATABASE_URL" up

seed:
	go run scripts/seed.go

lint:
	golangci-lint run

build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/kochappi-api ./cmd/api

run:
	./bin/kochappi-api

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down
```

---

## 3. Adding a New Feature — Step by Step

Follow this order to avoid circular dependencies:

1. **Domain entity** → `internal/domain/entity/new_feature.go`
2. **Output port** → add interface to `internal/application/port/output_port.go`
3. **Repository adapter** → `internal/adapter/persistence/postgres/new_feature_repository.go`
4. **DTOs** → add request/response structs to `internal/application/dto/`
5. **Use case** → `internal/application/service/new_feature/create_feature.go`
6. **HTTP handler** → `internal/adapter/http/handler/new_feature_handler.go`
7. **Route** → register in `internal/adapter/http/router.go`
8. **Tests** → `internal/application/service/new_feature/create_feature_test.go`

---

## 4. Adding a New Database Table — Step by Step

1. Write migration → `internal/adapter/persistence/postgres/migrations/XXX_create_new_table.sql`
2. Create DB model → `internal/adapter/persistence/postgres/model/new_table_model.go`
3. Create domain entity → `internal/domain/entity/new_entity.go`
4. Create repository → `internal/adapter/persistence/postgres/new_entity_repository.go`
5. Add port contract → `internal/application/port/output_port.go`

---

## 5. Environment Variables

```bash
# config/.env.local (never commit this file)
DATABASE_URL=postgresql://kochappi:password@localhost:5432/kochappi_dev
JWT_SECRET=your-local-secret
PORT=8080
ENV=development
LOG_LEVEL=debug
```
