package routes

import (
	"task-manager/controllers"
	"task-manager/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)

	r.POST("/login", controllers.Login)

	protected := r.Group("/api")

	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("tasks", controllers.CreateTask)
		protected.GET("tasks", controllers.GetTasks)
		protected.PUT("tasks/:id", controllers.UpdateTask)
		protected.DELETE("tasks/:id", controllers.DeleteTask)
	}
}
