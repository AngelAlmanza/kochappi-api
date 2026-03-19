package customers

import (
	"context"
	"errors"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
)

func TestGetCustomersUseCase_ShouldReturnAllCustomers(t *testing.T) {
	repo := &mock.MockCustomerRepository{
		GetAllFn: func(ctx context.Context) ([]entity.Customer, error) {
			return []entity.Customer{
				{ID: 1, Name: "Alice", Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)},
				{ID: 2, Name: "Bob", Birthdate: time.Date(1985, 6, 15, 0, 0, 0, 0, time.UTC)},
			}, nil
		},
	}

	useCase := NewGetCustomersUseCase(repo)
	result, err := useCase.Execute(context.Background())

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 customers, got %d", len(result))
	}
	if result[0].Name != "Alice" {
		t.Errorf("Expected first customer Alice, got %s", result[0].Name)
	}
}

func TestGetCustomersUseCase_ShouldReturnEmptyList(t *testing.T) {
	repo := &mock.MockCustomerRepository{
		GetAllFn: func(ctx context.Context) ([]entity.Customer, error) {
			return []entity.Customer{}, nil
		},
	}

	useCase := NewGetCustomersUseCase(repo)
	result, err := useCase.Execute(context.Background())

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected 0 customers, got %d", len(result))
	}
}

func TestGetCustomersUseCase_ShouldPropagateRepositoryError(t *testing.T) {
	repo := &mock.MockCustomerRepository{
		GetAllFn: func(ctx context.Context) ([]entity.Customer, error) {
			return nil, errors.New("db error")
		},
	}

	useCase := NewGetCustomersUseCase(repo)
	_, err := useCase.Execute(context.Background())

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
