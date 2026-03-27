package users

import (
	"context"
	"errors"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
)

func TestGetUsersUseCase_ShouldReturnAllUsers(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetAllFn: func(ctx context.Context, role *entity.Role) ([]entity.User, error) {
			return []entity.User{
				{ID: 1, Name: "Alice", Email: "alice@example.com", Role: entity.ROLE_TRAINER},
				{ID: 2, Name: "Bob", Email: "bob@example.com", Role: entity.ROLE_CLIENT},
			}, nil
		},
	}

	useCase := NewGetUsersUseCase(userRepo)
	result, err := useCase.Execute(context.Background(), nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 users, got %d", len(result))
	}
}

func TestGetUsersUseCase_ShouldFilterByRole(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetAllFn: func(ctx context.Context, role *entity.Role) ([]entity.User, error) {
			if role == nil || *role != entity.ROLE_TRAINER {
				t.Error("Expected trainer role filter")
			}
			return []entity.User{
				{ID: 1, Name: "Alice", Email: "alice@example.com", Role: entity.ROLE_TRAINER},
			}, nil
		},
	}

	role := entity.ROLE_TRAINER
	useCase := NewGetUsersUseCase(userRepo)
	result, err := useCase.Execute(context.Background(), &role)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 1 user, got %d", len(result))
	}
	if result[0].Role != string(entity.ROLE_TRAINER) {
		t.Errorf("Expected role trainer, got %s", result[0].Role)
	}
}

func TestGetUsersUseCase_ShouldReturnErrorOnRepositoryFailure(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetAllFn: func(ctx context.Context, role *entity.Role) ([]entity.User, error) {
			return nil, errors.New("database error")
		},
	}

	useCase := NewGetUsersUseCase(userRepo)
	_, err := useCase.Execute(context.Background(), nil)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
