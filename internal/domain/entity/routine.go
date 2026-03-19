package entity

import "time"

type Routine struct {
	ID         int
	CustomerID int
	TemplateID *int
	Name       string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewRoutine(customerID int, templateID *int, name string) *Routine {
	now := time.Now()
	return &Routine{
		CustomerID: customerID,
		TemplateID: templateID,
		Name:       name,
		IsActive:   false,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func (r *Routine) Activate() {
	r.IsActive = true
	r.UpdatedAt = time.Now()
}

func (r *Routine) Deactivate() {
	r.IsActive = false
	r.UpdatedAt = time.Now()
}
