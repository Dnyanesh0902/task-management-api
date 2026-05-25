package services

import (

	"task-management-api/models"
	"task-management-api/repositories"
)

func CreateTask(task *models.Task) error{
	return repositories.CreateTask(task)
}

func GetTasks(page int, limit int, status string, priority string, search string)([]models.Task, int64, error){
	return repositories.GetTasks(page, limit, status, priority, search)
}

func GetTask(id string)(models.Task, error){
	return repositories.GetTask(id)
}

func UpdateTask(task *models.Task) error{
	return repositories.UpdateTask(task)
}

func DeleteTask(task *models.Task) error{
	return repositories.DeleteTask(task)
}