package routines

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestGetRoutinePeriodsUseCase_ShouldReturnPeriods(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id, CustomerID: 1, Name: "Routine"}, nil
		},
	}
	endedAt := time.Now()
	periodRepo := &mock.MockRoutinePeriodRepository{
		GetByRoutineIDFn: func(ctx context.Context, routineID int) ([]entity.RoutinePeriod, error) {
			return []entity.RoutinePeriod{
				{ID: 1, RoutineID: routineID, StartedAt: time.Now(), EndedAt: &endedAt},
				{ID: 2, RoutineID: routineID, StartedAt: time.Now()},
			}, nil
		},
	}

	useCase := NewGetRoutinePeriodsUseCase(routineRepo, periodRepo)
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 periods, got %d", len(result))
	}
}

func TestGetRoutinePeriodsUseCase_ShouldFailWhenRoutineNotFound(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return nil, &domainerror.RoutineNotFoundError{ID: id}
		},
	}

	useCase := NewGetRoutinePeriodsUseCase(routineRepo, &mock.MockRoutinePeriodRepository{})
	_, err := useCase.Execute(context.Background(), 999)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.RoutineNotFoundError); !ok {
		t.Errorf("Expected RoutineNotFoundError, got %T", err)
	}
}
