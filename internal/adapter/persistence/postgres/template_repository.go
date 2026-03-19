package postgres

import (
	"context"
	"errors"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"

	"gorm.io/gorm"
)

type PostgresTemplateRepository struct {
	db *gorm.DB
}

func NewPostgresTemplateRepository(db *gorm.DB) *PostgresTemplateRepository {
	return &PostgresTemplateRepository{db: db}
}

func (r *PostgresTemplateRepository) GetAll(ctx context.Context) ([]entity.Template, error) {
	var models []model.TemplateModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	templates := make([]entity.Template, 0, len(models))
	for _, m := range models {
		templates = append(templates, *m.ToDomain())
	}
	return templates, nil
}

func (r *PostgresTemplateRepository) GetByID(ctx context.Context, id int) (*entity.Template, error) {
	var m model.TemplateModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.TemplateNotFoundError{ID: id}
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *PostgresTemplateRepository) Create(ctx context.Context, template *entity.Template) error {
	m := model.TemplateModelFromDomain(template)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	template.ID = m.ID
	return nil
}

func (r *PostgresTemplateRepository) Update(ctx context.Context, template *entity.Template) error {
	m := model.TemplateModelFromDomain(template)
	return r.db.WithContext(ctx).Save(m).Error
}

func (r *PostgresTemplateRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.TemplateModel{}, id).Error
}
