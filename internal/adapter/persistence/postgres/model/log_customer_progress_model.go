package model

import (
	"time"

	"kochappi/internal/domain/entity"
)

type LogCustomerProgressModel struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	CheckDate  time.Time `gorm:"type:date;not null"`
	Weight     int       `gorm:"type:integer;not null"`
	CustomerID int       `gorm:"type:integer;not null;index"`
	CreatedAt  time.Time `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"type:timestamptz;autoUpdateTime"`
}

func (LogCustomerProgressModel) TableName() string {
	return "log_customer_progress"
}

func (m *LogCustomerProgressModel) ToDomain() *entity.LogCustomerProgress {
	return &entity.LogCustomerProgress{
		ID:         m.ID,
		CheckDate:  m.CheckDate,
		Weight:     m.Weight,
		CustomerID: m.CustomerID,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func LogCustomerProgressModelFromDomain(l *entity.LogCustomerProgress) *LogCustomerProgressModel {
	return &LogCustomerProgressModel{
		ID:         l.ID,
		CheckDate:  l.CheckDate,
		Weight:     l.Weight,
		CustomerID: l.CustomerID,
		CreatedAt:  l.CreatedAt,
		UpdatedAt:  l.UpdatedAt,
	}
}
