package sessions

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestCreateExerciseLogUseCase_ShouldCreateLog(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return &entity.WorkoutSession{
				ID: id, Status: entity.WorkoutStatusInProgress,
				ActualDate: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			}, nil
		},
	}
	logRepo := &mock.MockLogExerciseSessionRepository{
		CreateFn: func(ctx context.Context, log *entity.LogExerciseSession) error {
			log.ID = 1
			return nil
		},
	}
	detailRepo := &mock.MockRoutineDetailRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.RoutineDetail, error) {
			return &entity.RoutineDetail{ID: id}, nil
		},
	}

	uc := NewCreateExerciseLogUseCase(sessionRepo, logRepo, detailRepo)
	result, err := uc.Execute(context.Background(), 1, &dto.CreateExerciseLogRequest{
		RoutineDetailID: 1,
		SetNumber:       1,
		RepsDone:        10,
		Weight:          50.0,
		Notes:           "Good set",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 {
		t.Errorf("Expected ID 1, got %d", result.ID)
	}
	if result.RepsDone != 10 {
		t.Errorf("Expected RepsDone 10, got %d", result.RepsDone)
	}
}

func TestCreateExerciseLogUseCase_ShouldFailWhenSessionNotFound(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return nil, &domainerror.WorkoutSessionNotFoundError{ID: id}
		},
	}

	uc := NewCreateExerciseLogUseCase(sessionRepo, &mock.MockLogExerciseSessionRepository{}, &mock.MockRoutineDetailRepository{})
	_, err := uc.Execute(context.Background(), 999, &dto.CreateExerciseLogRequest{
		RoutineDetailID: 1, SetNumber: 1, RepsDone: 10, Weight: 50,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestCreateExerciseLogUseCase_ShouldFailWhenSessionNotInProgress(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return &entity.WorkoutSession{
				ID: id, Status: entity.WorkoutStatusPending,
				ActualDate: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			}, nil
		},
	}

	uc := NewCreateExerciseLogUseCase(sessionRepo, &mock.MockLogExerciseSessionRepository{}, &mock.MockRoutineDetailRepository{})
	_, err := uc.Execute(context.Background(), 1, &dto.CreateExerciseLogRequest{
		RoutineDetailID: 1, SetNumber: 1, RepsDone: 10, Weight: 50,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.InvalidSessionStatusTransitionError); !ok {
		t.Errorf("Expected InvalidSessionStatusTransitionError, got %T", err)
	}
}

func TestCreateExerciseLogUseCase_ShouldFailWhenRoutineDetailNotFound(t *testing.T) {
	sessionRepo := &mock.MockWorkoutSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.WorkoutSession, error) {
			return &entity.WorkoutSession{
				ID: id, Status: entity.WorkoutStatusInProgress,
				ActualDate: time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			}, nil
		},
	}
	detailRepo := &mock.MockRoutineDetailRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.RoutineDetail, error) {
			return nil, &domainerror.RoutineDetailNotFoundError{ID: id}
		},
	}

	uc := NewCreateExerciseLogUseCase(sessionRepo, &mock.MockLogExerciseSessionRepository{}, detailRepo)
	_, err := uc.Execute(context.Background(), 1, &dto.CreateExerciseLogRequest{
		RoutineDetailID: 999, SetNumber: 1, RepsDone: 10, Weight: 50,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.RoutineDetailNotFoundError); !ok {
		t.Errorf("Expected RoutineDetailNotFoundError, got %T", err)
	}
}
