package entity

import "time"

type LogCustomerProgress struct {
	ID         int
	CheckDate  time.Time
	Weight     int
	CustomerID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewLogCustomerProgress(customerID int, checkDate time.Time, weight int) *LogCustomerProgress {
	now := time.Now()
	return &LogCustomerProgress{
		CustomerID: customerID,
		CheckDate:  checkDate,
		Weight:     weight,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
