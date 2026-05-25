package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"task-management-api/cache"
	"task-management-api/dto"
	"task-management-api/models"
	"task-management-api/services"

	"github.com/gin-gonic/gin"
)
// CreateTask godoc
// @Summary Create task
// @Description Create a new task
// @Tags Tasks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateTaskRequest true "Create Task Request"
// @Success 201 {object} map[string]interface{}
// @Router /tasks [post]
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

	if err := services.CreateTask(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"failed to create task",
		})
		return
	}
	cache.ClearTaskCache()
	c.JSON(http.StatusCreated, gin.H{
		"message":"Task created successfully",
		"data" : task,
	})
}

// GetTasks godoc
// @Summary Get all tasks
// @Description Get tasks with pagination, search and filter
// @Tags Tasks
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Param status query string false "Task status"
// @Param priority query string false "Task priority"
// @Param search query string false "Search keyword"
// @Success 200 {object} map[string]interface{}
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	status := c.Query("status")
	priority := c.Query("priority")
	search := c.Query("search")

	cacheKey := fmt.Sprintf(
		"tasks:page=%d:limit=%d:status=%s:priority=%s:search=%s",
		page,
		limit,
		status,
		priority,
		search,
	)

	cachedData, err := cache.RedisClient.Get(cache.Ctx, cacheKey).Result()

	if err == nil {
		var response gin.H

		json.Unmarshal([]byte(cachedData), &response)

		c.JSON(http.StatusOK, response)
		return
	}

	tasks, total, err := services.GetTasks(page, limit, status, priority, search)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get tasks",
		})
		return
	}

	response := gin.H{
		"data": tasks,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	}

	jsonData, _ := json.Marshal(response)

	cache.RedisClient.Set(cache.Ctx, cacheKey, jsonData, time.Minute*5)

	c.JSON(http.StatusOK, response)
}
// GetTask godoc
// @Summary Get task by ID
// @Description Get single task by task ID
// @Tags Tasks
// @Security BearerAuth
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Router /tasks/{id} [get]
func GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := services.GetTask(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"task not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":task,
	})
}
// UpdateTask godoc
// @Summary Update task
// @Description Update task by ID
// @Tags Tasks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param request body dto.CreateTaskRequest true "Update Task Request"
// @Success 200 {object} map[string]interface{}
// @Router /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	task, err := services.GetTask(id)
	if err != nil{
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

	if err := services.UpdateTask(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"failed to update task",
		})
		return
	}
	task, _ = services.GetTask(id)
	cache.ClearTaskCache()
	c.JSON(http.StatusOK, gin.H{
		"message":"task updated successfully",
		"data":task,
	})
}
// DeleteTask godoc
// @Summary Delete task
// @Description Delete task by ID
// @Tags Tasks
// @Security BearerAuth
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context){
	id := c.Param("id")
	task,err := services.GetTask(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"task not found",
		})
		return
	}
	if err := services.DeleteTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	cache.ClearTaskCache()
	c.JSON(http.StatusOK, gin.H{
		"message":"Task deleted successfully",
	})
}