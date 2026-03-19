package model

import (
	"time"

	"kochappi/internal/domain/entity"
)

type TemplateDetailModel struct {
	ID           int            `gorm:"primaryKey;autoIncrement"`
	TemplateID   int            `gorm:"type:integer;not null;index"`
	Template     *TemplateModel `gorm:"constraint:OnDelete:CASCADE;foreignKey:TemplateID"`
	DayOfWeek    int16          `gorm:"type:smallint;not null"`
	ExerciseID   int            `gorm:"type:integer;not null"`
	Sets         int16          `gorm:"type:smallint;not null"`
	Reps         int16          `gorm:"type:smallint;not null"`
	DisplayOrder int16          `gorm:"type:smallint;not null"`
	CreatedAt    time.Time      `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"type:timestamptz;autoUpdateTime"`
}

func (TemplateDetailModel) TableName() string {
	return "template_detail"
}

func (m *TemplateDetailModel) ToDomain() *entity.TemplateDetail {
	return &entity.TemplateDetail{
		ID:           m.ID,
		TemplateID:   m.TemplateID,
		DayOfWeek:    m.DayOfWeek,
		ExerciseID:   m.ExerciseID,
		Sets:         m.Sets,
		Reps:         m.Reps,
		DisplayOrder: m.DisplayOrder,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func TemplateDetailModelFromDomain(td *entity.TemplateDetail) *TemplateDetailModel {
	return &TemplateDetailModel{
		ID:           td.ID,
		TemplateID:   td.TemplateID,
		DayOfWeek:    td.DayOfWeek,
		ExerciseID:   td.ExerciseID,
		Sets:         td.Sets,
		Reps:         td.Reps,
		DisplayOrder: td.DisplayOrder,
		CreatedAt:    td.CreatedAt,
		UpdatedAt:    td.UpdatedAt,
	}
}
