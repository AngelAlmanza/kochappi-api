package users

import (
	"context"
	"errors"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
	vo "kochappi/internal/domain/value_object"
)

type CreateUserUseCase struct {
	userRepo       port.UserRepository
	passwordHasher port.PasswordHasher
}

func NewCreateUserUseCase(userRepo port.UserRepository, passwordHasher port.PasswordHasher) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo, passwordHasher: passwordHasher}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserDetailResponse, error) {
	email, err := vo.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if _, err := vo.NewPassword(req.Password); err != nil {
		return nil, err
	}

	existing, err := uc.userRepo.GetByEmail(ctx, email.String())
	var notFound *domainerror.UserNotFoundError
	if err != nil && !errors.As(err, &notFound) {
		return nil, err
	}
	if existing != nil {
		return nil, &domainerror.EmailAlreadyExistsError{Email: email.String()}
	}

	hash, err := uc.passwordHasher.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	user := entity.NewUser(req.Name, email.String(), hash, entity.Role(req.Role))
	if err := uc.userRepo.Create(ctx, user); err != nil {
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
