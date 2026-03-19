package templates

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
)

type AddTemplateDetailUseCase struct {
	templateRepo       port.TemplateRepository
	templateDetailRepo port.TemplateDetailRepository
	exerciseRepo       port.ExerciseRepository
}

func NewAddTemplateDetailUseCase(
	templateRepo port.TemplateRepository,
	templateDetailRepo port.TemplateDetailRepository,
	exerciseRepo port.ExerciseRepository,
) *AddTemplateDetailUseCase {
	return &AddTemplateDetailUseCase{
		templateRepo:       templateRepo,
		templateDetailRepo: templateDetailRepo,
		exerciseRepo:       exerciseRepo,
	}
}

func (uc *AddTemplateDetailUseCase) Execute(ctx context.Context, templateID int, req *dto.AddTemplateDetailRequest) (*dto.TemplateDetailResponse, error) {
	if _, err := uc.templateRepo.GetByID(ctx, templateID); err != nil {
		return nil, err
	}

	if _, err := uc.exerciseRepo.GetByID(ctx, req.ExerciseID); err != nil {
		return nil, err
	}

	detail := entity.NewTemplateDetail(templateID, req.DayOfWeek, req.ExerciseID, req.Sets, req.Reps, req.DisplayOrder)
	if err := uc.templateDetailRepo.Create(ctx, detail); err != nil {
		return nil, err
	}

	return &dto.TemplateDetailResponse{
		ID:           detail.ID,
		DayOfWeek:    detail.DayOfWeek,
		ExerciseID:   detail.ExerciseID,
		Sets:         detail.Sets,
		Reps:         detail.Reps,
		DisplayOrder: detail.DisplayOrder,
	}, nil
}
