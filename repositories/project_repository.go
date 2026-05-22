package repositories

import (
	"task-management-api/database"
	"task-management-api/models"
)

func CreateProject(project *models.Project) error{
	return database.DB.Create(project).Error
}

func GetProjects()([]models.Project, error){
	var project []models.Project
	err := database.DB.Find(&project).Error
	return project, err
}

func GetProject(id string)(models.Project, error){
	var project models.Project
	err := database.DB.First(&project, id).Error
	return  project, err
}

func UpdateProject(project *models.Project) error{
	return database.DB.Save(project).Error
}

func DeleteProject(project *models.Project) error{
	return  database.DB.Delete(project).Error
}