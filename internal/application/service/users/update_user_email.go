package users

import (
	"context"
	"errors"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	domainerror "kochappi/internal/domain/error"
	vo "kochappi/internal/domain/value_object"
)

type UpdateUserEmailUseCase struct {
	userRepo port.UserRepository
}

func NewUpdateUserEmailUseCase(userRepo port.UserRepository) *UpdateUserEmailUseCase {
	return &UpdateUserEmailUseCase{userRepo: userRepo}
}

func (uc *UpdateUserEmailUseCase) Execute(ctx context.Context, id int, req *dto.UpdateUserEmailRequest) (*dto.UserDetailResponse, error) {
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	email, err := vo.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user.Email == email.String() {
		return &dto.UserDetailResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      string(user.Role),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}, nil
	}

	existing, err := uc.userRepo.GetByEmail(ctx, email.String())
	var notFound *domainerror.UserNotFoundError
	if err != nil && !errors.As(err, &notFound) {
		return nil, err
	}
	if existing != nil {
		return nil, &domainerror.EmailAlreadyExistsError{Email: email.String()}
	}

	user.Email = email.String()
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return &dto.UserDetailResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
