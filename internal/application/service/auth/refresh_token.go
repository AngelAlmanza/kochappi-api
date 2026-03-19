package auth

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	domainerror "kochappi/internal/domain/error"
)

type RefreshTokenUseCase struct {
	userRepo      port.UserRepository
	tokenProvider port.TokenProvider
	refreshRepo   port.RefreshTokenRepository
}

func NewRefreshTokenUseCase(
	userRepo port.UserRepository,
	tokenProvider port.TokenProvider,
	refreshRepo port.RefreshTokenRepository,
) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		userRepo:      userRepo,
		tokenProvider: tokenProvider,
		refreshRepo:   refreshRepo,
	}
}

func (uc *RefreshTokenUseCase) Execute(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.TokenResponse, error) {
	userID, tokenID, err := uc.tokenProvider.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, &domainerror.InvalidTokenError{}
	}

	exists, err := uc.refreshRepo.Exists(ctx, tokenID)
	if err != nil || !exists {
		return nil, &domainerror.InvalidTokenError{}
	}

	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, &domainerror.InvalidTokenError{}
	}

	if err := uc.refreshRepo.DeleteByID(ctx, tokenID); err != nil {
		return nil, err
	}

	accessToken, err := uc.tokenProvider.GenerateAccessToken(user.ID, string(user.Role))
	if err != nil {
		return nil, err
	}

	newRefreshToken, newTokenID, expiresAt, err := uc.tokenProvider.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err := uc.refreshRepo.Store(ctx, user.ID, newTokenID, expiresAt); err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
