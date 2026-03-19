package templates

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestRemoveTemplateDetailUseCase_ShouldRemoveDetail(t *testing.T) {
	deleted := false
	detailRepo := &mock.MockTemplateDetailRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.TemplateDetail, error) {
			return &entity.TemplateDetail{ID: id, TemplateID: 1}, nil
		},
		DeleteByIDFn: func(ctx context.Context, id int) error {
			deleted = true
			return nil
		},
	}

	useCase := NewRemoveTemplateDetailUseCase(detailRepo)
	err := useCase.Execute(context.Background(), 1, 5)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !deleted {
		t.Error("Expected DeleteByID to be called")
	}
}

func TestRemoveTemplateDetailUseCase_ShouldReturnNotFoundWhenDetailDoesNotExist(t *testing.T) {
	detailRepo := &mock.MockTemplateDetailRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.TemplateDetail, error) {
			return nil, &domainerror.TemplateDetailNotFoundError{ID: id}
		},
	}

	useCase := NewRemoveTemplateDetailUseCase(detailRepo)
	err := useCase.Execute(context.Background(), 1, 99)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.TemplateDetailNotFoundError); !ok {
		t.Errorf("Expected TemplateDetailNotFoundError, got %T", err)
	}
}

func TestRemoveTemplateDetailUseCase_ShouldReturnNotFoundWhenDetailBelongsToDifferentTemplate(t *testing.T) {
	detailRepo := &mock.MockTemplateDetailRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.TemplateDetail, error) {
			// Detail belongs to template 2, not 1
			return &entity.TemplateDetail{ID: id, TemplateID: 2}, nil
		},
	}

	useCase := NewRemoveTemplateDetailUseCase(detailRepo)
	err := useCase.Execute(context.Background(), 1, 5)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.TemplateDetailNotFoundError); !ok {
		t.Errorf("Expected TemplateDetailNotFoundError, got %T", err)
	}
}
