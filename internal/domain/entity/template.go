package entity

import "time"

type Template struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTemplate(name, description string) *Template {
	now := time.Now()
	return &Template{
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
