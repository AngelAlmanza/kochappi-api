package customers

import (
	"context"

	"kochappi/internal/application/port"
)

type DeleteCustomerUseCase struct {
	customerRepo port.CustomerRepository
}

func NewDeleteCustomerUseCase(customerRepo port.CustomerRepository) *DeleteCustomerUseCase {
	return &DeleteCustomerUseCase{customerRepo: customerRepo}
}

func (uc *DeleteCustomerUseCase) Execute(ctx context.Context, id int) error {
	_, err := uc.customerRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return uc.customerRepo.Delete(ctx, id)
}
