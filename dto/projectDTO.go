package dto

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=100"`

	Description string `json:"description" binding:"required"`
}