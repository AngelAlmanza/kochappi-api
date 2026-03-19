package templates

import (
	"context"

	"kochappi/internal/application/port"
)

// DeleteTemplateUseCase deletes a template.
// The database handles cascading:
//   - template_detail rows are deleted (ON DELETE CASCADE)
//   - routines.template_id is set to NULL (ON DELETE SET NULL)
type DeleteTemplateUseCase struct {
	templateRepo port.TemplateRepository
}

func NewDeleteTemplateUseCase(templateRepo port.TemplateRepository) *DeleteTemplateUseCase {
	return &DeleteTemplateUseCase{templateRepo: templateRepo}
}

func (uc *DeleteTemplateUseCase) Execute(ctx context.Context, id int) error {
	if _, err := uc.templateRepo.GetByID(ctx, id); err != nil {
		return err
	}
	return uc.templateRepo.Delete(ctx, id)
}
