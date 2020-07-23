package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/carbans/gintodo/models"
)


type CreateTaskInput struct {
	Title  string `json:"title" binding:"required"`
	Owner string `json:"owner" binding:"required"`
}

type UpdateTaskInput struct {
	Title  string `json:"title"`
	Owner string `json:"owner"`
}


// GET /tasks
// Get all tasks

func FindTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context){
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	task := models.Task{Title: input.Title, Owner: input.Owner}
	models.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// GET /tasks/:id
// Find a task

func FindTask(c *gin.Context){
	var task models.Task

	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}


// PATCH /tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {
	// Get model if exist
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	models.DB.Model(&task).Updates(input)
  
	c.JSON(http.StatusOK, gin.H{"data": task})
  }

// DELETE /tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	models.DB.Delete(&task)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }
