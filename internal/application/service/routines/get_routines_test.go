package routines

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
)

func TestGetRoutinesUseCase_ShouldReturnAllRoutines(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetAllFn: func(ctx context.Context) ([]entity.Routine, error) {
			return []entity.Routine{
				{ID: 1, CustomerID: 1, Name: "Routine 1"},
				{ID: 2, CustomerID: 2, Name: "Routine 2"},
			}, nil
		},
	}

	useCase := NewGetRoutinesUseCase(routineRepo)
	result, err := useCase.Execute(context.Background(), nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 routines, got %d", len(result))
	}
}

func TestGetRoutinesUseCase_ShouldFilterByCustomerID(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByCustomerIDFn: func(ctx context.Context, customerID int) ([]entity.Routine, error) {
			return []entity.Routine{
				{ID: 1, CustomerID: customerID, Name: "Routine 1"},
			}, nil
		},
	}

	customerID := 1
	useCase := NewGetRoutinesUseCase(routineRepo)
	result, err := useCase.Execute(context.Background(), &customerID)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 1 routine, got %d", len(result))
	}
	if result[0].CustomerID != 1 {
		t.Errorf("Expected customerID 1, got %d", result[0].CustomerID)
	}
}
