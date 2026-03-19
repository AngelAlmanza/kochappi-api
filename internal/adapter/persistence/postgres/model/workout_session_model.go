package model

import (
	"time"

	"kochappi/internal/domain/entity"
)

type WorkoutSessionModel struct {
	ID           int        `gorm:"primaryKey;autoIncrement"`
	RoutineID    int        `gorm:"type:integer;not null;index"`
	ScheduledDay int16      `gorm:"type:smallint;not null"`
	ActualDate   time.Time  `gorm:"type:date;not null"`
	Status       string     `gorm:"type:varchar(255);not null"`
	StartedAt    *time.Time `gorm:"type:timestamptz"`
	FinishedAt   *time.Time `gorm:"type:timestamptz"`
	CreatedAt    time.Time  `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"type:timestamptz;autoUpdateTime"`
}

func (WorkoutSessionModel) TableName() string {
	return "workout_session"
}

func (m *WorkoutSessionModel) ToDomain() *entity.WorkoutSession {
	return &entity.WorkoutSession{
		ID:           m.ID,
		RoutineID:    m.RoutineID,
		ScheduledDay: m.ScheduledDay,
		ActualDate:   m.ActualDate,
		Status:       entity.WorkoutStatus(m.Status),
		StartedAt:    m.StartedAt,
		FinishedAt:   m.FinishedAt,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func WorkoutSessionModelFromDomain(ws *entity.WorkoutSession) *WorkoutSessionModel {
	return &WorkoutSessionModel{
		ID:           ws.ID,
		RoutineID:    ws.RoutineID,
		ScheduledDay: ws.ScheduledDay,
		ActualDate:   ws.ActualDate,
		Status:       string(ws.Status),
		StartedAt:    ws.StartedAt,
		FinishedAt:   ws.FinishedAt,
		CreatedAt:    ws.CreatedAt,
		UpdatedAt:    ws.UpdatedAt,
	}
}
