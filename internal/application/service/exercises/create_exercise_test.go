package exercises

import (
	"context"
	"errors"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
)

func TestCreateExerciseUseCase_ShouldCreateExercise(t *testing.T) {
	repo := &mock.MockExerciseRepository{
		CreateFn: func(ctx context.Context, exercise *entity.Exercise) error {
			exercise.ID = 1
			return nil
		},
	}

	useCase := NewCreateExerciseUseCase(repo)
	result, err := useCase.Execute(context.Background(), &dto.CreateExerciseRequest{
		Name:     "Squat",
		VideoURL: "https://example.com/squat.mp4",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 {
		t.Errorf("Expected ID 1, got %d", result.ID)
	}
	if result.Name != "Squat" {
		t.Errorf("Expected name Squat, got %s", result.Name)
	}
	if result.VideoURL != "https://example.com/squat.mp4" {
		t.Errorf("Unexpected VideoURL: %s", result.VideoURL)
	}
}

func TestCreateExerciseUseCase_ShouldCreateExerciseWithoutVideoURL(t *testing.T) {
	repo := &mock.MockExerciseRepository{
		CreateFn: func(ctx context.Context, exercise *entity.Exercise) error {
			exercise.ID = 2
			return nil
		},
	}

	useCase := NewCreateExerciseUseCase(repo)
	result, err := useCase.Execute(context.Background(), &dto.CreateExerciseRequest{
		Name: "Deadlift",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.VideoURL != "" {
		t.Errorf("Expected empty VideoURL, got %s", result.VideoURL)
	}
}

func TestCreateExerciseUseCase_ShouldPropagateRepositoryError(t *testing.T) {
	repo := &mock.MockExerciseRepository{
		CreateFn: func(ctx context.Context, exercise *entity.Exercise) error {
			return errors.New("db error")
		},
	}

	useCase := NewCreateExerciseUseCase(repo)
	_, err := useCase.Execute(context.Background(), &dto.CreateExerciseRequest{Name: "Squat"})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
