package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"task-management-api/models"
	"task-management-api/database"
	"task-management-api/dto"
)

func CreateProject(c *gin.Context){
	var input dto.CreateProjectRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	project := models.Project{
		Name:  input.Name,
		Description: input.Description,
	}
	if err := database.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to create",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message" :"Project created successfully",
		"data": project,
	})
}

func GetProjects(c *gin.Context) {
	var project []models.Project

	if err := database.DB.Find(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": project,
	})
}

func GetProject(c *gin.Context){
	id := c.Param("id")
	var project models.Project

	if err := database.DB.First(&project, id).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"project not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":project,
	})
}

func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project

	if err := database.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"project not found",
		})
		return
	}

	var input dto.CreateProjectRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":err.Error(),
		})
		return
	}
	project.Name = input.Name
	project.Description= input.Description
	if err := database.DB.Save(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"failed to update project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"project updated successfully",
		"data":project,
	})
}

func DeleteProject(c *gin.Context){
	id := c.Param("id")
	var project models.Project

	if err := database.DB.First(&project, id).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"project not found",
		})
		return
	}
	if err := database.DB.Delete(&project).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to delete project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"Project deleted successfully",
	})
}