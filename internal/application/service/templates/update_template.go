package templates

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type UpdateTemplateUseCase struct {
	templateRepo port.TemplateRepository
}

func NewUpdateTemplateUseCase(templateRepo port.TemplateRepository) *UpdateTemplateUseCase {
	return &UpdateTemplateUseCase{templateRepo: templateRepo}
}

func (uc *UpdateTemplateUseCase) Execute(ctx context.Context, id int, req *dto.UpdateTemplateRequest) (*dto.TemplateResponse, error) {
	template, err := uc.templateRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	template.Name = req.Name
	template.Description = req.Description

	if err := uc.templateRepo.Update(ctx, template); err != nil {
		return nil, err
	}

	return &dto.TemplateResponse{
		ID:          template.ID,
		Name:        template.Name,
		Description: template.Description,
	}, nil
}
