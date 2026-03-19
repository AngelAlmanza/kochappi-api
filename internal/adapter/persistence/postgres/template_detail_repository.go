package postgres

import (
	"context"
	"errors"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"

	"gorm.io/gorm"
)

type PostgresTemplateDetailRepository struct {
	db *gorm.DB
}

func NewPostgresTemplateDetailRepository(db *gorm.DB) *PostgresTemplateDetailRepository {
	return &PostgresTemplateDetailRepository{db: db}
}

func (r *PostgresTemplateDetailRepository) GetByTemplateID(ctx context.Context, templateID int) ([]entity.TemplateDetail, error) {
	var models []model.TemplateDetailModel
	if err := r.db.WithContext(ctx).Where("template_id = ?", templateID).Order("display_order ASC").Find(&models).Error; err != nil {
		return nil, err
	}

	details := make([]entity.TemplateDetail, 0, len(models))
	for _, m := range models {
		details = append(details, *m.ToDomain())
	}
	return details, nil
}

func (r *PostgresTemplateDetailRepository) GetByID(ctx context.Context, id int) (*entity.TemplateDetail, error) {
	var m model.TemplateDetailModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.TemplateDetailNotFoundError{ID: id}
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *PostgresTemplateDetailRepository) Create(ctx context.Context, detail *entity.TemplateDetail) error {
	m := model.TemplateDetailModelFromDomain(detail)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	detail.ID = m.ID
	return nil
}

func (r *PostgresTemplateDetailRepository) CreateBulk(ctx context.Context, details []*entity.TemplateDetail) error {
	if len(details) == 0 {
		return nil
	}

	models := make([]*model.TemplateDetailModel, 0, len(details))
	for _, d := range details {
		models = append(models, model.TemplateDetailModelFromDomain(d))
	}

	if err := r.db.WithContext(ctx).Create(&models).Error; err != nil {
		return err
	}

	for i, m := range models {
		details[i].ID = m.ID
	}
	return nil
}

func (r *PostgresTemplateDetailRepository) DeleteByID(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.TemplateDetailModel{}, id).Error
}
