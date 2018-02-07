package main

import (
	"blog/models"
	"blog/router"
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

	db := models.InitDB()
	defer db.Close()

	// 初始化路由
	router.InitRouter()
}
