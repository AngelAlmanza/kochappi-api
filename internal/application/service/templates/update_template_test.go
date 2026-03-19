package templates

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestUpdateTemplateUseCase_ShouldUpdateTemplate(t *testing.T) {
	repo := &mock.MockTemplateRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Template, error) {
			return &entity.Template{ID: id, Name: "Old Name", Description: "Old desc"}, nil
		},
		UpdateFn: func(ctx context.Context, template *entity.Template) error {
			return nil
		},
	}

	useCase := NewUpdateTemplateUseCase(repo)
	result, err := useCase.Execute(context.Background(), 1, &dto.UpdateTemplateRequest{
		Name:        "New Name",
		Description: "New desc",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Name != "New Name" || result.Description != "New desc" {
		t.Errorf("Unexpected result: %+v", result)
	}
}

func TestUpdateTemplateUseCase_ShouldReturnNotFoundError(t *testing.T) {
	repo := &mock.MockTemplateRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Template, error) {
			return nil, &domainerror.TemplateNotFoundError{ID: id}
		},
	}

	useCase := NewUpdateTemplateUseCase(repo)
	_, err := useCase.Execute(context.Background(), 99, &dto.UpdateTemplateRequest{Name: "X"})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.TemplateNotFoundError); !ok {
		t.Errorf("Expected TemplateNotFoundError, got %T", err)
	}
}
