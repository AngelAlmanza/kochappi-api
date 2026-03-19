package customers

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestCreateCustomerUseCase_ShouldCreateCustomer(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return &entity.User{ID: id, Role: entity.ROLE_CLIENT}, nil
		},
	}
	customerRepo := &mock.MockCustomerRepository{
		GetByUserIDFn: func(ctx context.Context, userID int) (*entity.Customer, error) {
			return nil, &domainerror.UserNotCustomerError{ID: userID}
		},
		CreateFn: func(ctx context.Context, customer *entity.Customer) error {
			customer.ID = 1
			return nil
		},
	}

	useCase := NewCreateCustomerUseCase(customerRepo, userRepo)
	result, err := useCase.Execute(context.Background(), &dto.CreateCustomerRequest{
		Name:      "Alice",
		Birthdate: "1990-01-01",
		UserID:    5,
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 || result.Name != "Alice" {
		t.Errorf("Unexpected result: %+v", result)
	}
	if result.Birthdate != "1990-01-01" {
		t.Errorf("Expected birthdate 1990-01-01, got %s", result.Birthdate)
	}
}

func TestCreateCustomerUseCase_ShouldFailWhenUserNotFound(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return nil, &domainerror.UserNotFoundError{Identifier: "99"}
		},
	}

	useCase := NewCreateCustomerUseCase(&mock.MockCustomerRepository{}, userRepo)
	_, err := useCase.Execute(context.Background(), &dto.CreateCustomerRequest{
		Name:      "Alice",
		Birthdate: "1990-01-01",
		UserID:    99,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.UserNotFoundError); !ok {
		t.Errorf("Expected UserNotFoundError, got %T", err)
	}
}

func TestCreateCustomerUseCase_ShouldFailWhenUserIsNotClient(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return &entity.User{ID: id, Role: entity.ROLE_TRAINER}, nil
		},
	}

	useCase := NewCreateCustomerUseCase(&mock.MockCustomerRepository{}, userRepo)
	_, err := useCase.Execute(context.Background(), &dto.CreateCustomerRequest{
		Name:      "Alice",
		Birthdate: "1990-01-01",
		UserID:    1,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.UserNotCustomerError); !ok {
		t.Errorf("Expected UserNotCustomerError, got %T", err)
	}
}

func TestCreateCustomerUseCase_ShouldFailWhenCustomerAlreadyExists(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return &entity.User{ID: id, Role: entity.ROLE_CLIENT}, nil
		},
	}
	customerRepo := &mock.MockCustomerRepository{
		GetByUserIDFn: func(ctx context.Context, userID int) (*entity.Customer, error) {
			return &entity.Customer{ID: 1, UserID: userID}, nil
		},
	}

	useCase := NewCreateCustomerUseCase(customerRepo, userRepo)
	_, err := useCase.Execute(context.Background(), &dto.CreateCustomerRequest{
		Name:      "Alice",
		Birthdate: "1990-01-01",
		UserID:    5,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.CustomerAlreadyExistsError); !ok {
		t.Errorf("Expected CustomerAlreadyExistsError, got %T", err)
	}
}

func TestCreateCustomerUseCase_ShouldFailWithInvalidBirthdate(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return &entity.User{ID: id, Role: entity.ROLE_CLIENT}, nil
		},
	}
	customerRepo := &mock.MockCustomerRepository{
		GetByUserIDFn: func(ctx context.Context, userID int) (*entity.Customer, error) {
			return nil, &domainerror.UserNotCustomerError{ID: userID}
		},
	}

	useCase := NewCreateCustomerUseCase(customerRepo, userRepo)
	_, err := useCase.Execute(context.Background(), &dto.CreateCustomerRequest{
		Name:      "Alice",
		Birthdate: "not-a-date",
		UserID:    5,
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.InvalidBirthdateError); !ok {
		t.Errorf("Expected InvalidBirthdateError, got %T", err)
	}
}
