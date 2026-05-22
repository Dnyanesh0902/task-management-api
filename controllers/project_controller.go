package controllers

import (
	"net/http"
	"task-management-api/dto"
	"task-management-api/models"
	"task-management-api/services"
	"github.com/gin-gonic/gin"
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
	if err := services.CreateProject(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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
	project, err := services.GetProjects()
	if err != nil {
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
	project, err :=services.GetProject(id)

	if err != nil{
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
	project, err := services.GetProject(id)
	if err != nil {
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
	if err := services.UpdateProject(&project); err != nil {
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
	project, err := services.GetProject(id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"project not found",
		})
		return
	}
	if err := services.DeleteProject(&project); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to delete project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"Project deleted successfully",
	})
}