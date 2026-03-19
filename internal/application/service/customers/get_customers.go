package customers

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetCustomersUseCase struct {
	customerRepo port.CustomerRepository
}

func NewGetCustomersUseCase(customerRepo port.CustomerRepository) *GetCustomersUseCase {
	return &GetCustomersUseCase{
		customerRepo: customerRepo,
	}
}

func (uc *GetCustomersUseCase) Execute(ctx context.Context) ([]dto.CustomerResponse, error) {
	customers, err := uc.customerRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerResponses []dto.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, dto.CustomerResponse{
			ID:        customer.ID,
			Name:      customer.Name,
			Birthdate: customer.Birthdate.Format(time.DateOnly),
		})
	}

	return customerResponses, nil
}
