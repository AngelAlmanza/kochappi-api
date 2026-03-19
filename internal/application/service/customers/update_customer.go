package customers

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	domainerror "kochappi/internal/domain/error"
)

type UpdateCustomerUseCase struct {
	customerRepo port.CustomerRepository
}

func NewUpdateCustomerUseCase(customerRepo port.CustomerRepository) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{customerRepo: customerRepo}
}

func (uc *UpdateCustomerUseCase) Execute(ctx context.Context, id int, req *dto.UpdateCustomerRequest) (*dto.CustomerResponse, error) {
	customer, err := uc.customerRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	birthdate, err := time.Parse(time.DateOnly, req.Birthdate)
	if err != nil {
		return nil, &domainerror.InvalidBirthdateError{Birthdate: req.Birthdate}
	}

	customer.Name = req.Name
	customer.Birthdate = birthdate

	if err := uc.customerRepo.Update(ctx, customer); err != nil {
		return nil, err
	}

	return &dto.CustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Birthdate: customer.Birthdate.Format(time.DateOnly),
	}, nil
}
