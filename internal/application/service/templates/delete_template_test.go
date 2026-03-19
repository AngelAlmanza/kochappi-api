package templates

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestDeleteTemplateUseCase_ShouldDeleteTemplate(t *testing.T) {
	deleted := false
	repo := &mock.MockTemplateRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Template, error) {
			return &entity.Template{ID: id, Name: "PPL"}, nil
		},
		DeleteFn: func(ctx context.Context, id int) error {
			deleted = true
			return nil
		},
	}

	useCase := NewDeleteTemplateUseCase(repo)
	err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !deleted {
		t.Error("Expected Delete to be called")
	}
}

func TestDeleteTemplateUseCase_ShouldReturnNotFoundError(t *testing.T) {
	repo := &mock.MockTemplateRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Template, error) {
			return nil, &domainerror.TemplateNotFoundError{ID: id}
		},
	}

	useCase := NewDeleteTemplateUseCase(repo)
	err := useCase.Execute(context.Background(), 99)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.TemplateNotFoundError); !ok {
		t.Errorf("Expected TemplateNotFoundError, got %T", err)
	}
}
