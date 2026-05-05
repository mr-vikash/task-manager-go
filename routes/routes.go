package routes

import (
	"task-manager/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)

	r.POST("tasks", controllers.CreateTask)
	r.GET("tasks", controllers.GetTasks)
	r.PUT("tasks/:id", controllers.UpdateTask)
	r.DELETE("tasks/:id", controllers.DeleteTask)
}
