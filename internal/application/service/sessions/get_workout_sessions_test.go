package sessions

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestGetWorkoutSessionsUseCase_ShouldReturnAllByRoutine(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id, Name: "Test"}, nil
		},
	}
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByCriteriaFn: func(ctx context.Context, criteria port.WorkoutSessionCriteria) ([]entity.WorkoutSession, error) {
			return []entity.WorkoutSession{
				{ID: 1, RoutineID: *criteria.RoutineID, Status: entity.WorkoutStatusPending, ActualDate: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)},
				{ID: 2, RoutineID: *criteria.RoutineID, Status: entity.WorkoutStatusCompleted, ActualDate: time.Date(2026, 1, 14, 0, 0, 0, 0, time.UTC)},
			}, nil
		},
	}

	uc := NewGetWorkoutSessionsUseCase(routineRepo, sessionRepo)
	result, err := uc.Execute(context.Background(), 1, nil, nil, nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 sessions, got %d", len(result))
	}
}

func TestGetWorkoutSessionsUseCase_ShouldFilterByDateRange(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id}, nil
		},
	}
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByCriteriaFn: func(ctx context.Context, criteria port.WorkoutSessionCriteria) ([]entity.WorkoutSession, error) {
			return []entity.WorkoutSession{
				{ID: 1, RoutineID: *criteria.RoutineID, ActualDate: *criteria.DateFrom},
			}, nil
		},
	}

	uc := NewGetWorkoutSessionsUseCase(routineRepo, sessionRepo)
	from := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2026, 1, 31, 0, 0, 0, 0, time.UTC)
	result, err := uc.Execute(context.Background(), 1, nil, &from, &to)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 1 session, got %d", len(result))
	}
}

func TestGetWorkoutSessionsUseCase_ShouldFilterByStatus(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id}, nil
		},
	}
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByCriteriaFn: func(ctx context.Context, criteria port.WorkoutSessionCriteria) ([]entity.WorkoutSession, error) {
			return []entity.WorkoutSession{
				{ID: 1, RoutineID: *criteria.RoutineID, Status: entity.WorkoutStatus(*criteria.Status), ActualDate: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)},
			}, nil
		},
	}

	uc := NewGetWorkoutSessionsUseCase(routineRepo, sessionRepo)
	status := "pending"
	result, err := uc.Execute(context.Background(), 1, &status, nil, nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 1 session, got %d", len(result))
	}
}

func TestGetWorkoutSessionsUseCase_ShouldFailWhenRoutineNotFound(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return nil, &domainerror.RoutineNotFoundError{ID: id}
		},
	}

	uc := NewGetWorkoutSessionsUseCase(routineRepo, &mock.MockWorkoutSessionRepository{})
	_, err := uc.Execute(context.Background(), 999, nil, nil, nil)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.RoutineNotFoundError); !ok {
		t.Errorf("Expected RoutineNotFoundError, got %T", err)
	}
}
