package auth

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestResetPasswordUseCase_ShouldResetPasswordWithValidOTP(t *testing.T) {
	var updatedUser *entity.User
	var deletedUserID string

	otpExpiry := time.Now().Add(10 * time.Minute)
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			user := &entity.User{
				ID:           "user-1",
				Email:        "john@example.com",
				PasswordHash: "old_hash",
				OTPCode:      "123456",
				OTPExpiresAt: &otpExpiry,
			}
			return user, nil
		},
		UpdateFn: func(ctx context.Context, user *entity.User) error {
			updatedUser = user
			return nil
		},
	}

	refreshRepo := &mock.MockRefreshTokenRepository{
		DeleteAllByUserIDFn: func(ctx context.Context, userID string) error {
			deletedUserID = userID
			return nil
		},
	}

	useCase := NewResetPasswordUseCase(userRepo, &mock.MockPasswordHasher{}, refreshRepo)

	resp, err := useCase.Execute(context.Background(), &dto.ResetPasswordRequest{
		Email:       "john@example.com",
		OTPCode:     "123456",
		NewPassword: "newsecurepass",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatal("Expected response, got nil")
	}
	if updatedUser == nil {
		t.Fatal("Expected user to be updated")
	}
	if updatedUser.PasswordHash == "old_hash" {
		t.Error("Expected password hash to be changed")
	}
	if updatedUser.OTPCode != "" {
		t.Error("Expected OTP to be cleared after reset")
	}
	if deletedUserID != "user-1" {
		t.Errorf("Expected all refresh tokens for 'user-1' to be deleted, got '%s'", deletedUserID)
	}
}

func TestResetPasswordUseCase_ShouldFailWithInvalidOTP(t *testing.T) {
	otpExpiry := time.Now().Add(10 * time.Minute)
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return &entity.User{
				ID:           "user-1",
				Email:        "john@example.com",
				OTPCode:      "123456",
				OTPExpiresAt: &otpExpiry,
			}, nil
		},
	}

	useCase := NewResetPasswordUseCase(userRepo, &mock.MockPasswordHasher{}, &mock.MockRefreshTokenRepository{})

	_, err := useCase.Execute(context.Background(), &dto.ResetPasswordRequest{
		Email:       "john@example.com",
		OTPCode:     "000000",
		NewPassword: "newsecurepass",
	})

	if err == nil {
		t.Fatal("Expected error for invalid OTP, got nil")
	}
	if _, ok := err.(*domainerror.InvalidOTPError); !ok {
		t.Errorf("Expected InvalidOTPError, got %T: %v", err, err)
	}
}

func TestResetPasswordUseCase_ShouldFailWithExpiredOTP(t *testing.T) {
	otpExpiry := time.Now().Add(-1 * time.Minute)
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return &entity.User{
				ID:           "user-1",
				Email:        "john@example.com",
				OTPCode:      "123456",
				OTPExpiresAt: &otpExpiry,
			}, nil
		},
	}

	useCase := NewResetPasswordUseCase(userRepo, &mock.MockPasswordHasher{}, &mock.MockRefreshTokenRepository{})

	_, err := useCase.Execute(context.Background(), &dto.ResetPasswordRequest{
		Email:       "john@example.com",
		OTPCode:     "123456",
		NewPassword: "newsecurepass",
	})

	if err == nil {
		t.Fatal("Expected error for expired OTP, got nil")
	}
	if _, ok := err.(*domainerror.InvalidOTPError); !ok {
		t.Errorf("Expected InvalidOTPError, got %T: %v", err, err)
	}
}

func TestResetPasswordUseCase_ShouldFailWithShortNewPassword(t *testing.T) {
	useCase := NewResetPasswordUseCase(
		&mock.MockUserRepository{},
		&mock.MockPasswordHasher{},
		&mock.MockRefreshTokenRepository{},
	)

	_, err := useCase.Execute(context.Background(), &dto.ResetPasswordRequest{
		Email:       "john@example.com",
		OTPCode:     "123456",
		NewPassword: "short",
	})

	if err == nil {
		t.Fatal("Expected error for short password, got nil")
	}
}

func TestResetPasswordUseCase_ShouldFailForNonExistentUser(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return nil, &domainerror.UserNotFoundError{Identifier: email}
		},
	}

	useCase := NewResetPasswordUseCase(userRepo, &mock.MockPasswordHasher{}, &mock.MockRefreshTokenRepository{})

	_, err := useCase.Execute(context.Background(), &dto.ResetPasswordRequest{
		Email:       "nonexistent@example.com",
		OTPCode:     "123456",
		NewPassword: "newsecurepass",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.InvalidOTPError); !ok {
		t.Errorf("Expected InvalidOTPError, got %T: %v", err, err)
	}
}
