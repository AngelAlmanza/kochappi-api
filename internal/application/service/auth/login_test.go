package auth

import (
	"context"
	"errors"
	"testing"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestLoginUseCase_ShouldLoginWithValidCredentials(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return &entity.User{
				ID:           1,
				Name:         "John Doe",
				Email:        "john@example.com",
				PasswordHash: "hashed_securepass123",
				Role:         entity.ROLE_TRAINER,
			}, nil
		},
	}

	useCase := NewLoginUseCase(
		userRepo,
		&mock.MockPasswordHasher{},
		&mock.MockTokenProvider{},
		&mock.MockRefreshTokenRepository{},
	)

	resp, err := useCase.Execute(context.Background(), &dto.LoginRequest{
		Email:    "john@example.com",
		Password: "securepass123",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatal("Expected response, got nil")
	}
	if resp.User.ID != 1 {
		t.Errorf("Expected user ID 1, got %d", resp.User.ID)
	}
	if resp.AccessToken == "" {
		t.Error("Expected access token to be set")
	}
	if resp.RefreshToken == "" {
		t.Error("Expected refresh token to be set")
	}
}

func TestLoginUseCase_ShouldFailWithWrongPassword(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return &entity.User{
				ID:           1,
				Email:        "john@example.com",
				PasswordHash: "hashed_correctpassword",
				Role:         entity.ROLE_TRAINER,
			}, nil
		},
	}

	passwordHasher := &mock.MockPasswordHasher{
		CompareFn: func(hashedPassword, password string) error {
			return errors.New("password mismatch")
		},
	}

	useCase := NewLoginUseCase(
		userRepo,
		passwordHasher,
		&mock.MockTokenProvider{},
		&mock.MockRefreshTokenRepository{},
	)

	_, err := useCase.Execute(context.Background(), &dto.LoginRequest{
		Email:    "john@example.com",
		Password: "wrongpassword",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.InvalidCredentialsError); !ok {
		t.Errorf("Expected InvalidCredentialsError, got %T: %v", err, err)
	}
}

func TestLoginUseCase_ShouldFailWithNonExistentEmail(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return nil, &domainerror.UserNotFoundError{Identifier: email}
		},
	}

	useCase := NewLoginUseCase(
		userRepo,
		&mock.MockPasswordHasher{},
		&mock.MockTokenProvider{},
		&mock.MockRefreshTokenRepository{},
	)

	_, err := useCase.Execute(context.Background(), &dto.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "securepass123",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.InvalidCredentialsError); !ok {
		t.Errorf("Expected InvalidCredentialsError, got %T: %v", err, err)
	}
}

func TestLoginUseCase_ShouldFailWithInvalidEmail(t *testing.T) {
	useCase := NewLoginUseCase(
		&mock.MockUserRepository{},
		&mock.MockPasswordHasher{},
		&mock.MockTokenProvider{},
		&mock.MockRefreshTokenRepository{},
	)

	_, err := useCase.Execute(context.Background(), &dto.LoginRequest{
		Email:    "invalid-email",
		Password: "securepass123",
	})

	if err == nil {
		t.Fatal("Expected error for invalid email, got nil")
	}
	if _, ok := err.(*domainerror.InvalidCredentialsError); !ok {
		t.Errorf("Expected InvalidCredentialsError, got %T: %v", err, err)
	}
}
