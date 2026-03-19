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

func TestRefreshTokenUseCase_ShouldRotateTokens(t *testing.T) {
	var deletedTokenID string
	var storedTokenID string

	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return &entity.User{
				ID:   1,
				Role: entity.ROLE_TRAINER,
			}, nil
		},
	}

	refreshRepo := &mock.MockRefreshTokenRepository{
		ExistsFn: func(ctx context.Context, tokenID string) (bool, error) {
			return true, nil
		},
		DeleteByIDFn: func(ctx context.Context, tokenID string) error {
			deletedTokenID = tokenID
			return nil
		},
		StoreFn: func(ctx context.Context, userID int, tokenID string, expiresAt int64) error {
			storedTokenID = tokenID
			return nil
		},
	}

	useCase := NewRefreshTokenUseCase(userRepo, &mock.MockTokenProvider{}, refreshRepo)

	resp, err := useCase.Execute(context.Background(), &dto.RefreshTokenRequest{
		RefreshToken: "valid_refresh_token",
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatal("Expected response, got nil")
	}
	if resp.AccessToken == "" {
		t.Error("Expected new access token")
	}
	if resp.RefreshToken == "" {
		t.Error("Expected new refresh token")
	}
	if deletedTokenID != "token_id_123" {
		t.Errorf("Expected old token ID 'token_id_123' to be deleted, got '%s'", deletedTokenID)
	}
	if storedTokenID == "" {
		t.Error("Expected new refresh token to be stored")
	}
}

func TestRefreshTokenUseCase_ShouldFailWhenUserNotFound(t *testing.T) {
	userRepo := &mock.MockUserRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.User, error) {
			return nil, &domainerror.UserNotFoundError{Identifier: "1"}
		},
	}

	refreshRepo := &mock.MockRefreshTokenRepository{
		ExistsFn: func(ctx context.Context, tokenID string) (bool, error) {
			return true, nil
		},
		DeleteByIDFn: func(ctx context.Context, tokenID string) error {
			return nil
		},
	}

	useCase := NewRefreshTokenUseCase(userRepo, &mock.MockTokenProvider{}, refreshRepo)

	_, err := useCase.Execute(context.Background(), &dto.RefreshTokenRequest{
		RefreshToken: "valid_refresh_token",
	})

	if err == nil {
		t.Fatal("Expected error when user not found, got nil")
	}
	if _, ok := err.(*domainerror.InvalidTokenError); !ok {
		t.Errorf("Expected InvalidTokenError, got %T: %v", err, err)
	}
}

func TestRefreshTokenUseCase_ShouldFailWithInvalidToken(t *testing.T) {
	tokenProvider := &mock.MockTokenProvider{
		ValidateRefreshTokenFn: func(tokenString string) (int, string, error) {
			return 0, "", errors.New("invalid token")
		},
	}

	useCase := NewRefreshTokenUseCase(
		&mock.MockUserRepository{},
		tokenProvider,
		&mock.MockRefreshTokenRepository{},
	)

	_, err := useCase.Execute(context.Background(), &dto.RefreshTokenRequest{
		RefreshToken: "invalid_token",
	})

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.InvalidTokenError); !ok {
		t.Errorf("Expected InvalidTokenError, got %T: %v", err, err)
	}
}

func TestRefreshTokenUseCase_ShouldFailWithRevokedToken(t *testing.T) {
	refreshRepo := &mock.MockRefreshTokenRepository{
		ExistsFn: func(ctx context.Context, tokenID string) (bool, error) {
			return false, nil
		},
	}

	useCase := NewRefreshTokenUseCase(
		&mock.MockUserRepository{},
		&mock.MockTokenProvider{},
		refreshRepo,
	)

	_, err := useCase.Execute(context.Background(), &dto.RefreshTokenRequest{
		RefreshToken: "revoked_token",
	})

	if err == nil {
		t.Fatal("Expected error for revoked token, got nil")
	}
	if _, ok := err.(*domainerror.InvalidTokenError); !ok {
		t.Errorf("Expected InvalidTokenError, got %T: %v", err, err)
	}
}
