package controllers

import (
	"net/http"
	"task-management-api/database"
	"task-management-api/dto"
	"task-management-api/models"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var input dto.CreateTaskRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	task := models.Task {
		Title:  input.Title,
		Description: input.Description,
		Status: input.Status,
		Priority: input.Priority,
		ProjectID: input.ProjectID,
		AssignedToID: input.AssignedToID,
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"failed to create task",
		})
		return
	}
	database.DB.
		Preload("Project").
		Preload("AssignedUser").
		First(&task, task.ID)
		
	c.JSON(http.StatusCreated, gin.H{
		"message":"Task created successfully",
		"data" : task,
	})
}


func GetTasks(c *gin.Context) {
	var task []models.Task

	if err := database.DB.Preload("Project").Preload("AssignedUser").Find(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to get tasks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		
		"data":task,
	})
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := database.DB.Preload("Project").Preload("AssignedUser").First(&task, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"task not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":task,
	})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"task not found",
		})
		return
	}
	var input dto.CreateTaskRequest

	if err :=c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":err.Error(),
		})
		return
	}
	task.Title = input.Title
	task.Description = input.Description
	task.Status = input.Status
	task.Priority = input.Priority
	task.ProjectID = input.ProjectID
	task.AssignedToID = input.AssignedToID

	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message":"failed to update task",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"task updated successfully",
		"data":task,
	})
}

func DeleteTask(c *gin.Context){
	id := c.Param("id")
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"task not found",
		})
		return
	}
	if err := database.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"Task deleted successfully",
	})
}