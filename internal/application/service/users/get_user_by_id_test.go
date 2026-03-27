package users

import (
	"context"
	"errors"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestGetUserByIDUseCase_ShouldReturnUser(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return &entity.User{ID: id, Name: "Alice", Email: "alice@example.com", Role: entity.ROLE_TRAINER}, nil
		},
	}

	useCase := NewGetUserByIDUseCase(userRepo)
	result, err := useCase.Execute(context.Background(), 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 {
		t.Errorf("Expected ID 1, got %d", result.ID)
	}
	if result.Email != "alice@example.com" {
		t.Errorf("Expected email alice@example.com, got %s", result.Email)
	}
}

func TestGetUserByIDUseCase_ShouldReturnNotFoundError(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return nil, &domainerror.UserNotFoundError{Identifier: "99"}
		},
	}

	useCase := NewGetUserByIDUseCase(userRepo)
	_, err := useCase.Execute(context.Background(), 99)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	var notFound *domainerror.UserNotFoundError
	if !errors.As(err, &notFound) {
		t.Errorf("Expected UserNotFoundError, got %T", err)
	}
}
