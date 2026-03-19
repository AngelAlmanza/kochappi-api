package entity

import "time"

type Exercise struct {
	ID        int
	Name      string
	VideoURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewExercise(name, videoURL string) *Exercise {
	now := time.Now()
	return &Exercise{
		Name:      name,
		VideoURL:  videoURL,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
