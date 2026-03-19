package exercises

import (
	"context"
	"errors"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
)

func TestGetExercisesUseCase_ShouldReturnAllExercises(t *testing.T) {
	repo := &mock.MockExerciseRepository{
		GetAllFn: func(ctx context.Context) ([]entity.Exercise, error) {
			return []entity.Exercise{
				{ID: 1, Name: "Squat", VideoURL: "https://example.com/squat.mp4"},
				{ID: 2, Name: "Deadlift", VideoURL: ""},
			}, nil
		},
	}

	useCase := NewGetExercisesUseCase(repo)
	result, err := useCase.Execute(context.Background())

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 exercises, got %d", len(result))
	}
	if result[0].ID != 1 || result[0].Name != "Squat" {
		t.Errorf("Unexpected first exercise: %+v", result[0])
	}
}

func TestGetExercisesUseCase_ShouldReturnEmptyList(t *testing.T) {
	repo := &mock.MockExerciseRepository{
		GetAllFn: func(ctx context.Context) ([]entity.Exercise, error) {
			return []entity.Exercise{}, nil
		},
	}

	useCase := NewGetExercisesUseCase(repo)
	result, err := useCase.Execute(context.Background())

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected 0 exercises, got %d", len(result))
	}
}

func TestGetExercisesUseCase_ShouldPropagateRepositoryError(t *testing.T) {
	repo := &mock.MockExerciseRepository{
		GetAllFn: func(ctx context.Context) ([]entity.Exercise, error) {
			return nil, errors.New("db error")
		},
	}

	useCase := NewGetExercisesUseCase(repo)
	_, err := useCase.Execute(context.Background())

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
