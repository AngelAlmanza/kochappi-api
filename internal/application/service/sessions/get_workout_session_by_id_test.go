package sessions

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestGetWorkoutSessionByIDUseCase_ShouldReturnWithLogs(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return &entity.WorkoutSession{
				ID: id, RoutineID: 1, Status: entity.WorkoutStatusInProgress,
				ActualDate: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			}, nil
		},
	}
	logRepo := &mock.MockLogExerciseSessionRepository{
		GetByWorkoutSessionIDFn: func(ctx context.Context, workoutSessionID int) ([]entity.LogExerciseSession, error) {
			return []entity.LogExerciseSession{
				{ID: 10, WorkoutSessionID: workoutSessionID, RoutineDetailID: 1, SetNumber: 1, RepsDone: 10, Weight: 50},
			}, nil
		},
	}

	uc := NewGetWorkoutSessionByIDUseCase(sessionRepo, logRepo)
	result, err := uc.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 {
		t.Errorf("Expected ID 1, got %d", result.ID)
	}
	if len(result.ExerciseLogs) != 1 {
		t.Errorf("Expected 1 log, got %d", len(result.ExerciseLogs))
	}
}

func TestGetWorkoutSessionByIDUseCase_ShouldReturnEmptyLogs(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return &entity.WorkoutSession{
				ID: id, RoutineID: 1, Status: entity.WorkoutStatusPending,
				ActualDate: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			}, nil
		},
	}
	logRepo := &mock.MockLogExerciseSessionRepository{
		GetByWorkoutSessionIDFn: func(ctx context.Context, workoutSessionID int) ([]entity.LogExerciseSession, error) {
			return []entity.LogExerciseSession{}, nil
		},
	}

	uc := NewGetWorkoutSessionByIDUseCase(sessionRepo, logRepo)
	result, err := uc.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.ExerciseLogs) != 0 {
		t.Errorf("Expected 0 logs, got %d", len(result.ExerciseLogs))
	}
}

func TestGetWorkoutSessionByIDUseCase_ShouldFailWhenNotFound(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return nil, &domainerror.WorkoutSessionNotFoundError{ID: id}
		},
	}

	uc := NewGetWorkoutSessionByIDUseCase(sessionRepo, &mock.MockLogExerciseSessionRepository{})
	_, err := uc.Execute(context.Background(), 999)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.WorkoutSessionNotFoundError); !ok {
		t.Errorf("Expected WorkoutSessionNotFoundError, got %T", err)
	}
}
