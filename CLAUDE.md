# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Development Commands

```bash
make dev              # Hot reload with air
make build            # Production build (Linux, CGO disabled)
make test             # All tests with coverage
make test-unit        # Unit tests only (./internal/...)
make test-int         # Integration tests (./test/integration/...)
make lint             # golangci-lint
make swagger          # Generate Swagger docs (swag)
make docker-up        # Start PostgreSQL via docker-compose (port 5433)
make docker-down      # Stop PostgreSQL
```

Run a single test:
```bash
go test -run TestFunctionName ./internal/application/service/auth/...
```

## Architecture

Hexagonal architecture (Ports & Adapters) with strict dependency direction: adapters → application → domain.

**Module**: `kochappi` | **Framework**: Gin | **ORM**: GORM | **DB**: PostgreSQL 15+

### Layer Breakdown

- **`internal/domain/`** — Pure business logic with zero dependencies. Contains entities, value objects (Email, Password with embedded validation), and a `DomainError` interface with 30+ typed error implementations.

- **`internal/application/`** — Orchestration layer.
  - `port/output_port.go` — All repository and service interfaces (input ports are implicit in use case signatures).
  - `service/<feature>/` — Each use case is a struct with a single `Execute()` method. Grouped by domain: auth, customers, exercises, templates, routines, progress, sessions.
  - `dto/` — Request/response objects for HTTP serialization.

- **`internal/adapter/`** — Infrastructure implementations.
  - `http/` — Gin router, handlers (one per feature), middleware (JWT auth, centralized error mapping).
  - `persistence/postgres/` — GORM repositories. Each has a `model/` subpackage with `ToDomain()`/`ModelFromDomain()` converters.
  - `persistence/mock/` — Mock repositories for unit tests.
  - `auth/` — JWT provider (access + refresh tokens, HS256), bcrypt hasher, console OTP service.
  - `storage/` — Local file storage (`./uploads`).
  - `cron/` — Scheduled tasks (daily workout session generation at 10:00 UTC).
  - `config/` — Environment config loaded from `.env.local` via godotenv.

### Wiring

`cmd/api/main.go` is the composition root: initializes all adapters, repositories, use cases, and handlers via constructor injection — no service locator or DI framework.

### Error Handling

Domain errors implement `DomainError` interface (`Code()` string, `IsUserError()` bool). The `error_handler.go` middleware maps error codes to HTTP statuses:
- `NOT_FOUND` → 404
- `CONFLICT` → 409
- `UNPROCESSABLE_ENTITY` → 422
- `VALIDATION_ERROR` → 400
- `UNAUTHORIZED` → 401

### Schema Management

GORM AutoMigrate — no external migration tool. Schema changes happen through GORM model struct tags in `persistence/postgres/model/`.

## Conventions

- **Use cases**: `*UseCase` suffix, single `Execute()` method, constructor-injected dependencies
- **Repositories**: Interface in `port/output_port.go`, implementation in `persistence/postgres/`, mock in `persistence/mock/`
- **Handlers**: One per feature domain, registered in `router.go`
- **Database models**: Separate from domain entities; always convert via `ToDomain()`/`ModelFromDomain()`
- **Tests**: Follow `Test[Function]_[Scenario]` naming; use mock repositories from `persistence/mock/`
- **All DB operations**: Accept `context.Context` as first parameter


## Version Control
- Use Git for version control
- Write clear and descriptive commit messages that explain the changes made, they must be in english and follow the format: `type(scope): description`, where type can be `feat`, `fix`, `docs`, `style`, `refactor`, `test`, or `chore`.
