package main

import (
	"io"
	"os"
	"github.com/gin-gonic/gin"
	"blog/models"
	"blog/controllers"
)

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	db := models.InitDB()

	defer db.Close()
	r.GET("/get_all",controllers.GetAll)
	r.GET("/get_publish",controllers.GetPublish)
	r.POST("/create", controllers.Post)
	r.POST("/update",controllers.Update)
	r.POST("/delete",controllers.Delete)
	r.Run()
}

