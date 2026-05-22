package dto

type CreateTaskRequest struct {
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description" binding:"required"`

	Status       string `json:"status" binding:"required"`
	Priority     string `json:"priority" binding:"required"`

	ProjectID    uint   `json:"project_id" binding:"required"`
	AssignedToID uint   `json:"assigned_to_id" binding:"required"`
}