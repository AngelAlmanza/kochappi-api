package users

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetUserByIDUseCase struct {
	userRepo port.UserRepository
}

func NewGetUserByIDUseCase(userRepo port.UserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{userRepo: userRepo}
}

func (uc *GetUserByIDUseCase) Execute(ctx context.Context, id int) (*dto.UserDetailResponse, error) {
	u, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.UserDetailResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      string(u.Role),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
