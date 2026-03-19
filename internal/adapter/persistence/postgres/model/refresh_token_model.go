package model

import "time"

type RefreshTokenModel struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	UserID    string    `gorm:"type:uuid;index;not null"`
	ExpiresAt time.Time `gorm:"type:timestamptz;not null"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null"`
}

func (RefreshTokenModel) TableName() string {
	return "refresh_tokens"
}
