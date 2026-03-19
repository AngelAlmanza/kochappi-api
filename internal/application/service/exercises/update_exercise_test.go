package exercises

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestUpdateExerciseUseCase_ShouldUpdateExercise(t *testing.T) {
	repo := &mock.MockExerciseRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Exercise, error) {
			return &entity.Exercise{ID: id, Name: "Old Name", VideoURL: ""}, nil
		},
		UpdateFn: func(ctx context.Context, exercise *entity.Exercise) error {
			return nil
		},
	}

	useCase := NewUpdateExerciseUseCase(repo)
	result, err := useCase.Execute(context.Background(), 1, &dto.UpdateExerciseRequest{
		Name:     "New Name",
		VideoURL: "https://example.com/new.mp4",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Name != "New Name" {
		t.Errorf("Expected name 'New Name', got %s", result.Name)
	}
	if result.VideoURL != "https://example.com/new.mp4" {
		t.Errorf("Unexpected VideoURL: %s", result.VideoURL)
	}
}

func TestUpdateExerciseUseCase_ShouldReturnNotFoundError(t *testing.T) {
	repo := &mock.MockExerciseRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Exercise, error) {
			return nil, &domainerror.ExerciseNotFoundError{ID: id}
		},
	}

	useCase := NewUpdateExerciseUseCase(repo)
	_, err := useCase.Execute(context.Background(), 99, &dto.UpdateExerciseRequest{Name: "X"})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ExerciseNotFoundError); !ok {
		t.Errorf("Expected ExerciseNotFoundError, got %T", err)
	}
}
