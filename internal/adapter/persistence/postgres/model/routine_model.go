package model

import (
	"time"

	"kochappi/internal/domain/entity"
)

type RoutineModel struct {
	ID         int            `gorm:"primaryKey;autoIncrement"`
	CustomerID int            `gorm:"type:integer;not null;index"`
	TemplateID *int           `gorm:"type:integer"`
	Template   *TemplateModel `gorm:"constraint:OnDelete:SET NULL;foreignKey:TemplateID"`
	Name       string         `gorm:"type:varchar(255);not null"`
	IsActive   bool           `gorm:"type:boolean;not null;default:false"`
	CreatedAt  time.Time      `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"type:timestamptz;autoUpdateTime"`
}

func (RoutineModel) TableName() string {
	return "routines"
}

func (m *RoutineModel) ToDomain() *entity.Routine {
	return &entity.Routine{
		ID:         m.ID,
		CustomerID: m.CustomerID,
		TemplateID: m.TemplateID,
		Name:       m.Name,
		IsActive:   m.IsActive,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func RoutineModelFromDomain(r *entity.Routine) *RoutineModel {
	return &RoutineModel{
		ID:         r.ID,
		CustomerID: r.CustomerID,
		TemplateID: r.TemplateID,
		Name:       r.Name,
		IsActive:   r.IsActive,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}
