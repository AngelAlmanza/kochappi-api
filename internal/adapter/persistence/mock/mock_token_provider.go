package mock

import (
	"fmt"
	"time"
)

type MockTokenProvider struct {
	GenerateAccessTokenFn  func(userID int, role string) (string, error)
	GenerateRefreshTokenFn func(userID int) (string, string, int64, error)
	ValidateAccessTokenFn  func(tokenString string) (int, string, error)
	ValidateRefreshTokenFn func(tokenString string) (int, string, error)
}

func (p *MockTokenProvider) GenerateAccessToken(userID int, role string) (string, error) {
	if p.GenerateAccessTokenFn != nil {
		return p.GenerateAccessTokenFn(userID, role)
	}
	return fmt.Sprintf("access_token_%d", userID), nil
}

func (p *MockTokenProvider) GenerateRefreshToken(userID int) (string, string, int64, error) {
	if p.GenerateRefreshTokenFn != nil {
		return p.GenerateRefreshTokenFn(userID)
	}
	return fmt.Sprintf("refresh_token_%d", userID), "token_id_123", time.Now().Add(7 * 24 * time.Hour).Unix(), nil
}

func (p *MockTokenProvider) ValidateAccessToken(tokenString string) (int, string, error) {
	if p.ValidateAccessTokenFn != nil {
		return p.ValidateAccessTokenFn(tokenString)
	}
	return 1, "trainer", nil
}

func (p *MockTokenProvider) ValidateRefreshToken(tokenString string) (int, string, error) {
	if p.ValidateRefreshTokenFn != nil {
		return p.ValidateRefreshTokenFn(tokenString)
	}
	return 1, "token_id_123", nil
}
