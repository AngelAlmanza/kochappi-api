package templates

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
)

type CreateTemplateUseCase struct {
	templateRepo       port.TemplateRepository
	templateDetailRepo port.TemplateDetailRepository
	exerciseRepo       port.ExerciseRepository
}

func NewCreateTemplateUseCase(
	templateRepo port.TemplateRepository,
	templateDetailRepo port.TemplateDetailRepository,
	exerciseRepo port.ExerciseRepository,
) *CreateTemplateUseCase {
	return &CreateTemplateUseCase{
		templateRepo:       templateRepo,
		templateDetailRepo: templateDetailRepo,
		exerciseRepo:       exerciseRepo,
	}
}

func (uc *CreateTemplateUseCase) Execute(ctx context.Context, req *dto.CreateTemplateRequest) (*dto.TemplateWithDetailsResponse, error) {
	// Validate all exercises exist in a single query (deduplicated IDs)
	if len(req.Details) > 0 {
		seen := make(map[int]struct{}, len(req.Details))
		ids := make([]int, 0, len(req.Details))
		for _, d := range req.Details {
			if _, ok := seen[d.ExerciseID]; !ok {
				// struct{}{} is an empty struct, it doesn't take up any memory
				// It is used as a value in the map to indicate the presence of a key
				seen[d.ExerciseID] = struct{}{}
				ids = append(ids, d.ExerciseID)
			}
		}
		if _, err := uc.exerciseRepo.GetByIDs(ctx, ids); err != nil {
			return nil, err
		}
	}

	template := entity.NewTemplate(req.Name, req.Description)
	if err := uc.templateRepo.Create(ctx, template); err != nil {
		return nil, err
	}

	detailEntities := make([]*entity.TemplateDetail, 0, len(req.Details))
	for _, d := range req.Details {
		detailEntities = append(detailEntities, entity.NewTemplateDetail(template.ID, d.DayOfWeek, d.ExerciseID, d.Sets, d.Reps, d.DisplayOrder))
	}

	if len(detailEntities) > 0 {
		if err := uc.templateDetailRepo.CreateBulk(ctx, detailEntities); err != nil {
			return nil, err
		}
	}

	detailResponses := make([]dto.TemplateDetailResponse, 0, len(detailEntities))
	for _, d := range detailEntities {
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
