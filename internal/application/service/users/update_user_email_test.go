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

func TestUpdateUserEmailUseCase_ShouldUpdateEmail(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return &entity.User{ID: id, Name: "Alice", Email: "old@example.com", Role: entity.ROLE_TRAINER}, nil
		},
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return nil, &domainerror.UserNotFoundError{Identifier: email}
		},
		UpdateFn: func(ctx context.Context, user *entity.User) error {
			return nil
		},
	}

	useCase := NewUpdateUserEmailUseCase(userRepo)
	result, err := useCase.Execute(context.Background(), 1, &dto.UpdateUserEmailRequest{
		Email: "new@example.com",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Email != "new@example.com" {
		t.Errorf("Expected email new@example.com, got %s", result.Email)
	}
}

func TestUpdateUserEmailUseCase_ShouldReturnSameUserWhenEmailUnchanged(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return &entity.User{ID: id, Name: "Alice", Email: "alice@example.com", Role: entity.ROLE_TRAINER}, nil
		},
	}

	useCase := NewUpdateUserEmailUseCase(userRepo)
	result, err := useCase.Execute(context.Background(), 1, &dto.UpdateUserEmailRequest{
		Email: "alice@example.com",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Email != "alice@example.com" {
		t.Errorf("Expected email alice@example.com, got %s", result.Email)
	}
}

func TestUpdateUserEmailUseCase_ShouldReturnConflictWhenEmailTakenByOtherUser(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return &entity.User{ID: id, Name: "Alice", Email: "alice@example.com", Role: entity.ROLE_TRAINER}, nil
		},
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return &entity.User{ID: 2, Email: email}, nil
		},
	}

	useCase := NewUpdateUserEmailUseCase(userRepo)
	_, err := useCase.Execute(context.Background(), 1, &dto.UpdateUserEmailRequest{
		Email: "other@example.com",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	var emailExists *domainerror.EmailAlreadyExistsError
	if !errors.As(err, &emailExists) {
		t.Errorf("Expected EmailAlreadyExistsError, got %T", err)
	}
}

func TestUpdateUserEmailUseCase_ShouldReturnNotFoundError(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return nil, &domainerror.UserNotFoundError{Identifier: "99"}
		},
	}

	useCase := NewUpdateUserEmailUseCase(userRepo)
	_, err := useCase.Execute(context.Background(), 99, &dto.UpdateUserEmailRequest{
		Email: "new@example.com",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	var notFound *domainerror.UserNotFoundError
	if !errors.As(err, &notFound) {
		t.Errorf("Expected UserNotFoundError, got %T", err)
	}
}

func TestUpdateUserEmailUseCase_ShouldReturnErrorOnInvalidEmail(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return &entity.User{ID: id, Email: "alice@example.com"}, nil
		},
	}

	useCase := NewUpdateUserEmailUseCase(userRepo)
	_, err := useCase.Execute(context.Background(), 1, &dto.UpdateUserEmailRequest{
		Email: "not-valid",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
