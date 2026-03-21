package routines

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestAddRoutineDetailUseCase_ShouldAddDetail(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id, CustomerID: 1, Name: "Routine"}, nil
		},
	}
	exerciseRepo := &mock.MockExerciseRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Exercise, error) {
			return &entity.Exercise{ID: id, Name: "Exercise"}, nil
		},
	}
	detailRepo := &mock.MockRoutineDetailRepository{
		CreateFn: func(ctx context.Context, detail *entity.RoutineDetail) error {
			detail.ID = 10
			return nil
		},
	}

	useCase := NewAddRoutineDetailUseCase(routineRepo, detailRepo, exerciseRepo)
	result, err := useCase.Execute(context.Background(), 1, &dto.AddRoutineDetailRequest{
		DayOfWeek:    1,
		ExerciseID:   5,
		Sets:         3,
		Reps:         10,
		DisplayOrder: 1,
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 10 {
		t.Errorf("Expected detail ID 10, got %d", result.ID)
	}
}

func TestAddRoutineDetailUseCase_ShouldFailWhenRoutineNotFound(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return nil, &domainerror.RoutineNotFoundError{ID: id}
		},
	}

	useCase := NewAddRoutineDetailUseCase(routineRepo, &mock.MockRoutineDetailRepository{}, &mock.MockExerciseRepository{})
	_, err := useCase.Execute(context.Background(), 999, &dto.AddRoutineDetailRequest{
		DayOfWeek:    1,
		ExerciseID:   5,
		Sets:         3,
		Reps:         10,
		DisplayOrder: 1,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.RoutineNotFoundError); !ok {
		t.Errorf("Expected RoutineNotFoundError, got %T", err)
	}
}

func TestAddRoutineDetailUseCase_ShouldFailWhenExerciseNotFound(t *testing.T) {
	routineRepo := &mock.MockRoutineRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Routine, error) {
			return &entity.Routine{ID: id, CustomerID: 1, Name: "Routine"}, nil
		},
	}
	exerciseRepo := &mock.MockExerciseRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Exercise, error) {
			return nil, &domainerror.ExerciseNotFoundError{ID: id}
		},
	}

	useCase := NewAddRoutineDetailUseCase(routineRepo, &mock.MockRoutineDetailRepository{}, exerciseRepo)
	_, err := useCase.Execute(context.Background(), 1, &dto.AddRoutineDetailRequest{
		DayOfWeek:    1,
		ExerciseID:   999,
		Sets:         3,
		Reps:         10,
		DisplayOrder: 1,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ExerciseNotFoundError); !ok {
		t.Errorf("Expected ExerciseNotFoundError, got %T", err)
	}
}
