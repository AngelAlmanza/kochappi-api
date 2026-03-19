package auth

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	domainerror "kochappi/internal/domain/error"
	"kochappi/internal/domain/value_object"
)

type LoginUseCase struct {
	userRepo       port.UserRepository
	passwordHasher port.PasswordHasher
	tokenProvider  port.TokenProvider
	refreshRepo    port.RefreshTokenRepository
}

func NewLoginUseCase(
	userRepo port.UserRepository,
	passwordHasher port.PasswordHasher,
	tokenProvider port.TokenProvider,
	refreshRepo port.RefreshTokenRepository,
) *LoginUseCase {
	return &LoginUseCase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		tokenProvider:  tokenProvider,
		refreshRepo:    refreshRepo,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error) {
	email, err := value_object.NewEmail(req.Email)
	if err != nil {
		return nil, &domainerror.InvalidCredentialsError{}
	}

	user, err := uc.userRepo.GetByEmail(ctx, email.String())
	if err != nil {
		return nil, &domainerror.InvalidCredentialsError{}
	}

	if err := uc.passwordHasher.Compare(user.PasswordHash, req.Password); err != nil {
		return nil, &domainerror.InvalidCredentialsError{}
	}

	accessToken, err := uc.tokenProvider.GenerateAccessToken(user.ID, string(user.Role))
	if err != nil {
		return nil, err
	}

	refreshToken, tokenID, expiresAt, err := uc.tokenProvider.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err := uc.refreshRepo.Store(ctx, user.ID, tokenID, expiresAt); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		User: dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  string(user.Role),
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
