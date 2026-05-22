package services

import (
	
	"task-management-api/models"
	"task-management-api/repositories"
)

func CreateProject(project *models.Project) error{
	return repositories.CreateProject(project)
}

func GetProjects()([]models.Project, error){
	return  repositories.GetProjects()
}

func GetProject(id string)(models.Project, error){
	return  repositories.GetProject(id)
}

func UpdateProject(project *models.Project) error{
	return  repositories.UpdateProject(project)
}
func DeleteProject(project *models.Project) error{
	return  repositories.DeleteProject(project)
}