package repositories

import (
	"task-management-api/database"
	"task-management-api/models"
)

func CreateTask(task *models.Task) error {
	if err := database.DB.Create(task).Error; err != nil {
		return err
	}

	return database.DB.
		Preload("Project").
		Preload("AssignedUser").
		First(task, task.ID).Error
}

func GetTasks() ([]models.Task, error) {

	var tasks []models.Task

	err := database.DB.
		Preload("Project").
		Preload("AssignedUser").
		Find(&tasks).Error

	return tasks, err
}

func GetTask(id string)(models.Task, error){
	var task models.Task
	err := database.DB.Preload("Project").Preload("AssignedUser").First(&task, id).Error
	return task, err
}

// func UpdateTask(task *models.Task)error{
// 	return database.DB.Save(task).Error
// }
func UpdateTask(task *models.Task) error {
	return database.DB.Model(&models.Task{}).
		Where("id = ?", task.ID).
		Updates(map[string]interface{}{
			"title":          task.Title,
			"description":    task.Description,
			"status":         task.Status,
			"priority":       task.Priority,
			"project_id":     task.ProjectID,
			"assigned_to_id": task.AssignedToID,
		}).Error
}
func DeleteTask(task *models.Task)error{
	return  database.DB.Delete(task).Error
}