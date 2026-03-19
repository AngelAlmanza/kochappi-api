package templates

import (
	"context"
	"errors"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
)

func TestGetTemplatesUseCase_ShouldReturnAllTemplates(t *testing.T) {
	repo := &mock.MockTemplateRepository{
		GetAllFn: func(ctx context.Context) ([]entity.Template, error) {
			return []entity.Template{
				{ID: 1, Name: "PPL", Description: "Push Pull Legs"},
				{ID: 2, Name: "Full Body", Description: ""},
			}, nil
		},
	}

	useCase := NewGetTemplatesUseCase(repo)
	result, err := useCase.Execute(context.Background())

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 templates, got %d", len(result))
	}
	if result[0].Name != "PPL" {
		t.Errorf("Expected first template PPL, got %s", result[0].Name)
	}
}

func TestGetTemplatesUseCase_ShouldReturnEmptyList(t *testing.T) {
	repo := &mock.MockTemplateRepository{
		GetAllFn: func(ctx context.Context) ([]entity.Template, error) {
			return []entity.Template{}, nil
		},
	}

	useCase := NewGetTemplatesUseCase(repo)
	result, err := useCase.Execute(context.Background())

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected 0 templates, got %d", len(result))
	}
}

func TestGetTemplatesUseCase_ShouldPropagateRepositoryError(t *testing.T) {
	repo := &mock.MockTemplateRepository{
		GetAllFn: func(ctx context.Context) ([]entity.Template, error) {
			return nil, errors.New("db error")
		},
	}

	useCase := NewGetTemplatesUseCase(repo)
	_, err := useCase.Execute(context.Background())

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
