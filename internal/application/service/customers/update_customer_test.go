package customers

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestUpdateCustomerUseCase_ShouldUpdateCustomer(t *testing.T) {
	repo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id, Name: "Old Name", Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)}, nil
		},
		UpdateFn: func(ctx context.Context, customer *entity.Customer) error {
			return nil
		},
	}

	useCase := NewUpdateCustomerUseCase(repo)
	result, err := useCase.Execute(context.Background(), 1, &dto.UpdateCustomerRequest{
		Name:      "New Name",
		Birthdate: "1995-06-15",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Name != "New Name" {
		t.Errorf("Expected name 'New Name', got %s", result.Name)
	}
	if result.Birthdate != "1995-06-15" {
		t.Errorf("Expected birthdate 1995-06-15, got %s", result.Birthdate)
	}
}

func TestUpdateCustomerUseCase_ShouldReturnNotFoundError(t *testing.T) {
	repo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return nil, &domainerror.CustomerNotFoundError{ID: id}
		},
	}

	useCase := NewUpdateCustomerUseCase(repo)
	_, err := useCase.Execute(context.Background(), 99, &dto.UpdateCustomerRequest{
		Name:      "X",
		Birthdate: "1990-01-01",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.CustomerNotFoundError); !ok {
		t.Errorf("Expected CustomerNotFoundError, got %T", err)
	}
}

func TestUpdateCustomerUseCase_ShouldFailWithInvalidBirthdate(t *testing.T) {
	repo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id, Name: "Alice"}, nil
		},
	}

	useCase := NewUpdateCustomerUseCase(repo)
	_, err := useCase.Execute(context.Background(), 1, &dto.UpdateCustomerRequest{
		Name:      "Alice",
		Birthdate: "bad-date",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.InvalidBirthdateError); !ok {
		t.Errorf("Expected InvalidBirthdateError, got %T", err)
	}
}
