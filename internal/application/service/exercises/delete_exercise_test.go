package exercises

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestDeleteExerciseUseCase_ShouldDeleteExercise(t *testing.T) {
	deleted := false
	repo := &mock.MockExerciseRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Exercise, error) {
			return &entity.Exercise{ID: id, Name: "Squat"}, nil
		},
		DeleteFn: func(ctx context.Context, id int) error {
			deleted = true
			return nil
		},
	}

	useCase := NewDeleteExerciseUseCase(repo)
	err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !deleted {
		t.Error("Expected Delete to be called")
	}
}

func TestDeleteExerciseUseCase_ShouldReturnNotFoundError(t *testing.T) {
	repo := &mock.MockExerciseRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Exercise, error) {
			return nil, &domainerror.ExerciseNotFoundError{ID: id}
		},
	}

	useCase := NewDeleteExerciseUseCase(repo)
	err := useCase.Execute(context.Background(), 99)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ExerciseNotFoundError); !ok {
		t.Errorf("Expected ExerciseNotFoundError, got %T", err)
	}
}
