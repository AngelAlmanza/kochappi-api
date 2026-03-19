package customers

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestGetCustomerByIDUseCase_ShouldReturnCustomer(t *testing.T) {
	repo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id, Name: "Alice", Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)}, nil
		},
	}

	useCase := NewGetCustomerByIDUseCase(repo)
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 || result.Name != "Alice" {
		t.Errorf("Unexpected result: %+v", result)
	}
}

func TestGetCustomerByIDUseCase_ShouldReturnNotFoundError(t *testing.T) {
	repo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return nil, &domainerror.CustomerNotFoundError{ID: id}
		},
	}

	useCase := NewGetCustomerByIDUseCase(repo)
	_, err := useCase.Execute(context.Background(), 99)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.CustomerNotFoundError); !ok {
		t.Errorf("Expected CustomerNotFoundError, got %T", err)
	}
}
