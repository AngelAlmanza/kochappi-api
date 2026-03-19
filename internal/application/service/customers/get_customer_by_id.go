package customers

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetCustomerByIDUseCase struct {
	customerRepo port.CustomerRepository
}

func NewGetCustomerByIDUseCase(customerRepo port.CustomerRepository) *GetCustomerByIDUseCase {
	return &GetCustomerByIDUseCase{
		customerRepo: customerRepo,
	}
}

func (uc *GetCustomerByIDUseCase) Execute(ctx context.Context, id int) (*dto.CustomerResponse, error) {
	customer, err := uc.customerRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return &dto.CustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Birthdate: customer.Birthdate.Format(time.DateOnly),
	}, nil
}
