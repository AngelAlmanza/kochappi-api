package sessions

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestDeleteExerciseLogUseCase_ShouldDeleteLog(t *testing.T) {
	deleted := false
	logRepo := &mock.MockLogExerciseSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.LogExerciseSession, error) {
			return &entity.LogExerciseSession{ID: id}, nil
		},
		DeleteFn: func(ctx context.Context, id int) error {
			deleted = true
			return nil
		},
	}

	uc := NewDeleteExerciseLogUseCase(logRepo)
	err := uc.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !deleted {
		t.Error("Expected delete to be called")
	}
}

func TestDeleteExerciseLogUseCase_ShouldFailWhenNotFound(t *testing.T) {
	logRepo := &mock.MockLogExerciseSessionRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.LogExerciseSession, error) {
			return nil, &domainerror.LogExerciseSessionNotFoundError{ID: id}
		},
	}

	uc := NewDeleteExerciseLogUseCase(logRepo)
	err := uc.Execute(context.Background(), 999)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.LogExerciseSessionNotFoundError); !ok {
		t.Errorf("Expected LogExerciseSessionNotFoundError, got %T", err)
	}
}
