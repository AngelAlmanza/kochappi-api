package auth

import (
	"context"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestRegisterUseCase_ShouldRegisterNewTrainer(t *testing.T) {
	var createdUser *entity.User

	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return nil, &domainerror.UserNotFoundError{Identifier: email}
		},
		CreateFn: func(ctx context.Context, user *entity.User) error {
			createdUser = user
			return nil
		},
	}

	useCase := NewRegisterUseCase(
		userRepo,
		&mock.MockPasswordHasher{},
		&mock.MockTokenProvider{},
		&mock.MockRefreshTokenRepository{},
	)

	resp, err := useCase.Execute(context.Background(), &dto.RegisterRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "securepass123",
		Role:     "trainer",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatal("Expected response, got nil")
	}
	if resp.User.Email != "john@example.com" {
		t.Errorf("Expected email 'john@example.com', got '%s'", resp.User.Email)
	}
	if resp.User.Role != "trainer" {
		t.Errorf("Expected role 'trainer', got '%s'", resp.User.Role)
	}
	if resp.AccessToken == "" {
		t.Error("Expected access token to be set")
	}
	if resp.RefreshToken == "" {
		t.Error("Expected refresh token to be set")
	}
	if createdUser == nil {
		t.Fatal("Expected user to be persisted via Create")
	}
}

func TestRegisterUseCase_ShouldFailWhenEmailAlreadyExists(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return &entity.User{
				ID:    "existing-id",
				Email: email,
			}, nil
		},
	}

	useCase := NewRegisterUseCase(
		userRepo,
		&mock.MockPasswordHasher{},
		&mock.MockTokenProvider{},
		&mock.MockRefreshTokenRepository{},
	)

	_, err := useCase.Execute(context.Background(), &dto.RegisterRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "securepass123",
		Role:     "trainer",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.EmailAlreadyExistsError); !ok {
		t.Errorf("Expected EmailAlreadyExistsError, got %T: %v", err, err)
	}
}

func TestRegisterUseCase_ShouldFailWithInvalidEmail(t *testing.T) {
	useCase := NewRegisterUseCase(
		&mock.MockUserRepository{},
		&mock.MockPasswordHasher{},
		&mock.MockTokenProvider{},
		&mock.MockRefreshTokenRepository{},
	)

	_, err := useCase.Execute(context.Background(), &dto.RegisterRequest{
		Name:     "John Doe",
		Email:    "not-an-email",
		Password: "securepass123",
		Role:     "trainer",
	})

	if err == nil {
		t.Fatal("Expected error for invalid email, got nil")
	}
}

func TestRegisterUseCase_ShouldFailWithShortPassword(t *testing.T) {
	useCase := NewRegisterUseCase(
		&mock.MockUserRepository{},
		&mock.MockPasswordHasher{},
		&mock.MockTokenProvider{},
		&mock.MockRefreshTokenRepository{},
	)

	_, err := useCase.Execute(context.Background(), &dto.RegisterRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "short",
		Role:     "trainer",
	})

	if err == nil {
		t.Fatal("Expected error for short password, got nil")
	}
}
