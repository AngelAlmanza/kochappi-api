package routines

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestCreateRoutineUseCase_ShouldCreateRoutineWithoutDetails(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	routineRepo := &mock.MockRoutineRepository{
		CreateFn: func(ctx context.Context, routine *entity.Routine) error {
			routine.ID = 1
			return nil
		},
	}

	useCase := NewCreateRoutineUseCase(routineRepo, &mock.MockRoutineDetailRepository{}, &mock.MockRoutinePeriodRepository{}, customerRepo, &mock.MockTemplateRepository{}, &mock.MockExerciseRepository{})
	result, err := useCase.Execute(context.Background(), &dto.CreateRoutineRequest{
		CustomerID: 1,
		Name:       "My Routine",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 || result.Name != "My Routine" {
		t.Errorf("Unexpected result: %+v", result)
	}
	if len(result.Details) != 0 {
		t.Errorf("Expected 0 details, got %d", len(result.Details))
	}
}

func TestCreateRoutineUseCase_ShouldCreateRoutineWithDetails(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	routineRepo := &mock.MockRoutineRepository{
		CreateFn: func(ctx context.Context, routine *entity.Routine) error {
			routine.ID = 1
			return nil
		},
	}
	detailRepo := &mock.MockRoutineDetailRepository{
		CreateBulkFn: func(ctx context.Context, details []*entity.RoutineDetail) error {
			for i, d := range details {
				d.ID = i + 10
			}
			return nil
		},
	}
	exerciseRepo := &mock.MockExerciseRepository{
		GetByIDsFn: func(ctx context.Context, ids []int) ([]entity.Exercise, error) {
			exercises := make([]entity.Exercise, 0, len(ids))
			for _, id := range ids {
				exercises = append(exercises, entity.Exercise{ID: id, Name: "Exercise"})
			}
			return exercises, nil
		},
	}

	useCase := NewCreateRoutineUseCase(routineRepo, detailRepo, &mock.MockRoutinePeriodRepository{}, customerRepo, &mock.MockTemplateRepository{}, exerciseRepo)
	result, err := useCase.Execute(context.Background(), &dto.CreateRoutineRequest{
		CustomerID: 1,
		Name:       "My Routine",
		Details: []dto.CreateRoutineDetailRequest{
			{DayOfWeek: 1, ExerciseID: 5, Sets: 3, Reps: 10, DisplayOrder: 1},
			{DayOfWeek: 2, ExerciseID: 6, Sets: 4, Reps: 8, DisplayOrder: 2},
		},
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Details) != 2 {
		t.Errorf("Expected 2 details, got %d", len(result.Details))
	}
}

func TestCreateRoutineUseCase_ShouldFailWhenCustomerNotFound(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return nil, &domainerror.CustomerNotFoundError{ID: id}
		},
	}

	useCase := NewCreateRoutineUseCase(&mock.MockRoutineRepository{}, &mock.MockRoutineDetailRepository{}, &mock.MockRoutinePeriodRepository{}, customerRepo, &mock.MockTemplateRepository{}, &mock.MockExerciseRepository{})
	_, err := useCase.Execute(context.Background(), &dto.CreateRoutineRequest{
		CustomerID: 999,
		Name:       "My Routine",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.CustomerNotFoundError); !ok {
		t.Errorf("Expected CustomerNotFoundError, got %T", err)
	}
}

func TestCreateRoutineUseCase_ShouldFailWhenActiveRoutineExists(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	routineRepo := &mock.MockRoutineRepository{
		GetActiveByCustomerIDFn: func(ctx context.Context, customerID int) (*entity.Routine, error) {
			return &entity.Routine{ID: 99, CustomerID: customerID, IsActive: true}, nil
		},
	}

	useCase := NewCreateRoutineUseCase(routineRepo, &mock.MockRoutineDetailRepository{}, &mock.MockRoutinePeriodRepository{}, customerRepo, &mock.MockTemplateRepository{}, &mock.MockExerciseRepository{})
	_, err := useCase.Execute(context.Background(), &dto.CreateRoutineRequest{
		CustomerID: 1,
		Name:       "My Routine",
		IsActive:   true,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ActiveRoutineExistsError); !ok {
		t.Errorf("Expected ActiveRoutineExistsError, got %T", err)
	}
}

func TestCreateRoutineUseCase_ShouldCreatePeriodWhenActive(t *testing.T) {
	periodCreated := false
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	routineRepo := &mock.MockRoutineRepository{
		GetActiveByCustomerIDFn: func(ctx context.Context, customerID int) (*entity.Routine, error) {
			return nil, nil
		},
		CreateFn: func(ctx context.Context, routine *entity.Routine) error {
			routine.ID = 1
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

	useCase := NewCreateRoutineUseCase(routineRepo, &mock.MockRoutineDetailRepository{}, periodRepo, customerRepo, &mock.MockTemplateRepository{}, &mock.MockExerciseRepository{})
	result, err := useCase.Execute(context.Background(), &dto.CreateRoutineRequest{
		CustomerID: 1,
		Name:       "My Routine",
		IsActive:   true,
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !periodCreated {
		t.Error("Expected period to be created")
	}
	if !result.IsActive {
		t.Error("Expected routine to be active")
	}
}

func TestCreateRoutineUseCase_ShouldValidateTemplateWhenProvided(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	templateRepo := &mock.MockTemplateRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Template, error) {
			return nil, &domainerror.TemplateNotFoundError{ID: id}
		},
	}

	templateID := 999
	useCase := NewCreateRoutineUseCase(&mock.MockRoutineRepository{}, &mock.MockRoutineDetailRepository{}, &mock.MockRoutinePeriodRepository{}, customerRepo, templateRepo, &mock.MockExerciseRepository{})
	_, err := useCase.Execute(context.Background(), &dto.CreateRoutineRequest{
		CustomerID: 1,
		TemplateID: &templateID,
		Name:       "My Routine",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.TemplateNotFoundError); !ok {
		t.Errorf("Expected TemplateNotFoundError, got %T", err)
	}
}
