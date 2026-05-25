package controllers

import (
	"net/http"
	"task-management-api/database"
	"task-management-api/dto"
	"task-management-api/models"
	"task-management-api/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	if err :=c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Failed to hash password",
		})
		return
	}
	user.Password = hashedPassword

	if err := database.DB.Create(&user).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Email already exists or invalid data",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":"User Registered successfully.",
		"data": gin.H{
			"id": user.ID,
			"name":user.Name,
			"email": user.Email,
			"role": user.Role,
		},
	})
}

func Login(c *gin.Context){
	var input dto.LoginRequest
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":"Invalid email or password",
		})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password){
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":"Invalid email or password",
		})
		return
	}
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Failed to generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"login successful",
		"token":token,
	})
}

