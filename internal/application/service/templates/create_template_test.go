package templates

import (
	"context"
	"errors"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestCreateTemplateUseCase_ShouldCreateTemplateWithoutDetails(t *testing.T) {
	templateRepo := &mock.MockTemplateRepository{
		CreateFn: func(ctx context.Context, template *entity.Template) error {
			template.ID = 1
			return nil
		},
	}

	useCase := NewCreateTemplateUseCase(templateRepo, &mock.MockTemplateDetailRepository{}, &mock.MockExerciseRepository{})
	result, err := useCase.Execute(context.Background(), &dto.CreateTemplateRequest{
		Name:        "PPL",
		Description: "Push Pull Legs",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 || result.Name != "PPL" {
		t.Errorf("Unexpected result: %+v", result)
	}
	if len(result.Details) != 0 {
		t.Errorf("Expected 0 details, got %d", len(result.Details))
	}
}

func TestCreateTemplateUseCase_ShouldCreateTemplateWithDetails(t *testing.T) {
	templateRepo := &mock.MockTemplateRepository{
		CreateFn: func(ctx context.Context, template *entity.Template) error {
			template.ID = 1
			return nil
		},
	}
	detailRepo := &mock.MockTemplateDetailRepository{
		CreateBulkFn: func(ctx context.Context, details []*entity.TemplateDetail) error {
			for i, d := range details {
				d.ID = i + 10
			}
			return nil
		},
	}
	exerciseRepo := &mock.MockExerciseRepository{
		GetByIDsFn: func(ctx context.Context, ids []int) ([]entity.Exercise, error) {
			exercises := make([]entity.Exercise, 0, len(ids))
			for _, id := range ids {
				exercises = append(exercises, entity.Exercise{ID: id, Name: "Exercise"})
			}
			return exercises, nil
		},
	}

	useCase := NewCreateTemplateUseCase(templateRepo, detailRepo, exerciseRepo)
	result, err := useCase.Execute(context.Background(), &dto.CreateTemplateRequest{
		Name: "PPL",
		Details: []dto.CreateTemplateDetailRequest{
			{DayOfWeek: 1, ExerciseID: 5, Sets: 3, Reps: 10, DisplayOrder: 1},
			{DayOfWeek: 1, ExerciseID: 6, Sets: 4, Reps: 8, DisplayOrder: 2},
		},
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Details) != 2 {
		t.Errorf("Expected 2 details, got %d", len(result.Details))
	}
}

func TestCreateTemplateUseCase_ShouldDeduplicateExerciseIDsInValidation(t *testing.T) {
	callCount := 0
	receivedIDs := []int{}

	exerciseRepo := &mock.MockExerciseRepository{
		GetByIDsFn: func(ctx context.Context, ids []int) ([]entity.Exercise, error) {
			callCount++
			receivedIDs = ids
			exercises := make([]entity.Exercise, 0, len(ids))
			for _, id := range ids {
				exercises = append(exercises, entity.Exercise{ID: id})
			}
			return exercises, nil
		},
	}
	templateRepo := &mock.MockTemplateRepository{
		CreateFn: func(ctx context.Context, template *entity.Template) error {
			template.ID = 1
			return nil
		},
	}

	useCase := NewCreateTemplateUseCase(templateRepo, &mock.MockTemplateDetailRepository{}, exerciseRepo)
	_, err := useCase.Execute(context.Background(), &dto.CreateTemplateRequest{
		Name: "PPL",
		Details: []dto.CreateTemplateDetailRequest{
			{DayOfWeek: 1, ExerciseID: 5, Sets: 3, Reps: 10, DisplayOrder: 1},
			{DayOfWeek: 2, ExerciseID: 5, Sets: 3, Reps: 10, DisplayOrder: 1}, // duplicate ExerciseID
			{DayOfWeek: 1, ExerciseID: 6, Sets: 4, Reps: 8, DisplayOrder: 2},
		},
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if callCount != 1 {
		t.Errorf("Expected GetByIDs to be called once, got %d", callCount)
	}
	if len(receivedIDs) != 2 {
		t.Errorf("Expected 2 unique IDs sent to GetByIDs, got %d: %v", len(receivedIDs), receivedIDs)
	}
}

func TestCreateTemplateUseCase_ShouldCallCreateBulkOnce(t *testing.T) {
	bulkCallCount := 0

	templateRepo := &mock.MockTemplateRepository{
		CreateFn: func(ctx context.Context, template *entity.Template) error {
			template.ID = 1
			return nil
		},
	}
	detailRepo := &mock.MockTemplateDetailRepository{
		CreateBulkFn: func(ctx context.Context, details []*entity.TemplateDetail) error {
			bulkCallCount++
			return nil
		},
	}
	exerciseRepo := &mock.MockExerciseRepository{
		GetByIDsFn: func(ctx context.Context, ids []int) ([]entity.Exercise, error) {
			exercises := make([]entity.Exercise, 0, len(ids))
			for _, id := range ids {
				exercises = append(exercises, entity.Exercise{ID: id})
			}
			return exercises, nil
		},
	}

	useCase := NewCreateTemplateUseCase(templateRepo, detailRepo, exerciseRepo)
	_, err := useCase.Execute(context.Background(), &dto.CreateTemplateRequest{
		Name: "PPL",
		Details: []dto.CreateTemplateDetailRequest{
			{DayOfWeek: 1, ExerciseID: 1, Sets: 3, Reps: 10, DisplayOrder: 1},
			{DayOfWeek: 2, ExerciseID: 2, Sets: 4, Reps: 8, DisplayOrder: 1},
			{DayOfWeek: 3, ExerciseID: 3, Sets: 3, Reps: 12, DisplayOrder: 1},
		},
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if bulkCallCount != 1 {
		t.Errorf("Expected CreateBulk to be called once, got %d", bulkCallCount)
	}
}

func TestCreateTemplateUseCase_ShouldFailWhenExerciseNotFound(t *testing.T) {
	exerciseRepo := &mock.MockExerciseRepository{
		GetByIDsFn: func(ctx context.Context, ids []int) ([]entity.Exercise, error) {
			return nil, &domainerror.ExerciseNotFoundError{ID: ids[0]}
		},
	}

	useCase := NewCreateTemplateUseCase(&mock.MockTemplateRepository{}, &mock.MockTemplateDetailRepository{}, exerciseRepo)
	_, err := useCase.Execute(context.Background(), &dto.CreateTemplateRequest{
		Name: "PPL",
		Details: []dto.CreateTemplateDetailRequest{
			{DayOfWeek: 1, ExerciseID: 999, Sets: 3, Reps: 10, DisplayOrder: 1},
		},
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ExerciseNotFoundError); !ok {
		t.Errorf("Expected ExerciseNotFoundError, got %T", err)
	}
}

func TestCreateTemplateUseCase_ShouldNotCallGetByIDsWhenNoDetails(t *testing.T) {
	called := false
	exerciseRepo := &mock.MockExerciseRepository{
		GetByIDsFn: func(ctx context.Context, ids []int) ([]entity.Exercise, error) {
			called = true
			return nil, errors.New("should not be called")
		},
	}
	templateRepo := &mock.MockTemplateRepository{
		CreateFn: func(ctx context.Context, template *entity.Template) error {
			template.ID = 1
			return nil
		},
	}

	useCase := NewCreateTemplateUseCase(templateRepo, &mock.MockTemplateDetailRepository{}, exerciseRepo)
	_, err := useCase.Execute(context.Background(), &dto.CreateTemplateRequest{Name: "PPL"})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if called {
		t.Error("GetByIDs should not be called when there are no details")
	}
}
