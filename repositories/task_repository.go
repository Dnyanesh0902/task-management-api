package repositories

import (
	"task-management-api/database"
	"task-management-api/models"

)

func CreateTask(task *models.Task) error{
	return  database.DB.Create(task).Error
}

func GetTasks()([]models.Task, error){
	var tasks []models.Task
	err := database.DB.Find(tasks).Error
	return tasks, err
}

func GetTask(id string)(models.Task, error){
	var task models.Task
	err := database.DB.First(&task, id).Error
	return task, err
}

func UpdateTask(task *models.Task)error{
	return database.DB.Save(task).Error
}

func DeleteTask(task *models.Task)error{
	return  database.DB.Delete(task).Error
}