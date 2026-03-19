package model

import (
	"time"

	"kochappi/internal/domain/entity"
)

type UserModel struct {
	ID           string     `gorm:"type:uuid;primaryKey"`
	Name         string     `gorm:"type:varchar(255);not null"`
	Email        string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash string     `gorm:"type:varchar(255);not null"`
	Role         string     `gorm:"type:varchar(20);not null"`
	OTPCode      string     `gorm:"type:varchar(6)"`
	OTPExpiresAt *time.Time `gorm:"type:timestamptz"`
	CreatedAt    time.Time  `gorm:"type:timestamptz;not null"`
	UpdatedAt    time.Time  `gorm:"type:timestamptz;not null"`
}

func (UserModel) TableName() string {
	return "users"
}

func (m *UserModel) ToDomainEntity() *entity.User {
	return &entity.User{
		ID:           m.ID,
		Name:         m.Name,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		Role:         entity.Role(m.Role),
		OTPCode:      m.OTPCode,
		OTPExpiresAt: m.OTPExpiresAt,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func UserModelFromEntity(user *entity.User) *UserModel {
	return &UserModel{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Role:         string(user.Role),
		OTPCode:      user.OTPCode,
		OTPExpiresAt: user.OTPExpiresAt,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
