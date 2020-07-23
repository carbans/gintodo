package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/carbans/gintodo/models"
)

// GET /tasks
// Get all books

func FindTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}