package routines

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestActivateRoutineUseCase_ShouldActivateInactiveRoutine(t *testing.T) {
	periodCreated := false
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id, CustomerID: 1, Name: "Routine", IsActive: false}, nil
		},
		GetActiveByCustomerIDFn: func(ctx context.Context, customerID int) (*entity.Routine, error) {
			return nil, nil
		},
		UpdateFn: func(ctx context.Context, routine *entity.Routine) error {
			return nil
		},
	}
	periodRepo := &mock.MockRoutinePeriodRepository{
		CreateFn: func(ctx context.Context, period *entity.RoutinePeriod) error {
			periodCreated = true
			period.ID = 1
			return nil
		},
	}

	useCase := NewActivateRoutineUseCase(routineRepo, periodRepo)
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !result.IsActive {
		t.Error("Expected routine to be active")
	}
	if !periodCreated {
		t.Error("Expected period to be created")
	}
}

func TestActivateRoutineUseCase_ShouldReturnEarlyWhenAlreadyActive(t *testing.T) {
	updateCalled := false
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id, CustomerID: 1, Name: "Routine", IsActive: true}, nil
		},
		UpdateFn: func(ctx context.Context, routine *entity.Routine) error {
			updateCalled = true
			return nil
		},
	}

	useCase := NewActivateRoutineUseCase(routineRepo, &mock.MockRoutinePeriodRepository{})
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !result.IsActive {
		t.Error("Expected routine to be active")
	}
	if updateCalled {
		t.Error("Expected Update not to be called")
	}
}

func TestActivateRoutineUseCase_ShouldFailWhenAnotherActiveExists(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id, CustomerID: 1, Name: "Routine", IsActive: false}, nil
		},
		GetActiveByCustomerIDFn: func(ctx context.Context, customerID int) (*entity.Routine, error) {
			return &entity.Routine{ID: 99, CustomerID: customerID, IsActive: true}, nil
		},
	}

	useCase := NewActivateRoutineUseCase(routineRepo, &mock.MockRoutinePeriodRepository{})
	_, err := useCase.Execute(context.Background(), 1)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ActiveRoutineExistsError); !ok {
		t.Errorf("Expected ActiveRoutineExistsError, got %T", err)
	}
}

func TestActivateRoutineUseCase_ShouldFailWhenNotFound(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return nil, &domainerror.RoutineNotFoundError{ID: id}
		},
	}

	useCase := NewActivateRoutineUseCase(routineRepo, &mock.MockRoutinePeriodRepository{})
	_, err := useCase.Execute(context.Background(), 999)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.RoutineNotFoundError); !ok {
		t.Errorf("Expected RoutineNotFoundError, got %T", err)
	}
}
