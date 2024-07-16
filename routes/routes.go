package routes

import (
	"taskManagement/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(router *gin.Engine) {
	router.POST("/task", controllers.CreateTask)
	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/task/:id", controllers.GetTaskByID)
	router.PUT("/task/:id", controllers.UpdateTask)
}
