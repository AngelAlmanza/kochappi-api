# Rules & Conventions

These are the team standards. Follow them in every PR.

---

## 1. Naming Conventions

| Element | Convention | Example |
|---------|-----------|---------|
| Files | snake_case | `create_routine.go` |
| Packages | lowercase | `routine`, `session` |
| Types / Structs | PascalCase | `Routine`, `CreateRoutineUseCase` |
| Functions / Methods | PascalCase | `Create()`, `GetByID()` |
| Constants | UPPER_SNAKE_CASE | `MAX_CLIENTS`, `ROUTINE_STATUS_ACTIVE` |
| Interfaces | `NameRepository` / `NameProvider` | `RoutineRepository`, `JWTProvider` |

---

## 2. Dependency Injection

All dependencies must be injected through constructors. **No global variables. No singletons.**

```go
// ✅ CORRECT
func NewCreateRoutineUseCase(
    routineRepo port.RoutineRepository,
    clientRepo port.ClientRepository,
) *CreateRoutineUseCase {
    return &CreateRoutineUseCase{
        routineRepo: routineRepo,
        clientRepo:  clientRepo,
    }
}

// ❌ WRONG — global variable
var globalRoutineRepo port.RoutineRepository

func CreateRoutine() {
    globalRoutineRepo.Create(...)
}
```

**Why?** Globals make testing impossible and hide dependencies.

---

## 3. Context Usage

Always pass `context.Context` as the **first parameter** for any function that does I/O. This enables cancellation and timeout propagation.

```go
// ✅ CORRECT
func (r *RoutineRepository) Create(ctx context.Context, routine *entity.Routine) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }
    return r.db.WithContext(ctx).Create(routine).Error
}

// ❌ WRONG — no context, no cancellation support
func (r *RoutineRepository) Create(routine *entity.Routine) error {
    return r.db.Create(routine).Error
}
```

---

## 4. Error Handling

Define domain-specific error types. Never use raw `errors.New("something went wrong")` for domain errors.

```go
// internal/domain/error/errors.go
package error

import "fmt"

type DomainError interface {
    error
    Code() string
    IsUserError() bool // true = safe to show to the end user
}

type ClientNotFoundError struct {
    ClientID string
}

func (e *ClientNotFoundError) Error() string {
    return fmt.Sprintf("client %s not found", e.ClientID)
}

func (e *ClientNotFoundError) Code() string {
    return "CLIENT_NOT_FOUND"
}

func (e *ClientNotFoundError) IsUserError() bool {
    return true
}
```

The HTTP handler maps `DomainError` codes to HTTP status codes in `handleError()`.

---

## 5. Repository Contracts

Every repository must implement its output port interface. This allows tests to swap the real DB for a mock.

```go
// PostgreSQL implementation
type PostgresRoutineRepository struct {
    db *gorm.DB
}

func (r *PostgresRoutineRepository) Create(ctx context.Context, routine *entity.Routine) error {
    return r.db.WithContext(ctx).Create(routine).Error
}

// Mock for unit tests — same interface, no database
type MockRoutineRepository struct {
    CreateFn func(ctx context.Context, routine *entity.Routine) error
}

func (r *MockRoutineRepository) Create(ctx context.Context, routine *entity.Routine) error {
    return r.CreateFn(ctx, routine)
}
```

---

## 6. No Business Logic in HTTP Handlers

Handlers are **orchestrators only**. All logic lives in use cases.

```go
// ❌ WRONG — business logic in handler
func (h *RoutineHandler) CreateRoutine(w http.ResponseWriter, r *http.Request) {
    if req.Weight > 300 {
        http.Error(w, "Weight too high", http.StatusBadRequest)
        return
    }
}

// ✅ CORRECT — logic in use case, handler just delegates
func (uc *RegisterTrainingSessionUseCase) Execute(ctx context.Context, req *dto.RegisterSessionRequest) error {
    if req.Weight > 300 {
        return ErrWeightTooHigh
    }
    return nil
}

func (h *SessionHandler) RegisterSession(w http.ResponseWriter, r *http.Request) {
    resp, err := h.registerSessionUseCase.Execute(r.Context(), &req)
    if err != nil {
        handleError(w, err)
        return
    }
    json.NewEncoder(w).Encode(resp)
}
```

---

## 7. Authentication vs Authorization

| Concern | Where | Responsibility |
|---------|-------|----------------|
| **Authentication** (who are you?) | Middleware | Verify JWT, set `trainer_id` in context |
| **Authorization** (can you do this?) | Use case | Check ownership, permissions |

```go
// Middleware: only checks token validity
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := extractToken(r)
        claims, err := jwtProvider.Verify(token)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        ctx := context.WithValue(r.Context(), "trainer_id", claims.TrainerID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Use case: checks that the authenticated trainer owns the resource
func (uc *CreateRoutineUseCase) Execute(ctx context.Context, req *dto.CreateRoutineRequest) error {
    client, err := uc.clientRepo.GetByID(ctx, req.ClientID)
    if err != nil {
        return err
    }
    if client.TrainerID != req.TrainerID {
        return ErrUnauthorized
    }
    return nil
}
```

---

## PR Checklist

Before opening a pull request, verify:

- [ ] No business logic in HTTP handlers
- [ ] All dependencies injected via constructors (no globals)
- [ ] Custom error types used for domain errors
- [ ] Unit tests written for use cases (with mocks)
- [ ] Integration tests written for repository changes
- [ ] Database models kept separate from domain entities
- [ ] No framework imports in domain layer
- [ ] `context.Context` passed to all I/O operations
- [ ] Authorization checks inside use cases, not handlers
- [ ] Naming follows conventions table above
- [ ] `make lint` passes
