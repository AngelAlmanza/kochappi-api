package auth

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
	"kochappi/internal/domain/value_object"
)

type RegisterUseCase struct {
	userRepo       port.UserRepository
	passwordHasher port.PasswordHasher
	tokenProvider  port.TokenProvider
	refreshRepo    port.RefreshTokenRepository
}

func NewRegisterUseCase(
	userRepo port.UserRepository,
	passwordHasher port.PasswordHasher,
	tokenProvider port.TokenProvider,
	refreshRepo port.RefreshTokenRepository,
) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		tokenProvider:  tokenProvider,
		refreshRepo:    refreshRepo,
	}
}

func (uc *RegisterUseCase) Execute(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	email, err := value_object.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if _, err := value_object.NewPassword(req.Password); err != nil {
		return nil, err
	}

	existing, _ := uc.userRepo.GetByEmail(ctx, email.String())
	if existing != nil {
		return nil, &domainerror.EmailAlreadyExistsError{Email: email.String()}
	}

	hashedPassword, err := uc.passwordHasher.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	user := entity.NewUser(
		req.Name,
		email.String(),
		hashedPassword,
		entity.Role(req.Role),
	)

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
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
