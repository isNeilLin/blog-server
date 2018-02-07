package main

import (
	"blog/controllers"
	"blog/models"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.DisableConsoleColor()

	// 设定环境
	gin.SetMode("debug")

	// 记录server日志
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 初始化
	r := gin.Default()
	db := models.InitDB()

	defer db.Close()

	// 路由
	r.GET("/get_all", controllers.GetAll)
	r.GET("/get_publish", controllers.GetPublish)
	r.GET("/get_private", controllers.GetPrivate)
	r.GET("/post", controllers.GetPostById)
	r.POST("/create", controllers.Post)
	r.POST("/update", controllers.Update)
	r.POST("/delete", controllers.Delete)
	r.GET("/get_tags",controllers.GetTags)
	r.GET("/getpostbytag",controllers.GetPostByTag)
	r.POST("/create_tag",controllers.AddTag)
	r.POST("/update_tag",controllers.UpdateTag)
	r.POST("/delete_tag",controllers.DeleteTag)
	r.Run()
}
