package dto

type RegisterUserRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`

	Email string `json:"email" binding:"required,email"`

	Password string `json:"password" binding:"required,min=4"`

	Role string `json:"role" binding:"required,oneof=Admin User"`
}