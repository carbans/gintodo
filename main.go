package main

import (
	"github.com/gin-gonic/gin"
	"github.com/carbans/gintodo/models"
	"github.com/carbans/gintodo/controllers"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	r.Run()
}