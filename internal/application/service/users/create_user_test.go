package users

import (
	"context"
	"errors"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestCreateUserUseCase_ShouldCreateUser(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return nil, &domainerror.UserNotFoundError{Identifier: email}
		},
		CreateFn: func(ctx context.Context, user *entity.User) error {
			user.ID = 1
			return nil
		},
	}
	hasher := &mock.MockPasswordHasher{
		HashFn: func(password string) (string, error) {
			return "hashed_" + password, nil
		},
	}

	useCase := NewCreateUserUseCase(userRepo, hasher)
	result, err := useCase.Execute(context.Background(), &dto.CreateUserRequest{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "password123",
		Role:     "trainer",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 {
		t.Errorf("Expected ID 1, got %d", result.ID)
	}
	if result.Role != "trainer" {
		t.Errorf("Expected role trainer, got %s", result.Role)
	}
}

func TestCreateUserUseCase_ShouldReturnErrorWhenEmailAlreadyExists(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return &entity.User{ID: 1, Email: email}, nil
		},
	}
	hasher := &mock.MockPasswordHasher{}

	useCase := NewCreateUserUseCase(userRepo, hasher)
	_, err := useCase.Execute(context.Background(), &dto.CreateUserRequest{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "password123",
		Role:     "trainer",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	var emailExists *domainerror.EmailAlreadyExistsError
	if !errors.As(err, &emailExists) {
		t.Errorf("Expected EmailAlreadyExistsError, got %T", err)
	}
}

func TestCreateUserUseCase_ShouldReturnErrorOnInvalidEmail(t *testing.T) {
	userRepo := &mock.MockUserRepository{}
	hasher := &mock.MockPasswordHasher{}

	useCase := NewCreateUserUseCase(userRepo, hasher)
	_, err := useCase.Execute(context.Background(), &dto.CreateUserRequest{
		Name:     "Alice",
		Email:    "not-an-email",
		Password: "password123",
		Role:     "trainer",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestCreateUserUseCase_ShouldReturnErrorOnShortPassword(t *testing.T) {
	userRepo := &mock.MockUserRepository{}
	hasher := &mock.MockPasswordHasher{}

	useCase := NewCreateUserUseCase(userRepo, hasher)
	_, err := useCase.Execute(context.Background(), &dto.CreateUserRequest{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "short",
		Role:     "trainer",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
