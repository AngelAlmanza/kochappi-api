package sessions

import (
	"context"

	"kochappi/internal/application/port"
)

type DeleteExerciseLogUseCase struct {
	logExerciseSessionRepo port.LogExerciseSessionRepository
}

func NewDeleteExerciseLogUseCase(logExerciseSessionRepo port.LogExerciseSessionRepository) *DeleteExerciseLogUseCase {
	return &DeleteExerciseLogUseCase{logExerciseSessionRepo: logExerciseSessionRepo}
}

func (uc *DeleteExerciseLogUseCase) Execute(ctx context.Context, id int) error {
	if _, err := uc.logExerciseSessionRepo.GetByID(ctx, id); err != nil {
		return err
	}

	return uc.logExerciseSessionRepo.Delete(ctx, id)
}
