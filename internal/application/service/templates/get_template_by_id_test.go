package templates

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestGetTemplateByIDUseCase_ShouldReturnTemplateWithDetails(t *testing.T) {
	templateRepo := &mock.MockTemplateRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Template, error) {
			return &entity.Template{ID: id, Name: "PPL", Description: "Push Pull Legs"}, nil
		},
	}
	detailRepo := &mock.MockTemplateDetailRepository{
		GetByTemplateIDFn: func(ctx context.Context, templateID int) ([]entity.TemplateDetail, error) {
			return []entity.TemplateDetail{
				{ID: 1, TemplateID: templateID, DayOfWeek: 1, ExerciseID: 10, Sets: 3, Reps: 10, DisplayOrder: 1},
				{ID: 2, TemplateID: templateID, DayOfWeek: 1, ExerciseID: 11, Sets: 4, Reps: 8, DisplayOrder: 2},
			}, nil
		},
	}

	useCase := NewGetTemplateByIDUseCase(templateRepo, detailRepo)
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 || result.Name != "PPL" {
		t.Errorf("Unexpected template: %+v", result)
	}
	if len(result.Details) != 2 {
		t.Errorf("Expected 2 details, got %d", len(result.Details))
	}
}

func TestGetTemplateByIDUseCase_ShouldReturnTemplateWithEmptyDetails(t *testing.T) {
	templateRepo := &mock.MockTemplateRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Template, error) {
			return &entity.Template{ID: id, Name: "Empty"}, nil
		},
	}
	detailRepo := &mock.MockTemplateDetailRepository{
		GetByTemplateIDFn: func(ctx context.Context, templateID int) ([]entity.TemplateDetail, error) {
			return []entity.TemplateDetail{}, nil
		},
	}

	useCase := NewGetTemplateByIDUseCase(templateRepo, detailRepo)
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Details) != 0 {
		t.Errorf("Expected 0 details, got %d", len(result.Details))
	}
}

func TestGetTemplateByIDUseCase_ShouldReturnNotFoundError(t *testing.T) {
	templateRepo := &mock.MockTemplateRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Template, error) {
			return nil, &domainerror.TemplateNotFoundError{ID: id}
		},
	}

	useCase := NewGetTemplateByIDUseCase(templateRepo, &mock.MockTemplateDetailRepository{})
	_, err := useCase.Execute(context.Background(), 99)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.TemplateNotFoundError); !ok {
		t.Errorf("Expected TemplateNotFoundError, got %T", err)
	}
}
