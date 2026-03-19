package dto

type CustomerResponse struct {
	ID        int    `json:"id" example:"1"`
	Name      string `json:"name" example:"John Doe"`
	Birthdate string `json:"birthdate" example:"1990-01-01"`
	// Gender    string `json:"gender" example:"male"`
}
