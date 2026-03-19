package templates

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetTemplatesUseCase struct {
	templateRepo port.TemplateRepository
}

func NewGetTemplatesUseCase(templateRepo port.TemplateRepository) *GetTemplatesUseCase {
	return &GetTemplatesUseCase{templateRepo: templateRepo}
}

func (uc *GetTemplatesUseCase) Execute(ctx context.Context) ([]dto.TemplateResponse, error) {
	tmplates, err := uc.templateRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.TemplateResponse, 0, len(tmplates))
	for _, t := range tmplates {
		responses = append(responses, dto.TemplateResponse{
			ID:          t.ID,
			Name:        t.Name,
			Description: t.Description,
		})
	}
	return responses, nil
}
