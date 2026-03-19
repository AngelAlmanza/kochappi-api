package templates

import (
	"context"

	"kochappi/internal/application/port"
	domainerror "kochappi/internal/domain/error"
)

type RemoveTemplateDetailUseCase struct {
	templateDetailRepo port.TemplateDetailRepository
}

func NewRemoveTemplateDetailUseCase(templateDetailRepo port.TemplateDetailRepository) *RemoveTemplateDetailUseCase {
	return &RemoveTemplateDetailUseCase{templateDetailRepo: templateDetailRepo}
}

func (uc *RemoveTemplateDetailUseCase) Execute(ctx context.Context, templateID, detailID int) error {
	detail, err := uc.templateDetailRepo.GetByID(ctx, detailID)
	if err != nil {
		return err
	}

	// Ensure the detail belongs to the requested template
	if detail.TemplateID != templateID {
		return &domainerror.TemplateDetailNotFoundError{ID: detailID}
	}

	return uc.templateDetailRepo.DeleteByID(ctx, detailID)
}
