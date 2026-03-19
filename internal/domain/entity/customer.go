package entity

import "time"

type Customer struct {
	ID        int
	Name      string
	Birthdate time.Time
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(name string, birthdate time.Time, userID int) *Customer {
	return &Customer{
		Name:      name,
		Birthdate: birthdate,
		UserID:    userID,
	}
}

func (c *Customer) GetAge() int {
	now := time.Now()
	years := now.Year() - c.Birthdate.Year()
	if now.Month() < c.Birthdate.Month() || (now.Month() == c.Birthdate.Month() && now.Day() < c.Birthdate.Day()) {
		years--
	}
	return years
}
