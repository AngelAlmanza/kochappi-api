package mock

import "time"

type MockTokenProvider struct {
	GenerateAccessTokenFn  func(userID string, role string) (string, error)
	GenerateRefreshTokenFn func(userID string) (string, string, int64, error)
	ValidateAccessTokenFn  func(tokenString string) (string, string, error)
	ValidateRefreshTokenFn func(tokenString string) (string, string, error)
}

func (p *MockTokenProvider) GenerateAccessToken(userID string, role string) (string, error) {
	if p.GenerateAccessTokenFn != nil {
		return p.GenerateAccessTokenFn(userID, role)
	}
	return "access_token_" + userID, nil
}

func (p *MockTokenProvider) GenerateRefreshToken(userID string) (string, string, int64, error) {
	if p.GenerateRefreshTokenFn != nil {
		return p.GenerateRefreshTokenFn(userID)
	}
	return "refresh_token_" + userID, "token_id_123", time.Now().Add(7 * 24 * time.Hour).Unix(), nil
}

func (p *MockTokenProvider) ValidateAccessToken(tokenString string) (string, string, error) {
	if p.ValidateAccessTokenFn != nil {
		return p.ValidateAccessTokenFn(tokenString)
	}
	return "user-1", "trainer", nil
}

func (p *MockTokenProvider) ValidateRefreshToken(tokenString string) (string, string, error) {
	if p.ValidateRefreshTokenFn != nil {
		return p.ValidateRefreshTokenFn(tokenString)
	}
	return "user-1", "token_id_123", nil
}
