package routines

import (
	"context"

	"kochappi/internal/application/port"
	domainerror "kochappi/internal/domain/error"
)

type RemoveRoutineDetailUseCase struct {
	routineDetailRepo port.RoutineDetailRepository
}

func NewRemoveRoutineDetailUseCase(routineDetailRepo port.RoutineDetailRepository) *RemoveRoutineDetailUseCase {
	return &RemoveRoutineDetailUseCase{routineDetailRepo: routineDetailRepo}
}

func (uc *RemoveRoutineDetailUseCase) Execute(ctx context.Context, routineID, detailID int) error {
	detail, err := uc.routineDetailRepo.GetByID(ctx, detailID)
	if err != nil {
		return err
	}

	if detail.RoutineID != routineID {
		return &domainerror.RoutineDetailNotFoundError{ID: detailID}
	}

	return uc.routineDetailRepo.DeleteByID(ctx, detailID)
}
