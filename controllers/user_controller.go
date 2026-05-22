package controllers

import (
	"net/http"
	"strconv"
	"task-management-api/dto"
	"task-management-api/models"
	"task-management-api/services"

	"github.com/gin-gonic/gin"
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

	if err := services.CreateUser(&user); err != nil {
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
	user, err := services.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message": "failed to fetch user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"data": user,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUser(id)

	if  err != nil {
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
	user, err := services.GetUser(id)
	if err != nil {
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

	if err := services.UpdateUser(&user); err != nil {
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
	_, err := strconv.Atoi(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Invalid User ID",
		})
		return
	}
	user, err := services.GetUser(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"student not found",
		})
		return
	}
	if err := services.DeleteUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message":"failed to delete user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"user deleted successfully",
	})
}