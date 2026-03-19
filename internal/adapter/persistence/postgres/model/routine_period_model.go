package model

import (
	"time"

	"kochappi/internal/domain/entity"
)

type RoutinePeriodModel struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	RoutineID int       `gorm:"type:integer;not null;index"`
	StartedAt time.Time  `gorm:"type:date;not null"`
	EndedAt   *time.Time `gorm:"type:date"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamptz;autoUpdateTime"`
}

func (RoutinePeriodModel) TableName() string {
	return "routine_periods"
}

func (m *RoutinePeriodModel) ToDomain() *entity.RoutinePeriod {
	return &entity.RoutinePeriod{
		ID:        m.ID,
		RoutineID: m.RoutineID,
		StartedAt: m.StartedAt,
		EndedAt:   m.EndedAt,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func RoutinePeriodModelFromDomain(rp *entity.RoutinePeriod) *RoutinePeriodModel {
	return &RoutinePeriodModel{
		ID:        rp.ID,
		RoutineID: rp.RoutineID,
		StartedAt: rp.StartedAt,
		EndedAt:   rp.EndedAt,
		CreatedAt: rp.CreatedAt,
		UpdatedAt: rp.UpdatedAt,
	}
}
