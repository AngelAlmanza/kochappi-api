package dto

type ProgressLogResponse struct {
	ID         int    `json:"id" example:"1"`
	CheckDate  string `json:"checkDate" example:"2024-01-15"`
	Weight     int    `json:"weight" example:"75"`
	CustomerID int    `json:"customerId" example:"1"`
}

type ProgressPhotoResponse struct {
	ID          int    `json:"id" example:"1"`
	URL         string `json:"url" example:"/uploads/progress_1_abc.jpg"`
	PictureType string `json:"pictureType" example:"front"`
	LogID       int    `json:"logId" example:"1"`
}

type ProgressLogWithPhotosResponse struct {
	ID         int                     `json:"id" example:"1"`
	CheckDate  string                  `json:"checkDate" example:"2024-01-15"`
	Weight     int                     `json:"weight" example:"75"`
	CustomerID int                     `json:"customerId" example:"1"`
	Photos     []ProgressPhotoResponse `json:"photos"`
}
