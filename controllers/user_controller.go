package controllers 

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"task-management-api/dto"
	"task-management-api/database"
	"task-management-api/models"
	
)

func CreateUser(c *gin.Context) {
	var input dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := models.User {
		Name:  input.Name,
		Email: input.Email,
		Password: input.Password,
		Role: input.Role,
	}

	if err :=database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"failed to create user",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":"User Created Successfully",
		"data": user,
	})
}


func GetUsers(c *gin.Context){
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message": "failed to fetch user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"data": users,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"User Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H {
		"data":user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"user not found",
		})
		return
	}
	
	var input dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.Role = input.Role

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to update user",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":"user updated successfully.",
		"data": user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message":"failed to delete user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"user deleted successfully",
	})
}