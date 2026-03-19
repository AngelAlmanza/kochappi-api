package dto

type ExerciseResponse struct {
	ID       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"Squat"`
	VideoURL string `json:"videoUrl" example:"https://example.com/squat.mp4"`
}
