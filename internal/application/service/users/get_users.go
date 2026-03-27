package users

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
)

type GetUsersUseCase struct {
	userRepo port.UserRepository
}

func NewGetUsersUseCase(userRepo port.UserRepository) *GetUsersUseCase {
	return &GetUsersUseCase{userRepo: userRepo}
}

func (uc *GetUsersUseCase) Execute(ctx context.Context, role *entity.Role) ([]dto.UserDetailResponse, error) {
	userList, err := uc.userRepo.GetAll(ctx, role)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.UserDetailResponse, 0, len(userList))
	for _, u := range userList {
		responses = append(responses, dto.UserDetailResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			Role:      string(u.Role),
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return responses, nil
}
