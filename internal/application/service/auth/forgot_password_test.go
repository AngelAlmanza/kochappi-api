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

func TestForgotPasswordUseCase_ShouldSendOTPForExistingUser(t *testing.T) {
	var updatedUser *entity.User
	var sentEmail, sentCode string

	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return &entity.User{
				ID:    "user-1",
				Email: "john@example.com",
			}, nil
		},
		UpdateFn: func(ctx context.Context, user *entity.User) error {
			updatedUser = user
			return nil
		},
	}

	otpService := &mock.MockOTPService{
		GenerateCodeFn: func() string {
			return "654321"
		},
		SendFn: func(ctx context.Context, email string, code string) error {
			sentEmail = email
			sentCode = code
			return nil
		},
	}

	useCase := NewForgotPasswordUseCase(userRepo, otpService, 10)

	resp, err := useCase.Execute(context.Background(), &dto.ForgotPasswordRequest{
		Email: "john@example.com",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatal("Expected response, got nil")
	}
	if updatedUser == nil {
		t.Fatal("Expected user to be updated with OTP")
	}
	if updatedUser.OTPCode != "654321" {
		t.Errorf("Expected OTP code '654321', got '%s'", updatedUser.OTPCode)
	}
	if updatedUser.OTPExpiresAt == nil {
		t.Error("Expected OTPExpiresAt to be set")
	}
	if sentEmail != "john@example.com" {
		t.Errorf("Expected OTP sent to 'john@example.com', got '%s'", sentEmail)
	}
	if sentCode != "654321" {
		t.Errorf("Expected OTP code '654321' sent, got '%s'", sentCode)
	}
}

func TestForgotPasswordUseCase_ShouldFailWhenSendFails(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return &entity.User{ID: "user-1", Email: email}, nil
		},
		UpdateFn: func(ctx context.Context, user *entity.User) error {
			return nil
		},
	}

	otpService := &mock.MockOTPService{
		SendFn: func(ctx context.Context, email string, code string) error {
			return errors.New("email service unavailable")
		},
	}

	useCase := NewForgotPasswordUseCase(userRepo, otpService, 10)

	_, err := useCase.Execute(context.Background(), &dto.ForgotPasswordRequest{
		Email: "john@example.com",
	})

	if err == nil {
		t.Fatal("Expected error when OTP send fails, got nil")
	}
}

func TestForgotPasswordUseCase_ShouldReturnSuccessForNonExistentEmail(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return nil, &domainerror.UserNotFoundError{Identifier: email}
		},
	}

	useCase := NewForgotPasswordUseCase(userRepo, &mock.MockOTPService{}, 10)

	resp, err := useCase.Execute(context.Background(), &dto.ForgotPasswordRequest{
		Email: "nonexistent@example.com",
	})

	if err != nil {
		t.Fatalf("Expected no error (to prevent email enumeration), got %v", err)
	}
	if resp == nil {
		t.Fatal("Expected response, got nil")
	}
}
