package exercises

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetExercisesUseCase struct {
	exerciseRepo port.ExerciseRepository
}

func NewGetExercisesUseCase(exerciseRepo port.ExerciseRepository) *GetExercisesUseCase {
	return &GetExercisesUseCase{exerciseRepo: exerciseRepo}
}

func (uc *GetExercisesUseCase) Execute(ctx context.Context) ([]dto.ExerciseResponse, error) {
	exercises, err := uc.exerciseRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ExerciseResponse, 0, len(exercises))
	for _, e := range exercises {
		responses = append(responses, dto.ExerciseResponse{
			ID:       e.ID,
			Name:     e.Name,
			VideoURL: e.VideoURL,
		})
	}
	return responses, nil
}
