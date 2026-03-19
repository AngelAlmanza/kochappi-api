package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTProvider struct {
	secret           string
	accessExpiryMin  int
	refreshExpiryDay int
}

func NewJWTProvider(secret string, accessExpiryMin, refreshExpiryDay int) *JWTProvider {
	return &JWTProvider{
		secret:           secret,
		accessExpiryMin:  accessExpiryMin,
		refreshExpiryDay: refreshExpiryDay,
	}
}

type AccessClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserID  string `json:"user_id"`
	TokenID string `json:"token_id"`
	jwt.RegisteredClaims
}

func (p *JWTProvider) GenerateAccessToken(userID string, role string) (string, error) {
	claims := AccessClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(p.accessExpiryMin) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "kochappi",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(p.secret))
}

func (p *JWTProvider) GenerateRefreshToken(userID string) (string, string, int64, error) {
	tokenID := uuid.New().String()
	expiresAt := time.Now().Add(time.Duration(p.refreshExpiryDay) * 24 * time.Hour)

	claims := RefreshClaims{
		UserID:  userID,
		TokenID: tokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "kochappi",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(p.secret))
	if err != nil {
		return "", "", 0, err
	}

	return tokenString, tokenID, expiresAt.Unix(), nil
}

func (p *JWTProvider) ValidateAccessToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(p.secret), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*AccessClaims)
	if !ok || !token.Valid {
		return "", "", jwt.ErrSignatureInvalid
	}

	return claims.UserID, claims.Role, nil
}

func (p *JWTProvider) ValidateRefreshToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(p.secret), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return "", "", jwt.ErrSignatureInvalid
	}

	return claims.UserID, claims.TokenID, nil
}
