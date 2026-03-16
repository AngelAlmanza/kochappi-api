# Project Structure

This is a **backend API server** written in Go. The structure follows Hexagonal Architecture (Ports & Adapters) to keep the business logic independent from HTTP, database, and other infrastructure concerns.

```
kochappi-api/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
│
├── internal/
│   ├── domain/                     # Domain Layer
│   │   ├── entity/                 # Core entities
│   │   │   ├── trainer.go
│   │   │   ├── client.go
│   │   │   ├── routine.go
│   │   │   ├── exercise.go
│   │   │   ├── session.go
│   │   │   └── user.go
│   │   │
│   │   ├── value_object/          # Value objects (immutable concepts)
│   │   │   ├── one_rm.go
│   │   │   ├── weight.go
│   │   │   ├── email.go
│   │   │   └── password.go
│   │   │
│   │   └── error/                 # Domain-specific errors
│   │       ├── routine_not_found.go
│   │       ├── invalid_client_assignment.go
│   │       └── errors.go
│   │
│   ├── application/               # Application Layer
│   │   ├── service/               # Use case services
│   │   │   ├── routine/
│   │   │   │   ├── create_routine.go
│   │   │   │   ├── update_routine.go
│   │   │   │   └── list_routines.go
│   │   │   │
│   │   │   ├── session/
│   │   │   │   ├── register_training_session.go
│   │   │   │   └── get_session_history.go
│   │   │   │
│   │   │   ├── one_rm/
│   │   │   │   ├── register_one_rm.go
│   │   │   │   └── get_one_rm_history.go
│   │   │   │
│   │   │   └── client/
│   │   │       ├── register_client.go
│   │   │       └── list_clients.go
│   │   │
│   │   ├── dto/                   # Data Transfer Objects
│   │   │   ├── routine_request.go
│   │   │   ├── routine_response.go
│   │   │   └── session_request.go
│   │   │
│   │   └── port/                  # Application Ports (Interfaces)
│   │       ├── input_port.go      # Inbound contracts
│   │       └── output_port.go     # Outbound contracts
│   │
│   ├── adapter/                   # Adapter Layer
│   │   ├── http/                  # HTTP/REST Adapters (Inbound)
│   │   │   ├── handler/
│   │   │   │   ├── routine_handler.go
│   │   │   │   ├── session_handler.go
│   │   │   │   ├── one_rm_handler.go
│   │   │   │   ├── auth_handler.go
│   │   │   │   └── client_handler.go
│   │   │   │
│   │   │   ├── middleware/
│   │   │   │   ├── auth_middleware.go
│   │   │   │   ├── cors_middleware.go
│   │   │   │   └── error_handler.go
│   │   │   │
│   │   │   └── router.go           # Routing setup
│   │   │
│   │   ├── persistence/           # Database Adapters (Outbound)
│   │   │   ├── postgres/
│   │   │   │   ├── routine_repository.go
│   │   │   │   ├── session_repository.go
│   │   │   │   ├── one_rm_repository.go
│   │   │   │   ├── client_repository.go
│   │   │   │   ├── user_repository.go
│   │   │   │   └── migrations/
│   │   │   │       ├── 001_create_users.sql
│   │   │   │       ├── 002_create_trainers.sql
│   │   │   │       ├── 003_create_clients.sql
│   │   │   │       ├── 004_create_routines.sql
│   │   │   │       └── ...
│   │   │   │
│   │   │   └── mock/              # Mock implementations for testing
│   │   │       ├── mock_routine_repository.go
│   │   │       └── mock_session_repository.go
│   │   │
│   │   ├── auth/                  # Authentication Adapter
│   │   │   ├── jwt_provider.go
│   │   │   ├── password_hasher.go
│   │   │   └── token_claims.go
│   │   │
│   │   └── config/                # Configuration
│   │       └── config.go
│   │
│   └── shared/                    # Shared utilities
│       ├── logger/
│       │   └── logger.go
│       │
│       ├── validator/
│       │   └── validator.go
│       │
│       └── pagination/
│           └── pagination.go
│
├── test/                          # Integration & E2E tests
│   ├── integration/
│   │   ├── routine_test.go
│   │   ├── session_test.go
│   │   └── testhelpers.go
│   │
│   ├── fixtures/
│   │   ├── trainer_fixture.go
│   │   ├── client_fixture.go
│   │   └── routine_fixture.go
│   │
│   └── docker-compose.test.yml    # Test database
│
├── pkg/                           # Public packages (if needed for library use)
│   └── errors/
│       └── errors.go
│
├── config/
│   ├── .env.example
│   └── .env.local
│
├── scripts/
│   ├── migrate.sh
│   └── seed.sh
│
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## Where Does Each Concept Live?

| Concept | Package | Example |
|---------|---------|---------|
| Business entities | `internal/domain/entity/` | `Trainer`, `Routine` |
| Immutable value types | `internal/domain/value_object/` | `Email`, `Weight` |
| Domain errors | `internal/domain/error/` | `ClientNotFoundError` |
| Use case logic | `internal/application/service/` | `CreateRoutineUseCase` |
| Input/output shapes | `internal/application/dto/` | `CreateRoutineRequest` |
| Repository contracts | `internal/application/port/` | `RoutineRepository` interface |
| HTTP handlers | `internal/adapter/http/handler/` | `RoutineHandler` |
| Middleware | `internal/adapter/http/middleware/` | `AuthMiddleware` |
| DB implementations | `internal/adapter/persistence/postgres/` | `PostgresRoutineRepository` |
| Auth (JWT) | `internal/adapter/auth/` | `JWTProvider` |
| Cross-cutting utils | `internal/shared/` | logger, validator, pagination |
