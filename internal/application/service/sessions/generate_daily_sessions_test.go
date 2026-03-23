package sessions

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
)

func TestGenerateDailySessionsUseCase_ShouldCreateForActiveRoutines(t *testing.T) {
	// Wednesday 2026-01-14 → ISO day 3 (Wednesday)
	date := time.Date(2026, 1, 14, 0, 0, 0, 0, time.UTC)

	routineRepo := &mock.MockRoutineRepository{
		GetRoutinesToGenerateSessionsFn: func(ctx context.Context, dayOfWeek int16, d time.Time) ([]entity.Routine, error) {
			if dayOfWeek == 3 {
				return []entity.Routine{{ID: 1, Name: "Routine A", IsActive: true}}, nil
			}
			return nil, nil
		},
	}
	sessionRepo := &mock.MockWorkoutSessionRepository{
		CreateBulkFn: func(ctx context.Context, sessions []*entity.WorkoutSession) error {
			for i := range sessions {
				sessions[i].ID = i + 1
			}
			return nil
		},
	}

	uc := NewGenerateDailySessionsUseCase(routineRepo, sessionRepo)
	result, err := uc.Execute(context.Background(), date)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.SessionsCreated != 1 {
		t.Errorf("Expected 1 session created, got %d", result.SessionsCreated)
	}
}

func TestGenerateDailySessionsUseCase_ShouldSkipNoMatchDay(t *testing.T) {
	// Monday 2026-01-12 → ISO day 1 (Monday)
	date := time.Date(2026, 1, 12, 0, 0, 0, 0, time.UTC)

	routineRepo := &mock.MockRoutineRepository{
		GetRoutinesToGenerateSessionsFn: func(ctx context.Context, dayOfWeek int16, d time.Time) ([]entity.Routine, error) {
			return nil, nil // No routines match Monday
		},
	}

	uc := NewGenerateDailySessionsUseCase(routineRepo, &mock.MockWorkoutSessionRepository{})
	result, err := uc.Execute(context.Background(), date)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.SessionsCreated != 0 {
		t.Errorf("Expected 0 sessions created, got %d", result.SessionsCreated)
	}
}

func TestGenerateDailySessionsUseCase_ShouldSkipExistingSessions(t *testing.T) {
	date := time.Date(2026, 1, 14, 0, 0, 0, 0, time.UTC) // Wednesday

	routineRepo := &mock.MockRoutineRepository{
		GetRoutinesToGenerateSessionsFn: func(ctx context.Context, dayOfWeek int16, d time.Time) ([]entity.Routine, error) {
			return nil, nil // Query already filters out existing sessions
		},
	}

	uc := NewGenerateDailySessionsUseCase(routineRepo, &mock.MockWorkoutSessionRepository{})
	result, err := uc.Execute(context.Background(), date)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.SessionsCreated != 0 {
		t.Errorf("Expected 0 sessions created (idempotent), got %d", result.SessionsCreated)
	}
}

func TestGenerateDailySessionsUseCase_ShouldHandleZeroActiveRoutines(t *testing.T) {
	date := time.Date(2026, 1, 14, 0, 0, 0, 0, time.UTC)

	routineRepo := &mock.MockRoutineRepository{
		GetRoutinesToGenerateSessionsFn: func(ctx context.Context, dayOfWeek int16, d time.Time) ([]entity.Routine, error) {
			return nil, nil
		},
	}

	uc := NewGenerateDailySessionsUseCase(routineRepo, &mock.MockWorkoutSessionRepository{})
	result, err := uc.Execute(context.Background(), date)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.SessionsCreated != 0 {
		t.Errorf("Expected 0 sessions created, got %d", result.SessionsCreated)
	}
}
