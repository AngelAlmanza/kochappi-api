package entity

import "time"

type ProgressPhoto struct {
	ID                    int
	URL                   string
	PictureType           string
	LogCustomerProgressID int
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func NewProgressPhoto(url, pictureType string, logCustomerProgressID int) *ProgressPhoto {
	now := time.Now()
	return &ProgressPhoto{
		URL:                   url,
		PictureType:           pictureType,
		LogCustomerProgressID: logCustomerProgressID,
		CreatedAt:             now,
		UpdatedAt:             now,
	}
}
