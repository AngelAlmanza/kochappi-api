package routines

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestDeactivateRoutineUseCase_ShouldDeactivateActiveRoutine(t *testing.T) {
	periodUpdated := false
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id, CustomerID: 1, Name: "Routine", IsActive: true}, nil
		},
		UpdateFn: func(ctx context.Context, routine *entity.Routine) error {
			return nil
		},
	}
	periodRepo := &mock.MockRoutinePeriodRepository{
		GetOngoingByRoutineIDFn: func(ctx context.Context, routineID int) (*entity.RoutinePeriod, error) {
			return &entity.RoutinePeriod{ID: 1, RoutineID: routineID, StartedAt: time.Now()}, nil
		},
		UpdateFn: func(ctx context.Context, period *entity.RoutinePeriod) error {
			periodUpdated = true
			return nil
		},
	}

	useCase := NewDeactivateRoutineUseCase(routineRepo, periodRepo)
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.IsActive {
		t.Error("Expected routine to be inactive")
	}
	if !periodUpdated {
		t.Error("Expected period to be updated")
	}
}

func TestDeactivateRoutineUseCase_ShouldReturnEarlyWhenAlreadyInactive(t *testing.T) {
	updateCalled := false
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id, CustomerID: 1, Name: "Routine", IsActive: false}, nil
		},
		UpdateFn: func(ctx context.Context, routine *entity.Routine) error {
			updateCalled = true
			return nil
		},
	}

	useCase := NewDeactivateRoutineUseCase(routineRepo, &mock.MockRoutinePeriodRepository{})
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.IsActive {
		t.Error("Expected routine to be inactive")
	}
	if updateCalled {
		t.Error("Expected Update not to be called")
	}
}

func TestDeactivateRoutineUseCase_ShouldHandleNoOngoingPeriod(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id, CustomerID: 1, Name: "Routine", IsActive: true}, nil
		},
		UpdateFn: func(ctx context.Context, routine *entity.Routine) error {
			return nil
		},
	}
	periodRepo := &mock.MockRoutinePeriodRepository{
		GetOngoingByRoutineIDFn: func(ctx context.Context, routineID int) (*entity.RoutinePeriod, error) {
			return nil, nil
		},
	}

	useCase := NewDeactivateRoutineUseCase(routineRepo, periodRepo)
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.IsActive {
		t.Error("Expected routine to be inactive")
	}
}

func TestDeactivateRoutineUseCase_ShouldFailWhenNotFound(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return nil, &domainerror.RoutineNotFoundError{ID: id}
		},
	}

	useCase := NewDeactivateRoutineUseCase(routineRepo, &mock.MockRoutinePeriodRepository{})
	_, err := useCase.Execute(context.Background(), 999)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.RoutineNotFoundError); !ok {
		t.Errorf("Expected RoutineNotFoundError, got %T", err)
	}
}
