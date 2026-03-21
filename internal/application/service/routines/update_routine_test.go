package routines

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestUpdateRoutineUseCase_ShouldUpdateName(t *testing.T) {
	routine := &entity.Routine{
		ID:         1,
		CustomerID: 1,
		Name:       "Old Name",
	}
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return routine, nil
		},
		UpdateFn: func(ctx context.Context, r *entity.Routine) error {
			r.Name = "New Name"
			return nil
		},
	}

	useCase := NewUpdateRoutineUseCase(routineRepo)
	result, err := useCase.Execute(context.Background(), 1, &dto.UpdateRoutineRequest{Name: "New Name"})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Name != "New Name" {
		t.Errorf("Expected name 'New Name', got '%s'", result.Name)
	}
}

func TestUpdateRoutineUseCase_ShouldFailWhenNotFound(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return nil, &domainerror.RoutineNotFoundError{ID: id}
		},
	}

	useCase := NewUpdateRoutineUseCase(routineRepo)
	_, err := useCase.Execute(context.Background(), 999, &dto.UpdateRoutineRequest{Name: "New Name"})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.RoutineNotFoundError); !ok {
		t.Errorf("Expected RoutineNotFoundError, got %T", err)
	}
}
