package services

import (
	"task-management-api/models"
	"task-management-api/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func GetUsers() ([]models.User, error){
	return  repositories.GetUsers()
}

func GetUser(id string)(models.User, error){
	return  repositories.GetUser(id)
}

func UpdateUser(user *models.User) error{
	return repositories.UpdateUser(user)
}
func DeleteUser(user *models.User) error{
	return  repositories.DeleteUser(user)
}