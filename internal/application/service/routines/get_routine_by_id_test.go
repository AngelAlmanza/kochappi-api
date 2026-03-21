package routines

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestGetRoutineByIDUseCase_ShouldReturnRoutineWithDetails(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id, CustomerID: 1, Name: "My Routine"}, nil
		},
	}
	detailRepo := &mock.MockRoutineDetailRepository{
		GetByRoutineIDFn: func(ctx context.Context, routineID int) ([]entity.RoutineDetail, error) {
			return []entity.RoutineDetail{
				{ID: 10, RoutineID: routineID, DayOfWeek: 1, ExerciseID: 5, Sets: 3, Reps: 10, DisplayOrder: 1},
			}, nil
		},
	}

	useCase := NewGetRoutineByIDUseCase(routineRepo, detailRepo)
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 || result.Name != "My Routine" {
		t.Errorf("Unexpected result: %+v", result)
	}
	if len(result.Details) != 1 {
		t.Errorf("Expected 1 detail, got %d", len(result.Details))
	}
}

func TestGetRoutineByIDUseCase_ShouldFailWhenNotFound(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return nil, &domainerror.RoutineNotFoundError{ID: id}
		},
	}

	useCase := NewGetRoutineByIDUseCase(routineRepo, &mock.MockRoutineDetailRepository{})
	_, err := useCase.Execute(context.Background(), 999)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.RoutineNotFoundError); !ok {
		t.Errorf("Expected RoutineNotFoundError, got %T", err)
	}
}
