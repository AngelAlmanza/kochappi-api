package templates

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestAddTemplateDetailUseCase_ShouldAddDetail(t *testing.T) {
	templateRepo := &mock.MockTemplateRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Template, error) {
			return &entity.Template{ID: id, Name: "PPL"}, nil
		},
	}
	exerciseRepo := &mock.MockExerciseRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Exercise, error) {
			return &entity.Exercise{ID: id, Name: "Squat"}, nil
		},
	}
	detailRepo := &mock.MockTemplateDetailRepository{
		CreateFn: func(ctx context.Context, detail *entity.TemplateDetail) error {
			detail.ID = 5
			return nil
		},
	}

	useCase := NewAddTemplateDetailUseCase(templateRepo, detailRepo, exerciseRepo)
	result, err := useCase.Execute(context.Background(), 1, &dto.AddTemplateDetailRequest{
		DayOfWeek:    2,
		ExerciseID:   10,
		Sets:         4,
		Reps:         8,
		DisplayOrder: 1,
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 5 || result.ExerciseID != 10 {
		t.Errorf("Unexpected result: %+v", result)
	}
}

func TestAddTemplateDetailUseCase_ShouldFailWhenTemplateNotFound(t *testing.T) {
	templateRepo := &mock.MockTemplateRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Template, error) {
			return nil, &domainerror.TemplateNotFoundError{ID: id}
		},
	}

	useCase := NewAddTemplateDetailUseCase(templateRepo, &mock.MockTemplateDetailRepository{}, &mock.MockExerciseRepository{})
	_, err := useCase.Execute(context.Background(), 99, &dto.AddTemplateDetailRequest{
		DayOfWeek: 1, ExerciseID: 1, Sets: 3, Reps: 10, DisplayOrder: 1,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.TemplateNotFoundError); !ok {
		t.Errorf("Expected TemplateNotFoundError, got %T", err)
	}
}

func TestAddTemplateDetailUseCase_ShouldFailWhenExerciseNotFound(t *testing.T) {
	templateRepo := &mock.MockTemplateRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Template, error) {
			return &entity.Template{ID: id}, nil
		},
	}
	exerciseRepo := &mock.MockExerciseRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Exercise, error) {
			return nil, &domainerror.ExerciseNotFoundError{ID: id}
		},
	}

	useCase := NewAddTemplateDetailUseCase(templateRepo, &mock.MockTemplateDetailRepository{}, exerciseRepo)
	_, err := useCase.Execute(context.Background(), 1, &dto.AddTemplateDetailRequest{
		DayOfWeek: 1, ExerciseID: 999, Sets: 3, Reps: 10, DisplayOrder: 1,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ExerciseNotFoundError); !ok {
		t.Errorf("Expected ExerciseNotFoundError, got %T", err)
	}
}
