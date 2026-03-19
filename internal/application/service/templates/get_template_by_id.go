package templates

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetTemplateByIDUseCase struct {
	templateRepo       port.TemplateRepository
	templateDetailRepo port.TemplateDetailRepository
}

func NewGetTemplateByIDUseCase(
	templateRepo port.TemplateRepository,
	templateDetailRepo port.TemplateDetailRepository,
) *GetTemplateByIDUseCase {
	return &GetTemplateByIDUseCase{
		templateRepo:       templateRepo,
		templateDetailRepo: templateDetailRepo,
	}
}

func (uc *GetTemplateByIDUseCase) Execute(ctx context.Context, id int) (*dto.TemplateWithDetailsResponse, error) {
	template, err := uc.templateRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	details, err := uc.templateDetailRepo.GetByTemplateID(ctx, id)
	if err != nil {
		return nil, err
	}

	detailResponses := make([]dto.TemplateDetailResponse, 0, len(details))
	for _, d := range details {
		detailResponses = append(detailResponses, dto.TemplateDetailResponse{
			ID:           d.ID,
			DayOfWeek:    d.DayOfWeek,
			ExerciseID:   d.ExerciseID,
			Sets:         d.Sets,
			Reps:         d.Reps,
			DisplayOrder: d.DisplayOrder,
		})
	}

	return &dto.TemplateWithDetailsResponse{
		ID:          template.ID,
		Name:        template.Name,
		Description: template.Description,
		Details:     detailResponses,
	}, nil
}
