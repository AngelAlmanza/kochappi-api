package model

import (
	"kochappi/internal/domain/entity"

	"time"
)

type CustomerModel struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Birthdate time.Time `gorm:"type:date;not null"`
	UserID    int       `gorm:"type:integer;uniqueIndex;not null"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamptz;autoUpdateTime"`
}

func (CustomerModel) TableName() string {
	return "customers"
}

func (c *CustomerModel) ToDomain() *entity.Customer {
	return &entity.Customer{
		ID:        c.ID,
		Name:      c.Name,
		Birthdate: c.Birthdate,
		UserID:    c.UserID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func CustomerModelFromDomain(customer *entity.Customer) *CustomerModel {
	return &CustomerModel{
		ID:        customer.ID,
		Name:      customer.Name,
		Birthdate: customer.Birthdate,
		UserID:    customer.UserID,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}
