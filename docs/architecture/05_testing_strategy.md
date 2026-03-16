# Testing Strategy

---

## Test Pyramid

```
        /\
       /  \          E2E Tests (5%)
      /────\         - Full API flow with real database
     /      \
    /────────\
   /          \      Integration Tests (25%)
  /            \     - Use cases + real repositories
 /──────────────\    - Containerized test database
/                \
/──────────────────\  Unit Tests (70%)
   Domain + Services  - Domain entities and value objects
                       - Use cases with mock repositories
```

---

## 1. Unit Tests — Use Cases

Test use cases by injecting mock repositories. No database required.

```go
// internal/application/service/routine/create_routine_test.go
package routine

import (
    "context"
    "testing"
    "kochappi/internal/application/dto"
    "kochappi/internal/domain/entity"
)

func TestCreateRoutineUseCase_ShouldCreateRoutineForOwnedClient(t *testing.T) {
    mockRoutineRepo := &MockRoutineRepository{
        CreateFn: func(ctx context.Context, routine *entity.Routine) error {
            return nil
        },
    }

    mockClientRepo := &MockClientRepository{
        GetByIDFn: func(ctx context.Context, id string) (*entity.Client, error) {
            return &entity.Client{
                ID:        "client-1",
                TrainerID: "trainer-1",
            }, nil
        },
    }

    useCase := NewCreateRoutineUseCase(mockRoutineRepo, mockClientRepo)

    response, err := useCase.Execute(context.Background(), &dto.CreateRoutineRequest{
        TrainerID: "trainer-1",
        ClientID:  "client-1",
        Name:      "Full Body",
    })

    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if response == nil {
        t.Errorf("Expected response, got nil")
    }
}

func TestCreateRoutineUseCase_ShouldFailForUnownedClient(t *testing.T) {
    mockClientRepo := &MockClientRepository{
        GetByIDFn: func(ctx context.Context, id string) (*entity.Client, error) {
            return &entity.Client{
                ID:        "client-1",
                TrainerID: "trainer-2", // Different trainer — should be rejected
            }, nil
        },
    }

    useCase := NewCreateRoutineUseCase(&MockRoutineRepository{}, mockClientRepo)

    _, err := useCase.Execute(context.Background(), &dto.CreateRoutineRequest{
        TrainerID: "trainer-1",
        ClientID:  "client-1",
    })

    if err == nil {
        t.Errorf("Expected authorization error, got nil")
    }
}
```

---

## 2. Integration Tests — Repositories

Test repositories against a real PostgreSQL database (containerized). These tests verify that SQL queries and migrations actually work.

```go
// test/integration/routine_test.go
package integration

import (
    "context"
    "testing"
    "kochappi/internal/adapter/persistence/postgres"
    "kochappi/internal/domain/entity"
    "kochappi/test"
)

func TestPostgresRoutineRepository_Create(t *testing.T) {
    db := test.SetupTestDB(t)
    defer test.TeardownTestDB(db)

    repo := postgres.NewRoutineRepository(db)

    routine := &entity.Routine{
        ID:       "routine-1",
        ClientID: "client-1",
        Name:     "Full Body",
    }

    err := repo.Create(context.Background(), routine)
    if err != nil {
        t.Fatalf("Failed to create routine: %v", err)
    }

    retrieved, err := repo.GetByID(context.Background(), "routine-1")
    if err != nil {
        t.Fatalf("Failed to retrieve routine: %v", err)
    }

    if retrieved.Name != "Full Body" {
        t.Errorf("Expected 'Full Body', got %s", retrieved.Name)
    }
}
```

---

## 3. E2E Tests — Full API Flow

Test complete request/response cycles via HTTP.

```go
// test/integration/api_test.go
package integration

import (
    "bytes"
    "encoding/json"
    "net/http"
    "testing"
    "kochappi/test"
)

func TestCreateRoutineE2E(t *testing.T) {
    server := test.SetupTestServer(t)
    defer server.Close()

    trainer := test.CreateTrainerFixture(t, server)
    token := test.GenerateJWT(trainer.ID)
    client := test.CreateClientFixture(t, server, trainer.ID)

    requestBody := map[string]interface{}{
        "client_id": client.ID,
        "name":      "Full Body Workout",
        "exercises": []map[string]interface{}{
            {
                "name":      "Bench Press",
                "sets":      3,
                "reps":      8,
                "video_url": "https://youtube.com/...",
            },
        },
    }

    body, _ := json.Marshal(requestBody)

    req, _ := http.NewRequest("POST", server.URL+"/api/v1/routines", bytes.NewBuffer(body))
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

    resp, _ := http.DefaultClient.Do(req)

    if resp.StatusCode != http.StatusCreated {
        t.Errorf("Expected 201, got %d", resp.StatusCode)
    }

    var response map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&response)

    if response["id"] == nil {
        t.Errorf("Expected routine ID in response")
    }
}
```

---

## 4. Test Helpers & Fixtures

```go
// test/testhelpers.go
package test

func SetupTestDB(t *testing.T) *gorm.DB {
    db, err := gorm.Open(postgres.Open(testDBURL))
    if err != nil {
        t.Fatalf("Failed to connect to test DB: %v", err)
    }
    runMigrations(db)
    return db
}

func TeardownTestDB(db *gorm.DB) {
    // Clean up data after test
}

// test/fixtures/trainer_fixture.go
func CreateTrainerFixture(t *testing.T, db *gorm.DB) *entity.Trainer {
    trainer := &entity.Trainer{
        ID:    "trainer-test-" + uuid.New().String(),
        Name:  "John Doe",
        Email: "john@example.com",
    }
    if err := db.Create(trainer).Error; err != nil {
        t.Fatalf("Failed to create trainer fixture: %v", err)
    }
    return trainer
}
```

---

## 5. Test Database (docker-compose.test.yml)

```yaml
version: '3.8'

services:
  postgres-test:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: kochappi_test
      POSTGRES_PASSWORD: password
      POSTGRES_DB: kochappi_test
    ports:
      - "5433:5432"
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U kochappi_test"]
      interval: 10s
      timeout: 5s
      retries: 5
```

Integration tests hit port `5433` so they don't conflict with the dev database on `5432`.

---

## 6. Running Tests

```bash
# Unit tests only (fast, no DB required)
go test ./internal/... -v

# Integration tests (requires Docker)
go test ./test/integration/... -v

# All tests with coverage report
go test ./... -v -cover

# Run a specific test by name
go test -run TestCreateRoutineUseCase ./internal/application/service/routine

# With race condition detector
go test ./... -race

# Generate HTML coverage report
go test ./... -cover -coverprofile=coverage.out
go tool cover -html=coverage.out
```
