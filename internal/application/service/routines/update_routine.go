package routines

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type UpdateRoutineUseCase struct {
	routineRepo port.RoutineRepository
}

func NewUpdateRoutineUseCase(routineRepo port.RoutineRepository) *UpdateRoutineUseCase {
	return &UpdateRoutineUseCase{routineRepo: routineRepo}
}

func (uc *UpdateRoutineUseCase) Execute(ctx context.Context, id int, req *dto.UpdateRoutineRequest) (*dto.RoutineResponse, error) {
	routine, err := uc.routineRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	routine.Name = req.Name

	if err := uc.routineRepo.Update(ctx, routine); err != nil {
		return nil, err
	}

	return &dto.RoutineResponse{
		ID:         routine.ID,
		CustomerID: routine.CustomerID,
		TemplateID: routine.TemplateID,
		Name:       routine.Name,
		IsActive:   routine.IsActive,
	}, nil
}
