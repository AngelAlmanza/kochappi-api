package model

import (
	"time"

	"kochappi/internal/domain/entity"
)

type TemplateModel struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:timestamptz;autoUpdateTime"`
}

func (TemplateModel) TableName() string {
	return "templates"
}

func (m *TemplateModel) ToDomain() *entity.Template {
	return &entity.Template{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func TemplateModelFromDomain(t *entity.Template) *TemplateModel {
	return &TemplateModel{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
