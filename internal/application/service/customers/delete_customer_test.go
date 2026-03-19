package customers

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestDeleteCustomerUseCase_ShouldDeleteCustomer(t *testing.T) {
	deleted := false
	repo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id, Name: "Alice", Birthdate: time.Now()}, nil
		},
		DeleteFn: func(ctx context.Context, id int) error {
			deleted = true
			return nil
		},
	}

	useCase := NewDeleteCustomerUseCase(repo)
	err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !deleted {
		t.Error("Expected Delete to be called")
	}
}

func TestDeleteCustomerUseCase_ShouldReturnNotFoundError(t *testing.T) {
	repo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return nil, &domainerror.CustomerNotFoundError{ID: id}
		},
	}

	useCase := NewDeleteCustomerUseCase(repo)
	err := useCase.Execute(context.Background(), 99)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.CustomerNotFoundError); !ok {
		t.Errorf("Expected CustomerNotFoundError, got %T", err)
	}
}

func TestDeleteCustomerUseCase_ShouldNotDeleteUser(t *testing.T) {
	userDeleted := false
	repo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id, UserID: 10, Name: "Alice", Birthdate: time.Now()}, nil
		},
		DeleteFn: func(ctx context.Context, id int) error {
			// Only the customer record is deleted; this mock only receives the customer id
			if id == 10 {
				userDeleted = true
			}
			return nil
		},
	}

	useCase := NewDeleteCustomerUseCase(repo)
	err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if userDeleted {
		t.Error("Delete should only remove customer, not user")
	}
}
