package entity

import (
	"time"

	"kochappi/internal/domain/value_object"
)

type ProgressPhoto struct {
	ID                    int
	URL                   string
	PictureType           value_object.PictureType
	LogCustomerProgressID int
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func NewProgressPhoto(url string, pictureType value_object.PictureType, logCustomerProgressID int) *ProgressPhoto {
	now := time.Now()
	return &ProgressPhoto{
		URL:                   url,
		PictureType:           pictureType,
		LogCustomerProgressID: logCustomerProgressID,
		CreatedAt:             now,
		UpdatedAt:             now,
	}
}
