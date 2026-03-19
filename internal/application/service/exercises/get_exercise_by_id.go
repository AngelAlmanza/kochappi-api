package exercises

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetExerciseByIDUseCase struct {
	exerciseRepo port.ExerciseRepository
}

func NewGetExerciseByIDUseCase(exerciseRepo port.ExerciseRepository) *GetExerciseByIDUseCase {
	return &GetExerciseByIDUseCase{exerciseRepo: exerciseRepo}
}

func (uc *GetExerciseByIDUseCase) Execute(ctx context.Context, id int) (*dto.ExerciseResponse, error) {
	exercise, err := uc.exerciseRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.ExerciseResponse{
		ID:       exercise.ID,
		Name:     exercise.Name,
		VideoURL: exercise.VideoURL,
	}, nil
}
