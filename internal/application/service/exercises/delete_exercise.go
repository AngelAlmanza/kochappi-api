package exercises

import (
	"context"

	"kochappi/internal/application/port"
)

type DeleteExerciseUseCase struct {
	exerciseRepo port.ExerciseRepository
}

func NewDeleteExerciseUseCase(exerciseRepo port.ExerciseRepository) *DeleteExerciseUseCase {
	return &DeleteExerciseUseCase{exerciseRepo: exerciseRepo}
}

func (uc *DeleteExerciseUseCase) Execute(ctx context.Context, id int) error {
	_, err := uc.exerciseRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return uc.exerciseRepo.Delete(ctx, id)
}
