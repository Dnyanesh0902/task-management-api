package controllers

import (
	"net/http"
	"strconv"
	"task-management-api/dto"
	"task-management-api/models"
	"task-management-api/services"
	"task-management-api/utils"
	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.RegisterUserRequest true "Create User Request"
// @Success 201 {object} map[string]interface{}
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var input dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to hash password",
		})
		return
	}
	user := models.User {
		Name:  input.Name,
		Email: input.Email,
		Password: hashedPassword,
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

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
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

// GetUser godoc
// @Summary Get user by ID
// @Description Get single user by ID
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [get]
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

// UpdateUser godoc
// @Summary Update user
// @Description Update user by ID
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.RegisterUserRequest true "Update User Request"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [put]
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

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by ID
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [delete]
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