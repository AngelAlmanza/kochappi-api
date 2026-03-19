package auth

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	domainerror "kochappi/internal/domain/error"
	"kochappi/internal/domain/value_object"
)

type ResetPasswordUseCase struct {
	userRepo       port.UserRepository
	passwordHasher port.PasswordHasher
	refreshRepo    port.RefreshTokenRepository
}

func NewResetPasswordUseCase(
	userRepo port.UserRepository,
	passwordHasher port.PasswordHasher,
	refreshRepo port.RefreshTokenRepository,
) *ResetPasswordUseCase {
	return &ResetPasswordUseCase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		refreshRepo:    refreshRepo,
	}
}

func (uc *ResetPasswordUseCase) Execute(ctx context.Context, req *dto.ResetPasswordRequest) (*dto.MessageResponse, error) {
	if _, err := value_object.NewPassword(req.NewPassword); err != nil {
		return nil, err
	}

	user, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, &domainerror.InvalidOTPError{}
	}

	if !user.IsOTPValid(req.OTPCode) {
		return nil, &domainerror.InvalidOTPError{}
	}

	hashedPassword, err := uc.passwordHasher.Hash(req.NewPassword)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = hashedPassword
	user.ClearOTP()

	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	if err := uc.refreshRepo.DeleteAllByUserID(ctx, user.ID); err != nil {
		return nil, err
	}

	return &dto.MessageResponse{
		Message: "Password has been reset successfully",
	}, nil
}
