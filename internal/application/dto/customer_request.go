package dto

type CreateCustomerRequest struct {
	Name      string `json:"name" binding:"required" example:"Jhon Doe"`
	Birthdate string `json:"birthdate" binding:"required" example:"1990-01-01"`
	UserID    int    `json:"userId" binding:"required" example:"1"`
}

type UpdateCustomerRequest struct {
	// ID        int    `json:"id" binding:"required" example:"1"`
	Name      string `json:"name" binding:"required" example:"Jhon Doe"`
	Birthdate string `json:"birthdate" binding:"required" example:"1990-01-01"`
}

// type DeleteCustomerRequest struct {
// 	ID int `json:"id" binding:"required" example:"1"`
// }

// type GetCustomerRequest struct {
// 	ID int `json:"id" binding:"required" example:"1"`
// }
