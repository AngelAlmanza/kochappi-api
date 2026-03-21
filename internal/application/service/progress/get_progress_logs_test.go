package progress

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestGetProgressLogsUseCase_ShouldReturnLogs(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	logRepo := &mock.MockLogCustomerProgressRepository{
		GetByCustomerIDFn: func(ctx context.Context, customerID int) ([]entity.LogCustomerProgress, error) {
			return []entity.LogCustomerProgress{
				{ID: 1, CheckDate: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC), Weight: 75, CustomerID: customerID},
				{ID: 2, CheckDate: time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC), Weight: 74, CustomerID: customerID},
			}, nil
		},
	}

	uc := NewGetProgressLogsUseCase(customerRepo, logRepo)
	result, err := uc.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("Expected 2 logs, got %d", len(result))
	}
	if result[0].Weight != 75 {
		t.Errorf("Expected weight 75, got %d", result[0].Weight)
	}
}

func TestGetProgressLogsUseCase_ShouldFailWhenCustomerNotFound(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return nil, &domainerror.CustomerNotFoundError{ID: id}
		},
	}

	uc := NewGetProgressLogsUseCase(customerRepo, &mock.MockLogCustomerProgressRepository{})
	_, err := uc.Execute(context.Background(), 99)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.CustomerNotFoundError); !ok {
		t.Errorf("Expected CustomerNotFoundError, got %T", err)
	}
}
