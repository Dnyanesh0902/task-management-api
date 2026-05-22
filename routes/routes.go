package routes

import (
	"task-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router * gin.Engine) {
	api := router.Group("/api")
	userRoutes := api.Group("/users")
	{
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.PUT("/:id",controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}
	taskRoutes := api.Group("/tasks")
	{
		taskRoutes.POST("/", controllers.CreateTask)
		taskRoutes.GET("/", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTask)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}
	projectRoutes := api.Group("/projects")
	{
		projectRoutes.POST("/",controllers.CreateProject)
		projectRoutes.GET("/",controllers.GetProjects)
		projectRoutes.GET("/:id", controllers.GetProjects)
		projectRoutes.PUT("/:id", controllers.UpdateProject)
		projectRoutes.DELETE("/:id", controllers.DeleteProject)
	}
}