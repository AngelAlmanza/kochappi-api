# Core Concepts

This document explains the five building blocks of the architecture with code examples.

---

## 1. Entities (Domain Layer)

Entities are core domain objects with **identity** (an ID) and **behavior** (methods that enforce business rules).

```go
// internal/domain/entity/trainer.go
package entity

type Trainer struct {
    ID        string
    UserID    string
    Name      string
    Email     string
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Domain rule: Trainer can have max 50 clients in v1
func (t *Trainer) CanAddClient(currentClientCount int) bool {
    return currentClientCount < 50
}
```

**Key rule:** Entities must have **zero imports** from external frameworks (no GORM tags, no Gin, etc.).

---

## 2. Value Objects (Domain Layer)

Value objects are **immutable** concepts that have no identity — two `Email` objects with the same value are equal.

```go
// internal/domain/value_object/email.go
package value_object

import "regexp"

type Email struct {
    value string // unexported — can only be set via constructor
}

func NewEmail(email string) (Email, error) {
    if !isValidEmail(email) {
        return Email{}, ErrInvalidEmail
    }
    return Email{value: email}, nil
}

func (e Email) String() string {
    return e.value
}
```

**Why value objects?** They centralize validation logic. You can never have an `Email` in memory that isn't valid — the constructor enforces it at creation time.

---

## 3. Repositories (Ports & Adapters)

Repositories define **how entities are persisted**, but the domain only sees the interface (port), not the implementation.

```go
// internal/application/port/output_port.go
package port

import (
    "context"
    "kochappi/internal/domain/entity"
)

// RoutineRepository defines the contract for routine persistence
type RoutineRepository interface {
    Create(ctx context.Context, routine *entity.Routine) error
    GetByID(ctx context.Context, id string) (*entity.Routine, error)
    Update(ctx context.Context, routine *entity.Routine) error
    List(ctx context.Context, clientID string) ([]*entity.Routine, error)
    Delete(ctx context.Context, id string) error
}

// ClientRepository defines the contract for client persistence
type ClientRepository interface {
    Create(ctx context.Context, client *entity.Client) error
    GetByID(ctx context.Context, id string) (*entity.Client, error)
    List(ctx context.Context, trainerID string) ([]*entity.Client, error)
}
```

The PostgreSQL implementation lives in `internal/adapter/persistence/postgres/` and implements this interface. Tests swap it for a mock — see [05_testing_strategy.md](./05_testing_strategy.md).

---

## 4. Use Cases / Services (Application Layer)

Use cases orchestrate domain logic to fulfill one specific business flow. They depend on **ports (interfaces)**, not on concrete implementations.

```go
// internal/application/service/routine/create_routine.go
package routine

import (
    "context"
    "kochappi/internal/application/dto"
    "kochappi/internal/application/port"
    "kochappi/internal/domain/entity"
)

type CreateRoutineUseCase struct {
    routineRepo port.RoutineRepository
    clientRepo  port.ClientRepository
}

func NewCreateRoutineUseCase(
    routineRepo port.RoutineRepository,
    clientRepo port.ClientRepository,
) *CreateRoutineUseCase {
    return &CreateRoutineUseCase{
        routineRepo: routineRepo,
        clientRepo:  clientRepo,
    }
}

func (uc *CreateRoutineUseCase) Execute(
    ctx context.Context,
    req *dto.CreateRoutineRequest,
) (*dto.RoutineResponse, error) {
    // 1. Validate client exists and belongs to trainer
    client, err := uc.clientRepo.GetByID(ctx, req.ClientID)
    if err != nil {
        return nil, ErrClientNotFound
    }

    if client.TrainerID != req.TrainerID {
        return nil, ErrUnauthorized
    }

    // 2. Create routine entity
    routine := entity.NewRoutine(
        req.ClientID,
        req.Name,
        req.Exercises,
    )

    // 3. Persist
    if err := uc.routineRepo.Create(ctx, routine); err != nil {
        return nil, err
    }

    return dto.FromRoutineEntity(routine), nil
}
```

**Notice:** the use case doesn't know if the repo is PostgreSQL, SQLite, or a mock. That's the point.

---

## 5. HTTP Handlers (Adapter Layer)

Handlers are thin. Their only job is to parse the HTTP request, call a use case, and format the response. **No business logic here.**

```go
// internal/adapter/http/handler/routine_handler.go
package handler

import (
    "encoding/json"
    "net/http"
    "kochappi/internal/adapter/http/middleware"
    "kochappi/internal/application/dto"
    "kochappi/internal/application/service/routine"
)

type RoutineHandler struct {
    createRoutineUseCase *routine.CreateRoutineUseCase
}

func NewRoutineHandler(
    createRoutineUseCase *routine.CreateRoutineUseCase,
) *RoutineHandler {
    return &RoutineHandler{
        createRoutineUseCase: createRoutineUseCase,
    }
}

// POST /api/v1/routines
func (h *RoutineHandler) CreateRoutine(w http.ResponseWriter, r *http.Request) {
    // 1. Extract trainer ID from JWT (set by middleware)
    trainerID := middleware.GetTrainerIDFromContext(r.Context())

    // 2. Parse request body
    var req dto.CreateRoutineRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    req.TrainerID = trainerID

    // 3. Execute use case
    response, err := h.createRoutineUseCase.Execute(r.Context(), &req)
    if err != nil {
        handleError(w, err)
        return
    }

    // 4. Return response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
```

---

## 6. Database Models vs Domain Entities

Database models and domain entities are kept **separate** to avoid coupling GORM tags or SQL concerns into the domain.

```go
// internal/adapter/persistence/postgres/model/routine_model.go
package model

// RoutineModel is the DB representation (has GORM tags)
type RoutineModel struct {
    ID        string `gorm:"primaryKey"`
    ClientID  string
    Name      string
    WeekStart time.Time
    Active    bool
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Convert to domain entity (used by repository when reading)
func (m *RoutineModel) ToDomainEntity() *entity.Routine {
    return &entity.Routine{
        ID:        m.ID,
        ClientID:  m.ClientID,
        Name:      m.Name,
        WeekStart: m.WeekStart,
        Active:    m.Active,
    }
}

// Create from domain entity (used by repository when writing)
func FromDomainEntity(routine *entity.Routine) *RoutineModel {
    return &RoutineModel{
        ID:        routine.ID,
        ClientID:  routine.ClientID,
        Name:      routine.Name,
        WeekStart: routine.WeekStart,
        Active:    routine.Active,
    }
}
```
