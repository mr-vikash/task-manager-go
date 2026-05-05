package controllers

import (
	"strconv"
	"task-manager/database"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}
	database.DB.Create(&task)
	c.JSON(200, task)
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(200, tasks)
}

func UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	var task models.Task

	err = database.DB.First(&task, id).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Task not Found",
		})
		return
	}
	var input models.Task

	err = c.BindJSON(&input)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Input",
		})
		return
	}

	err = database.DB.Model(&task).Updates(&input).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to update the task",
		})
		return
	}

	database.DB.First(&task, id)

	c.JSON(200, gin.H{
		"message": "Task updated successfully",
		"data":    task,
	})

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	err := database.DB.Find(&task, id).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Task not found",
		})
	}

	err = database.DB.Delete(&task).Error

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Something went wrong",
		})
	}

	c.JSON(200, gin.H{
		"message": "Task Deleted Successfully",
	})
}
