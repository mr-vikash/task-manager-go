package main

import (
	"task-manager/database"
	"task-manager/models"
	"task-manager/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Task{})

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8081")
}
