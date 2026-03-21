package progress

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestCreateProgressLogUseCase_ShouldCreateLog(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	logRepo := &mock.MockLogCustomerProgressRepository{
		CreateFn: func(ctx context.Context, log *entity.LogCustomerProgress) error {
			log.ID = 1
			return nil
		},
	}

	uc := NewCreateProgressLogUseCase(customerRepo, logRepo)
	result, err := uc.Execute(context.Background(), 1, &dto.CreateProgressLogRequest{
		CheckDate: "2024-01-15",
		Weight:    75,
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 {
		t.Errorf("Expected ID 1, got %d", result.ID)
	}
	if result.CheckDate != "2024-01-15" {
		t.Errorf("Expected checkDate 2024-01-15, got %s", result.CheckDate)
	}
	if result.Weight != 75 {
		t.Errorf("Expected weight 75, got %d", result.Weight)
	}
}

func TestCreateProgressLogUseCase_ShouldFailWithInvalidDate(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}

	uc := NewCreateProgressLogUseCase(customerRepo, &mock.MockLogCustomerProgressRepository{})
	_, err := uc.Execute(context.Background(), 1, &dto.CreateProgressLogRequest{
		CheckDate: "not-a-date",
		Weight:    75,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.InvalidCheckDateError); !ok {
		t.Errorf("Expected InvalidCheckDateError, got %T", err)
	}
}

func TestCreateProgressLogUseCase_ShouldFailWithInvalidWeight(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}

	uc := NewCreateProgressLogUseCase(customerRepo, &mock.MockLogCustomerProgressRepository{})
	_, err := uc.Execute(context.Background(), 1, &dto.CreateProgressLogRequest{
		CheckDate: "2024-01-15",
		Weight:    0,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.InvalidWeightError); !ok {
		t.Errorf("Expected InvalidWeightError, got %T", err)
	}
}

func TestCreateProgressLogUseCase_ShouldFailWhenCustomerNotFound(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return nil, &domainerror.CustomerNotFoundError{ID: id}
		},
	}

	uc := NewCreateProgressLogUseCase(customerRepo, &mock.MockLogCustomerProgressRepository{})
	_, err := uc.Execute(context.Background(), 99, &dto.CreateProgressLogRequest{
		CheckDate: "2024-01-15",
		Weight:    75,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.CustomerNotFoundError); !ok {
		t.Errorf("Expected CustomerNotFoundError, got %T", err)
	}
}
