package routes

import (
	"task-management-api/controllers"
	"task-management-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	userRoutes := api.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.POST("/",middleware.AuthorizeRole("Admin"), controllers.CreateUser)
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.PUT("/:id",middleware.AuthorizeRole("Admin"), controllers.UpdateUser)
		userRoutes.DELETE("/:id", middleware.AuthorizeRole("Admin"), controllers.DeleteUser)
	}
	taskRoutes := api.Group("/tasks")
	taskRoutes.Use(middleware.AuthMiddleware())
	{
		taskRoutes.POST("/", middleware.AuthorizeRole("Admin"), controllers.CreateTask)
		taskRoutes.GET("/", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTask)
		taskRoutes.PUT("/:id", middleware.AuthorizeRole("Admin"), controllers.UpdateTask)
		taskRoutes.DELETE("/:id", middleware.AuthorizeRole("Admin"), controllers.DeleteTask)
	}
	projectRoutes := api.Group("/projects")
	projectRoutes.Use(middleware.AuthMiddleware())
	{
		projectRoutes.POST("/",middleware.AuthorizeRole("Admin"), controllers.CreateProject)
		projectRoutes.GET("/",controllers.GetProjects)
		projectRoutes.GET("/:id", controllers.GetProject)
		projectRoutes.PUT("/:id", middleware.AuthorizeRole("Admin"), controllers.UpdateProject)
		projectRoutes.DELETE("/:id",middleware.AuthorizeRole("Admin"), controllers.DeleteProject)
	}
	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}
}