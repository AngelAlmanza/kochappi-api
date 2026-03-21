package routines

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestRemoveRoutineDetailUseCase_ShouldRemoveDetail(t *testing.T) {
	deleted := false
	detailRepo := &mock.MockRoutineDetailRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.RoutineDetail, error) {
			return &entity.RoutineDetail{ID: id, RoutineID: 1}, nil
		},
		DeleteByIDFn: func(ctx context.Context, id int) error {
			deleted = true
			return nil
		},
	}

	useCase := NewRemoveRoutineDetailUseCase(detailRepo)
	err := useCase.Execute(context.Background(), 1, 10)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !deleted {
		t.Error("Expected detail to be deleted")
	}
}

func TestRemoveRoutineDetailUseCase_ShouldFailWhenDetailNotFound(t *testing.T) {
	detailRepo := &mock.MockRoutineDetailRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.RoutineDetail, error) {
			return nil, &domainerror.RoutineDetailNotFoundError{ID: id}
		},
	}

	useCase := NewRemoveRoutineDetailUseCase(detailRepo)
	err := useCase.Execute(context.Background(), 1, 999)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.RoutineDetailNotFoundError); !ok {
		t.Errorf("Expected RoutineDetailNotFoundError, got %T", err)
	}
}

func TestRemoveRoutineDetailUseCase_ShouldFailWhenDetailBelongsToDifferentRoutine(t *testing.T) {
	detailRepo := &mock.MockRoutineDetailRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.RoutineDetail, error) {
			return &entity.RoutineDetail{ID: id, RoutineID: 99}, nil
		},
	}

	useCase := NewRemoveRoutineDetailUseCase(detailRepo)
	err := useCase.Execute(context.Background(), 1, 10)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.RoutineDetailNotFoundError); !ok {
		t.Errorf("Expected RoutineDetailNotFoundError, got %T", err)
	}
}
