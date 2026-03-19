package port

import (
	"context"
	"kochappi/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id int) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
}

type RefreshTokenRepository interface {
	Store(ctx context.Context, userID int, tokenID string, expiresAt int64) error
	Exists(ctx context.Context, tokenID string) (bool, error)
	DeleteByID(ctx context.Context, tokenID string) error
	DeleteAllByUserID(ctx context.Context, userID int) error
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type TokenProvider interface {
	GenerateAccessToken(userID int, role string) (string, error)
	GenerateRefreshToken(userID int) (tokenString string, tokenID string, expiresAt int64, err error)
	ValidateAccessToken(tokenString string) (userID int, role string, err error)
	ValidateRefreshToken(tokenString string) (userID int, tokenID string, err error)
}

type OTPService interface {
	GenerateCode() string
	Send(ctx context.Context, email string, code string) error
}
