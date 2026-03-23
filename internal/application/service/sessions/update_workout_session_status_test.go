package sessions

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestUpdateWorkoutSessionStatusUseCase_ShouldTransitionPendingToInProgress(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return &entity.WorkoutSession{
				ID: id, Status: entity.WorkoutStatusPending,
				ActualDate: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			}, nil
		},
		UpdateFn: func(ctx context.Context, session *entity.WorkoutSession) error {
			return nil
		},
	}

	uc := NewUpdateWorkoutSessionStatusUseCase(sessionRepo)
	result, err := uc.Execute(context.Background(), 1, "in_progress")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Status != "in_progress" {
		t.Errorf("Expected status in_progress, got %s", result.Status)
	}
}

func TestUpdateWorkoutSessionStatusUseCase_ShouldTransitionInProgressToCompleted(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return &entity.WorkoutSession{
				ID: id, Status: entity.WorkoutStatusInProgress,
				ActualDate: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			}, nil
		},
		UpdateFn: func(ctx context.Context, session *entity.WorkoutSession) error {
			return nil
		},
	}

	uc := NewUpdateWorkoutSessionStatusUseCase(sessionRepo)
	result, err := uc.Execute(context.Background(), 1, "completed")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Status != "completed" {
		t.Errorf("Expected status completed, got %s", result.Status)
	}
}

func TestUpdateWorkoutSessionStatusUseCase_ShouldTransitionPendingToSkipped(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return &entity.WorkoutSession{
				ID: id, Status: entity.WorkoutStatusPending,
				ActualDate: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			}, nil
		},
		UpdateFn: func(ctx context.Context, session *entity.WorkoutSession) error { return nil },
	}

	uc := NewUpdateWorkoutSessionStatusUseCase(sessionRepo)
	result, err := uc.Execute(context.Background(), 1, "skipped")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Status != "skipped" {
		t.Errorf("Expected status skipped, got %s", result.Status)
	}
}

func TestUpdateWorkoutSessionStatusUseCase_ShouldFailOnInvalidTransition(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return &entity.WorkoutSession{
				ID: id, Status: entity.WorkoutStatusCompleted,
				ActualDate: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			}, nil
		},
	}

	uc := NewUpdateWorkoutSessionStatusUseCase(sessionRepo)
	_, err := uc.Execute(context.Background(), 1, "skipped")

	if err == nil {
		t.Fatal("Expected error for invalid transition, got nil")
	}
	if _, ok := err.(*domainerror.InvalidSessionStatusTransitionError); !ok {
		t.Errorf("Expected InvalidSessionStatusTransitionError, got %T", err)
	}
}

func TestUpdateWorkoutSessionStatusUseCase_ShouldFailWhenNotFound(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return nil, &domainerror.WorkoutSessionNotFoundError{ID: id}
		},
	}

	uc := NewUpdateWorkoutSessionStatusUseCase(sessionRepo)
	_, err := uc.Execute(context.Background(), 999, "in_progress")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.WorkoutSessionNotFoundError); !ok {
		t.Errorf("Expected WorkoutSessionNotFoundError, got %T", err)
	}
}
