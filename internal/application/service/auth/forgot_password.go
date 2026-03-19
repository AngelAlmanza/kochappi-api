package auth

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type ForgotPasswordUseCase struct {
	userRepo       port.UserRepository
	otpService     port.OTPService
	otpExpiryMins  int
}

func NewForgotPasswordUseCase(
	userRepo port.UserRepository,
	otpService port.OTPService,
	otpExpiryMins int,
) *ForgotPasswordUseCase {
	return &ForgotPasswordUseCase{
		userRepo:      userRepo,
		otpService:    otpService,
		otpExpiryMins: otpExpiryMins,
	}
}

func (uc *ForgotPasswordUseCase) Execute(ctx context.Context, req *dto.ForgotPasswordRequest) (*dto.MessageResponse, error) {
	user, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return &dto.MessageResponse{
			Message: "If the email exists, a recovery code has been sent",
		}, nil
	}

	code := uc.otpService.GenerateCode()
	expiresAt := time.Now().Add(time.Duration(uc.otpExpiryMins) * time.Minute)
	user.SetOTP(code, expiresAt)

	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	if err := uc.otpService.Send(ctx, user.Email, code); err != nil {
		return nil, err
	}

	return &dto.MessageResponse{
		Message: "If the email exists, a recovery code has been sent",
	}, nil
}
