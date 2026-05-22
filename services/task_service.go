package services

import (

	"task-management-api/models"
	"task-management-api/repositories"
)

func CreateTask(task *models.Task) error{
	return repositories.CreateTask(task)
}

func GetTasks()([]models.Task, error){
	return repositories.GetTasks()
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