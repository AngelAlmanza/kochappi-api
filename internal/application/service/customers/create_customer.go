package customers

import (
	"context"
	"strconv"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

type CreateCustomerUseCase struct {
	customerRepo port.CustomerRepository
	userRepo     port.UserRepository
}

func NewCreateCustomerUseCase(customerRepo port.CustomerRepository, userRepo port.UserRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		customerRepo: customerRepo,
		userRepo:     userRepo,
	}
}

func (uc *CreateCustomerUseCase) Execute(ctx context.Context, req *dto.CreateCustomerRequest) (*dto.CustomerResponse, error) {
	// Validate user exists
	user, err := uc.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return nil, &domainerror.UserNotFoundError{Identifier: strconv.Itoa(req.UserID)}
	}

	// Validate user role is client
	if user.Role != entity.ROLE_CLIENT {
		return nil, &domainerror.UserNotCustomerError{ID: req.UserID}
	}

	// Validate no customer already linked to this user
	_, err = uc.customerRepo.GetByUserID(ctx, req.UserID)
	if err == nil {
		return nil, &domainerror.CustomerAlreadyExistsError{UserID: req.UserID}
	}
	if _, ok := err.(*domainerror.UserNotCustomerError); !ok {
		return nil, err
	}

	// Validate birthdate
	birthdate, err := time.Parse(time.DateOnly, req.Birthdate)
	if err != nil {
		return nil, &domainerror.InvalidBirthdateError{Birthdate: req.Birthdate}
	}

	customerEntity := entity.NewCustomer(req.Name, birthdate, req.UserID)
	if err := uc.customerRepo.Create(ctx, customerEntity); err != nil {
		return nil, err
	}

	return &dto.CustomerResponse{
		ID:        customerEntity.ID,
		Name:      customerEntity.Name,
		Birthdate: customerEntity.Birthdate.Format(time.DateOnly),
	}, nil
}
