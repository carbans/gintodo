package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/carbans/gintodo/models"
	"github.com/carbans/gintodo/controllers"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	//r.GET("/tasks", controllers.FindTasks)

	r.Run()
}