package sessions

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestUpdateExerciseLogUseCase_ShouldUpdateLog(t *testing.T) {
	logRepo := &mock.MockLogExerciseSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.LogExerciseSession, error) {
			return &entity.LogExerciseSession{
				ID: id, WorkoutSessionID: 1, RoutineDetailID: 1,
				SetNumber: 1, RepsDone: 8, Weight: 40,
			}, nil
		},
		UpdateFn: func(ctx context.Context, log *entity.LogExerciseSession) error {
			return nil
		},
	}

	uc := NewUpdateExerciseLogUseCase(logRepo)
	result, err := uc.Execute(context.Background(), 1, &dto.UpdateExerciseLogRequest{
		SetNumber: 1, RepsDone: 12, Weight: 55, Notes: "PR!",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.RepsDone != 12 {
		t.Errorf("Expected RepsDone 12, got %d", result.RepsDone)
	}
	if result.Weight != 55 {
		t.Errorf("Expected Weight 55, got %f", result.Weight)
	}
}

func TestUpdateExerciseLogUseCase_ShouldFailWhenNotFound(t *testing.T) {
	logRepo := &mock.MockLogExerciseSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.LogExerciseSession, error) {
			return nil, &domainerror.LogExerciseSessionNotFoundError{ID: id}
		},
	}

	uc := NewUpdateExerciseLogUseCase(logRepo)
	_, err := uc.Execute(context.Background(), 999, &dto.UpdateExerciseLogRequest{
		SetNumber: 1, RepsDone: 10, Weight: 50,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.LogExerciseSessionNotFoundError); !ok {
		t.Errorf("Expected LogExerciseSessionNotFoundError, got %T", err)
	}
}
