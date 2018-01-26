package main

import (
	"io"
	"os"
	"github.com/gin-gonic/gin"
	Controller "blog/controllers"
)

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.GET("/ping", Controller.Ping)
	r.GET("/user/:name/*action", Controller.User)
	r.GET("/welcome", Controller.Welcome)
	r.POST("/form_post", Controller.FormPost)
	r.POST("/upload", Controller.Upload)

	r.Run()
}

