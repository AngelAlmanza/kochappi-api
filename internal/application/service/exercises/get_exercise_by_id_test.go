package exercises

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestGetExerciseByIDUseCase_ShouldReturnExercise(t *testing.T) {
	repo := &mock.MockExerciseRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Exercise, error) {
			return &entity.Exercise{ID: id, Name: "Squat", VideoURL: "https://example.com/squat.mp4"}, nil
		},
	}

	useCase := NewGetExerciseByIDUseCase(repo)
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 || result.Name != "Squat" {
		t.Errorf("Unexpected result: %+v", result)
	}
}

func TestGetExerciseByIDUseCase_ShouldReturnNotFoundError(t *testing.T) {
	repo := &mock.MockExerciseRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Exercise, error) {
			return nil, &domainerror.ExerciseNotFoundError{ID: id}
		},
	}

	useCase := NewGetExerciseByIDUseCase(repo)
	_, err := useCase.Execute(context.Background(), 99)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ExerciseNotFoundError); !ok {
		t.Errorf("Expected ExerciseNotFoundError, got %T", err)
	}
}
