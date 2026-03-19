package dto

type CreateExerciseRequest struct {
	Name     string `json:"name" binding:"required" example:"Squat"`
	VideoURL string `json:"videoUrl" example:"https://example.com/squat.mp4"`
}

type UpdateExerciseRequest struct {
	Name     string `json:"name" binding:"required" example:"Squat"`
	VideoURL string `json:"videoUrl" example:"https://example.com/squat.mp4"`
}
